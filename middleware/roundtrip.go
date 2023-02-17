package middleware

import (
	"errors"
	"github.com/idiomatic-go/accessevents/data"
	"github.com/idiomatic-go/accessevents/log"
	"net/http"
	"time"
)

type wrapper struct {
	rt http.RoundTripper
}

// RoundTrip - implementation of the RoundTrip interface for a transport, also logs an access entry
func (w *wrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	var start = time.Now().UTC()

	// !panic
	if w == nil || w.rt == nil {
		return nil, errors.New("invalid handler round tripper configuration : http.RoundTripper is nil")
	}
	resp, err := w.rt.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	entry := data.NewHttpEntry(data.EgressTraffic, start, time.Since(start), req, resp, "", nil)
	log.Write[log.LogOutputHandler, data.JsonFormatter](entry)
	return resp, nil
}

func WrapDefaultTransport() {
	if http.DefaultClient.Transport == nil {
		http.DefaultClient.Transport = &wrapper{http.DefaultTransport}
	} else {
		http.DefaultClient.Transport = WrapRoundTripper(http.DefaultClient.Transport)
	}
}

func WrapRoundTripper(rt http.RoundTripper) http.RoundTripper {
	return &wrapper{rt}
}
