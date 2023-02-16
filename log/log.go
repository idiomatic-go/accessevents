package log

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

const (
	errorName     = "error"
	errorNilEntry = "access data entry is nil"
	errorEmptyFmt = "%v log entries are empty"
)

var ingressOperators []accessdata.Operator
var egressOperators []accessdata.Operator

// InitIngressOperators - allows configuration of access log attributes for ingress traffic
func InitIngressOperators(config []accessdata.Operator) error {
	var err error
	ingressOperators, err = accessdata.InitOperators(config)
	return err
}

// InitEgressOperators - allows configuration of access log attributes for egress traffic
func InitEgressOperators(config []accessdata.Operator) error {
	var err error
	egressOperators, err = accessdata.InitOperators(config)
	return err
}

// Log - templated function handling writing the access log entry utilizing the OutputHandler and Formatter
func Log[O OutputHandler, F accessdata.Formatter](entry *accessdata.Entry) {
	var o O
	var f F
	if entry == nil {
		o.Write([]accessdata.Operator{{errorName, errorNilEntry}}, accessdata.NewEntry(), f)
		return
	}
	if entry.IsIngress() {
		if !opt.ingress {
			return
		}
		if len(ingressOperators) == 0 {
			o.Write(emptyOperators(entry), accessdata.NewEntry(), f)
			return
		}
		o.Write(ingressOperators, entry, f)
	} else {
		if !opt.egress {
			return
		}
		if len(egressOperators) == 0 {
			o.Write(emptyOperators(entry), accessdata.NewEntry(), f)
			return
		}
		o.Write(egressOperators, entry, f)
	}
}

func emptyOperators(entry *accessdata.Entry) []accessdata.Operator {
	return []accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}
}
