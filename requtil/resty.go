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

	"github.com/go-resty/resty/v2"
	"moul.io/http2curl"
)

type RestyRetry struct {
	Attempts   uint
	Wait       time.Duration
	MaxWait    time.Duration
	RetryAfter resty.RetryAfterFunc
}

var (
	logger           Logger = &DullLogger{}
	errLoggerNotInit        = errors.New("logger not initialized")
	defaultRetry            = &RestyRetry{
		Attempts: 2,
		Wait:     1 * time.Second,
		MaxWait:  3 * time.Second,
		RetryAfter: func(c *resty.Client, resp *resty.Response) (time.Duration, error) {
			logger.Warningf("url=%+v, retryCount=%d, err=%+v", c.BaseURL, c.RetryCount, resp.Error())
			return 0, nil // default backoff
		},
	}
)

func Init(loggerIns Logger) {
	logger = loggerIns
}

func NewLongConnClient(timeout time.Duration, retry *RestyRetry) *resty.Client {
	if retry == nil {
		retry = defaultRetry
	}
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
		}).
		SetRetryCount(int(retry.Attempts)).
		SetRetryWaitTime(retry.Wait).
		SetRetryMaxWaitTime(retry.MaxWait).
		SetRetryAfter(retry.RetryAfter)
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
	client *resty.Client, req *Req,
) (httpCode int, resp []byte, err error) {
	if logger == nil {
		return httpCode, nil, errLoggerNotInit
	}
	if req.ApiName == "" {
		return httpCode, nil, errors.New("api name not set")
	}

	if req.RateLimiter != nil {
		locErr := req.RateLimiter.Wait(context.Background())
		if locErr != nil {
			logger.Warningf("rate limiter wait error, err=%s", locErr)
			// won't return
		}
	}
	return requestLongConn(client, req)
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
