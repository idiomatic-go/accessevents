package middleware

import (
	"github.com/felixge/httpsnoop"
	"github.com/idiomatic-go/accessevents/data"
	"github.com/idiomatic-go/accessevents/log"
	"net/http"
	"time"
)

// HttpHostMetricsHandler - http handler that captures metrics about an ingress request, also logs an access
// entry
func HttpHostMetricsHandler(appHandler http.Handler, msg string) http.Handler {
	wrappedH := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now().UTC()
		m := httpsnoop.CaptureMetrics(appHandler, w, req)
		//log.Printf("%s %s (code=%d dt=%s written=%d)", r.Method, r.URL, m.Code, m.Duration, m.Written)
		resp := new(http.Response)
		resp.StatusCode = m.Code
		resp.ContentLength = m.Written
		entry := data.NewHttpEntry(data.IngressTraffic, start, time.Since(start), req, resp, "", nil)
		log.Write[log.LogOutputHandler, data.JsonFormatter](entry)
	})
	return wrappedH
}
