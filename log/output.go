package log

import (
	"fmt"
	"github.com/idiomatic-go/accessevents/data"
	"log"
)

// OutputHandler - template parameter for log output
type OutputHandler interface {
	Write(items []data.Operator, data *data.Entry, formatter data.Formatter)
}

// NilOutputHandler - no output
type NilOutputHandler struct{}

func (NilOutputHandler) Write(_ []data.Operator, _ *data.Entry, _ data.Formatter) {
}

// DebugOutputHandler - output to stdio
type DebugOutputHandler struct{}

func (DebugOutputHandler) Write(items []data.Operator, data *data.Entry, formatter data.Formatter) {
	fmt.Printf("%v\n", formatter.Format(items, data))
}

// TestOutputHandler - special use case of DebugOutputHandler for testing examples
type TestOutputHandler struct{}

func (TestOutputHandler) Write(items []data.Operator, data *data.Entry, formatter data.Formatter) {
	fmt.Printf("test: Write() -> [%v]\n", formatter.Format(items, data))
}

// LogOutputHandler - output to log
type LogOutputHandler struct{}

func (LogOutputHandler) Write(items []data.Operator, data *data.Entry, formatter data.Formatter) {
	log.Println(formatter.Format(items, data))
}
