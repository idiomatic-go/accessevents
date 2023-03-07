package middleware

import (
	"errors"
	"github.com/idiomatic-go/accessevents/data"
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
	entry := data.NewEgressEntry(start, time.Since(start), req, resp, "", nil)
	defaultLogFn(entry)
	return resp, nil
}

// WrapTransport - provides a RoundTrip wrapper that applies controller controllers
func WrapTransport(client *http.Client) {
	if client == nil || client == http.DefaultClient {
		if http.DefaultClient.Transport == nil {
			http.DefaultClient.Transport = &wrapper{http.DefaultTransport}
		} else {
			http.DefaultClient.Transport = WrapRoundTripper(http.DefaultClient.Transport)
		}
	} else {
		if client.Transport == nil {
			client.Transport = &wrapper{http.DefaultTransport}
		} else {
			client.Transport = WrapRoundTripper(client.Transport)
		}
	}
}

// WrapRoundTripper - wrap a round tripper
func WrapRoundTripper(rt http.RoundTripper) http.RoundTripper {
	return &wrapper{rt}
}
