package middleware

import (
	"github.com/idiomatic-go/accessevents/data"
	"github.com/idiomatic-go/accessevents/log"
)

// SetLogFn - allows setting an application configured logging function
func SetLogFn(fn data.Accessor) {
	if fn != nil {
		logFn = fn
	}
}

var logFn = defaultLogFn

var defaultLogFn data.Accessor = func(e *data.Entry) {
	log.Write[log.LogOutputHandler, data.JsonFormatter](e)
}
