package log

import (
	"fmt"
	"github.com/idiomatic-go/accessevents/data"
	"net/http"
	"reflect"
	"time"
)

func ExampleLog_Error() {
	start := time.Now()

	Write[TestOutputHandler, data.TextFormatter](nil)
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start, time.Since(start), map[string]string{data.ActName: "egress-route"}, nil, nil, ""))

	//Output:
	//test: Write() -> [access data entry is nil]
	//test: Write() -> [{"error":"egress log entries are empty"}]

}

func ExampleLog_Origin() {
	name := "ingress-origin-route"
	start := time.Now()

	data.SetOrigin(data.Origin{Region: "us-west", Zone: "dfw", Service: "test-service", InstanceId: "123456-7890-1234"})
	err := InitIngressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Value: data.DurationOperator, Name: "duration_ms"},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}, {Value: data.OriginRegionOperator}, {Value: data.OriginZoneOperator}, {Value: data.OriginServiceOperator}, {Value: data.OriginInstanceIdOperator},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	entry := data.NewHttpIngressEntry(start1, time.Since(start), map[string]string{data.ActName: name}, nil, 0, 0, "")
	Write[TestOutputHandler, data.JsonFormatter](entry)
	Write[TestOutputHandler, data.TextFormatter](entry)

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"ingress","route_name":"ingress-origin-route","region":"us-west","zone":"dfw","service":"test-service","instance_id":"123456-7890-1234"}]
	//test: Write() -> [0001-01-01 00:00:00.000000,0,ingress,ingress-origin-route,us-west,dfw,test-service,123456-7890-1234]

}

func ExampleLog_Ping() {
	name := "ingress-ping-route"
	url := "https://www.google.com/search"

	req, _ := http.NewRequest("", url, nil)
	data.SetPingRoutes([]data.PingRoute{{Traffic: "ingress", Pattern: "/search"}})
	start := time.Now()
	err := InitIngressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Value: data.DurationOperator, Name: "duration_ms"},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	entry := data.NewHttpIngressEntry(start1, time.Since(start), map[string]string{data.ActName: name}, req, 0, 0, "")
	Write[TestOutputHandler, data.JsonFormatter](entry)

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"ping","route_name":"ingress-ping-route"}]

}

func ExampleLog_Timeout() {
	start := time.Now()

	err := InitEgressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Name: "duration_ms", Value: data.DurationOperator},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}, {Value: data.TimeoutDurationOperator}, {Name: "static", Value: "value"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start1, time.Since(start), map[string]string{data.ActName: "handler-route", data.TimeoutName: "5000"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"egress","route_name":"handler-route","timeout_ms":5000,"static":"value"}]

}

func ExampleLog_RateLimiter_500() {
	start := time.Now()

	err := InitEgressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Name: "duration", Value: data.DurationOperator},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}, {Value: data.RateLimitOperator}, {Value: data.RateBurstOperator}, {Name: "static2", Value: "value2"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start1, time.Since(start), map[string]string{data.ActName: "handler-route", data.RateLimitName: "500", data.RateBurstName: "10"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration":0,"traffic":"egress","route_name":"handler-route","rate_limit":500,"rate_burst":10,"static2":"value2"}]

}

func ExampleLog_Failover() {
	start := time.Now()

	err := InitEgressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Name: "duration", Value: data.DurationOperator},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}, {Value: data.FailoverOperator}, {Name: "static2", Value: "value2"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start1, time.Since(start), map[string]string{data.ActName: "handler-route", data.FailoverName: "true"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration":0,"traffic":"egress","route_name":"handler-route","failover":true,"static2":"value2"}]

}

func ExampleLog_Retry() {
	start := time.Now()

	err := InitEgressOperators([]data.Operator{{Value: data.StartTimeOperator}, {Value: data.DurationOperator, Name: "duration_ms"},
		{Value: data.TrafficOperator}, {Value: data.RouteNameOperator}, {Value: data.RetryOperator},
		{Value: data.RetryRateLimitOperator}, {Value: data.RetryRateBurstOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start1, time.Since(start), map[string]string{data.ActName: "handler-route", data.RetryName: "true", data.RetryRateLimitName: "123", data.RetryRateBurstName: "67"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"egress","route_name":"handler-route","retry":true,"retry_rate_limit":123,"retry_rate_burst":67}]

}

func ExampleLog_Request() {
	req, _ := http.NewRequest("", "www.google.com/search/documents", nil)
	req.Header.Add("customer", "Ted's Bait & Tackle")

	var start time.Time
	err := InitEgressOperators([]data.Operator{{Value: data.RequestProtocolOperator}, {Value: data.RequestMethodOperator}, {Value: data.RequestUrlOperator},
		{Value: data.RequestPathOperator}, {Value: data.RequestHostOperator}, {Value: "%REQ(customer)%"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start, time.Since(start), map[string]string{data.ActName: "handler-route"}, nil, nil, ""))
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start, time.Since(start), map[string]string{data.ActName: "handler-route"}, req, nil, ""))

	//Output:
	//test: Write() -> [{"protocol":null,"method":null,"url":null,"path":null,"host":null,"customer":null}]
	//test: Write() -> [{"protocol":"HTTP/1.1","method":"GET","url":"www.google.com/search/documents","path":"www.google.com/search/documents","host":null,"customer":"Ted's Bait & Tackle"}]

}

func ExampleLog_Response() {
	resp := &http.Response{StatusCode: 404, ContentLength: 1234}

	err := InitEgressOperators([]data.Operator{{Value: data.ResponseStatusCodeOperator}, {Value: data.ResponseBytesReceivedOperator}, {Value: data.StatusFlagsOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start time.Time
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start, time.Since(start), map[string]string{data.ActName: "handler-route"}, nil, nil, "UT"))
	Write[TestOutputHandler, data.JsonFormatter](data.NewHttpEgressEntry(start, time.Since(start), map[string]string{data.ActName: "handler-route"}, nil, resp, "UT"))

	//Output:
	//test: Write() -> [{"status_code":0,"bytes_received":0,"status_flags":"UT"}]
	//test: Write() -> [{"status_code":404,"bytes_received":1234,"status_flags":"UT"}]

}

func _Example_Log_State() {
	t := time.Duration(time.Millisecond * 500)
	i := reflect.TypeOf(t)
	a := any(t)

	fmt.Printf("test 1 -> %v\n", a)

	fmt.Printf("test 2 -> %v\n", i)

	//Output:
	//fail
}
