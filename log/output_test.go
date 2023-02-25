package log

import (
	"fmt"
	"github.com/idiomatic-go/accessevents/data"
)

func ExampleOutputHandler() {
	fmt.Printf("test: Output[NilOutputHandler,data.TextFormatter](nil,nil)\n")
	logTest[NilOutputHandler, data.TextFormatter](nil, nil)

	fmt.Printf("test: Output[DebugOutputHandler,data.JsonFormatter](operators,data)\n")
	ops := []data.Operator{{"error", "message"}}
	logTest[DebugOutputHandler, data.JsonFormatter](ops, data.NewEmptyEntry())

	fmt.Printf("test: Output[TestOutputHandler,data.JsonFormatter](nil,nil)\n")
	logTest[TestOutputHandler, data.JsonFormatter](nil, nil)

	fmt.Printf("test: Output[TestOutputHandler,data.JsonFormatter](ops,data)\n")
	logTest[TestOutputHandler, data.JsonFormatter](ops, data.NewEmptyEntry())

	fmt.Printf("test: Output[LogOutputHandler,data.JsonFormatter](ops,data)\n")
	logTest[LogOutputHandler, data.JsonFormatter](ops, data.NewEmptyEntry())

	//Output:
	//test: Output[NilOutputHandler,data.TextFormatter](nil,nil)
	//test: Output[DebugOutputHandler,data.JsonFormatter](operators,data)
	//{"error":"message"}
	//test: Output[TestOutputHandler,data.JsonFormatter](nil,nil)
	//test: Write() -> [{}]
	//test: Output[TestOutputHandler,data.JsonFormatter](ops,data)
	//test: Write() -> [{"error":"message"}]
	//test: Output[LogOutputHandler,data.JsonFormatter](ops,data)

}

func logTest[O OutputHandler, F data.Formatter](items []data.Operator, data *data.Entry) {
	var o O
	var f F
	o.Write(items, data, f)
}
