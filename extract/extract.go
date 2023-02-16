package extract

import (
	"errors"
	"github.com/idiomatic-go/accessevents/data"
	//"github.com/idiomatic-go/motif/runtime"
	//"github.com/idiomatic-go/motif/template"
	//"github.com/idiomatic-go/resiliency/actuator"
	"net/http"
	"reflect"
	"strings"
)

type ErrorHandleFn func(location string, errs ...error)

type messageHandler func(l *data.Entry) bool
type pkg struct{}

var (
	pkgPath      = reflect.TypeOf(any(pkg{})).PkgPath()
	locInit      = pkgPath + "/initialize"
	locDo        = pkgPath + "/do"
	url          string
	c            chan *data.Entry
	client                      = http.DefaultClient
	handler      messageHandler = do
	errorHandler ErrorHandleFn
	operators    = []data.Operator{
		{Name: "start_time", Value: data.StartTimeOperator},
		{Name: "duration_ms", Value: data.DurationOperator},
		{Name: "traffic", Value: data.TrafficOperator},
		{Name: "route_name", Value: data.RouteNameOperator},

		{Name: "region", Value: data.OriginRegionOperator},
		{Name: "zone", Value: data.OriginZoneOperator},
		//{Name: "sub_zone", Value: accessdata.OriginSubZoneOperator},
		{Name: "service", Value: data.OriginServiceOperator},
		{Name: "instance_id", Value: data.OriginInstanceIdOperator},

		{Name: "method", Value: data.RequestMethodOperator},
		{Name: "host", Value: data.RequestHostOperator},
		{Name: "path", Value: data.RequestPathOperator},
		{Name: "protocol", Value: data.RequestProtocolOperator},
		{Name: "request_id", Value: data.RequestIdOperator},
		{Name: "forwarded", Value: data.RequestForwardedForOperator},

		{Name: "status_code", Value: data.ResponseStatusCodeOperator},
		{Name: "status_flags", Value: data.StatusFlagsOperator},
		//{Name: "start_time", Value: data.ResponseBytesReceivedOperator},
		//{}Name: "start_time", Value: data.ResponseBytesSentOperator},

		{Name: "timeout_ms", Value: data.TimeoutDurationOperator},
		{Name: "rate_limit", Value: data.RateLimitOperator},
		{Name: "rate_burst", Value: data.RateBurstOperator},
		{Name: "retry", Value: data.RetryOperator},
		{Name: "retry_rate_limit", Value: data.RetryRateLimitOperator},
		{Name: "retry_rate_burst", Value: data.RetryRateBurstOperator},
		{Name: "failover", Value: data.FailoverOperator},
	}
)

/*
// Initialize - templated function to initialize extract
func Initialize[E template.ErrorHandler](uri string, newClient *http.Client) *runtime.Status {
	errorHandler = template.NewErrorHandleFn[E]()
	if uri == "" {
		return errorHandler(locInit, errors.New("invalid argument: uri is empty"))
	}
	u, err1 := urlpkg.Parse(uri)
	if err1 != nil {
		return errorHandler(locInit, err1)
	}
	url = u.String()
	c = make(chan *data.Entry, 100)
	go receive()
	if newClient != nil {
		client = newClient
	}
	//actuator.EnableExtract(extract)
	return runtime.NewStatusOK()
}


*/
func Shutdown() {
	if c != nil {
		close(c)
	}
}

func extract(entry *data.Entry) {
	if entry != nil {
		c <- entry
	}
}

func do(entry *data.Entry) bool {
	if entry == nil {
		errorHandler(locDo, errors.New("invalid argument: access log data is nil"))
		return false
	}
	// let's not extract the extract, the extract, the extract ...
	if entry.Url == url {
		return false
	}
	var req *http.Request
	var err error

	reader := strings.NewReader(data.WriteJson(operators, entry))
	req, err = http.NewRequest(http.MethodPost, url, reader)
	if err == nil {
		_, err = client.Do(req)
	}
	if err != nil {
		errorHandler(locDo, err)
		return false
	}
	return true
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			go handler(msg)
		default:
		}
	}
}
