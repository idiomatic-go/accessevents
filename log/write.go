package log

import (
	"fmt"
	"github.com/idiomatic-go/accessevents/data"
)

const (
	errorName     = "error"
	errorNilEntry = "access data entry is nil"
	errorEmptyFmt = "%v log entries are empty"
)

var ingressOperators []data.Operator
var egressOperators []data.Operator

// Write - templated function handling writing the access data utilizing the OutputHandler and Formatter
func Write[O OutputHandler, F data.Formatter](entry *data.Entry) {
	var o O
	var f F
	if entry == nil {
		o.Write([]data.Operator{{errorName, errorNilEntry}}, data.NewEntry(), f)
		return
	}
	if entry.IsIngress() {
		if !opt.ingress {
			return
		}
		if len(ingressOperators) == 0 {
			o.Write(emptyOperators(entry), data.NewEntry(), f)
			return
		}
		o.Write(ingressOperators, entry, f)
	} else {
		if !opt.egress {
			return
		}
		if len(egressOperators) == 0 {
			o.Write(emptyOperators(entry), data.NewEntry(), f)
			return
		}
		o.Write(egressOperators, entry, f)
	}
}

func emptyOperators(entry *data.Entry) []data.Operator {
	return []data.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}
}
