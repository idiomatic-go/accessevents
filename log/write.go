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
	var operators []data.Operator
	switch entry.Traffic {
	case data.IngressTraffic, data.PingTraffic:
		if entry.Traffic == data.IngressTraffic && !opt.ingress {
			return
		}
		if entry.Traffic == data.PingTraffic && !opt.ping {
			return
		}
		operators = ingressOperators
	case data.EgressTraffic:
		if !opt.egress {
			return
		}
		operators = egressOperators
	}
	if len(operators) == 0 {
		operators = emptyOperators(entry)
	}
	o.Write(operators, entry, f)
}

func emptyOperators(entry *data.Entry) []data.Operator {
	return []data.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}
}
