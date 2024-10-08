package requtil

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Method int

const (
	MethodUnknown = iota
	POST
	GET
)

const (
	TypeJSON       = "json"
	TypeXML        = "xml"
	TypeUrlencoded = "urlencoded"
	TypeForm       = "form"
	TypeFormData   = "form-data"
	TypeHTML       = "html"
	TypeText       = "text"
	TypeMultipart  = "multipart"

	errHttpStatusTooEarly = "Client.Timeout exceeded while awaiting headers"
)

type Req struct {
	ApiName    string
	Url        string
	Method     Method
	TargetType string

	FileName  string // for file request
	Query     map[string]string
	Body      interface{}
	Headers   map[string]string
	Cookies   []*http.Cookie
	RequestId string

	TimeOut     time.Duration
	RateLimiter *rate.Limiter
	RespCheck   func(content []byte) (err error)
}
