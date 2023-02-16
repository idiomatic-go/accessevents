package log

import (
	"encoding/json"
	"errors"
	"github.com/idiomatic-go/accessevents/data"
)

// CreateEgressOperators - provides creation of egress operators from a []byte
func CreateEgressOperators(buf []byte) error {
	operators, err := ReadOperators(buf)
	if err != nil {
		return err
	}
	return InitEgressOperators(operators)
}

// CreateIngressOperators - provides creation of ingress operators from a []byte
func CreateIngressOperators(buf []byte) error {
	operators, err := ReadOperators(buf)
	if err != nil {
		return err
	}
	return InitIngressOperators(operators)
}

// ReadOperators - read the operators from a []byte
func ReadOperators(buf []byte) ([]data.Operator, error) {
	var operators []data.Operator

	if buf == nil {
		return nil, errors.New("invalid argument: buffer is nil")
	}
	err1 := json.Unmarshal(buf, &operators)
	return operators, err1
}
