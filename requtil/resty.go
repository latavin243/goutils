package requtil

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/go-resty/resty/v2"
	"moul.io/http2curl"

	"github.com/latavin243/goutils/fnwrap"
)

var (
	logger           Logger = &DullLogger{}
	errLoggerNotInit        = errors.New("logger not initialized")
)

func Init(loggerIns Logger) {
	logger = loggerIns
}

func NewLongConnClient(timeout time.Duration) *resty.Client {
	return resty.New().
		SetTimeout(timeout).
		SetTransport(NewLongConnTransport()).
		SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
			if logger == nil {
				return errLoggerNotInit
			}
			cmd, _ := http2curl.GetCurlCommand(r)
			logger.Infof("request curl command: %s", cmd.String())
			return nil
		})
}

func NewLongConnTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          500,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
}

func RequestLongConn(
	client *resty.Client, req *Req, retrySettings *RetrySettings,
) (httpCode int, resp []byte, err error) {
	if logger == nil {
		return httpCode, nil, errLoggerNotInit
	}
	if req.ApiName == "" {
		return httpCode, nil, errors.New("api name not set")
	}
	if retrySettings == nil {
		retrySettings = &RetrySettings{
			Attempts:        2,
			Delay:           1 * time.Second,
			OnRetryCallback: func(n uint, err error) { logger.Warningf("apiName=%+v, retry=%d, err=%s", req.ApiName, n, err) },
		}
	}

	requestFunc := func() error {
		var locErr error
		if req.RateLimiter != nil {
			locErr = req.RateLimiter.Wait(context.Background())
			if locErr != nil {
				logger.Warningf("rate limiter wait error, err=%s", locErr)
				// won't return
			}
		}
		httpCode, resp, locErr = requestLongConn(client, req)
		return locErr
	}

	err = fnwrap.New(requestFunc).
		WithRetry(
			retry.Attempts(retrySettings.Attempts),
			retry.Delay(retrySettings.Delay),
			retry.OnRetry(retrySettings.OnRetryCallback),
		).
		Do()
	return httpCode, resp, err
}

func requestLongConn(
	client *resty.Client, req *Req,
) (statusCode int, content []byte, err error) {
	if req.TimeOut < 3*time.Second {
		req.TimeOut = 3 * time.Second
	}

	request := client.R().
		SetBody(req.Body).
		SetQueryParams(req.Query).
		SetHeaders(req.Headers).
		SetHeader("X-Request-Id", req.RequestId)

	if req.Cookies != nil && len(req.Cookies) != 0 {
		request.SetCookies(req.Cookies)
	}

	// start := time.Now()
	var rawResp *resty.Response
	switch req.Method {
	case POST:
		rawResp, err = request.Post(req.Url)
	case GET:
		rawResp, err = request.Get(req.Url)
	default:
		rawResp, err = request.Get(req.Url)
	}
	// elapsed := float64(time.Since(start) / time.Millisecond)

	if err != nil {
		statusCode = http.StatusInternalServerError
		logger.Errorf("resty request err, err=%s", err)
		if strings.Contains(err.Error(), errHttpStatusTooEarly) {
			statusCode = http.StatusTooEarly
		}
		return statusCode, nil, err
	}

	statusCode = rawResp.StatusCode()
	respContent := rawResp.Body()
	if statusCode != http.StatusOK {
		logger.Errorf("request [%s] failed, req=%+v, respContent=%s, statusCode=%d", req.ApiName, req, string(respContent), statusCode)
		return statusCode, respContent, fmt.Errorf("request [%s] failed", req.ApiName)
	}

	// custom resp content check
	if req.RespCheck != nil {
		err := req.RespCheck(respContent)
		if err != nil {
			return statusCode, respContent, fmt.Errorf("request [%s] failed, content error", req.ApiName)
		}
	}

	return statusCode, respContent, nil
}
