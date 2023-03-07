package middleware

import (
	"fmt"
	"github.com/idiomatic-go/accessevents/data"
	"github.com/idiomatic-go/accessevents/log"
	"net/http"
)

var (
	isEnabled    = false
	googleUrl    = "https://www.google.com/search?q=test"
	instagramUrl = "https://www.instagram.com"
	config       = []data.Operator{
		{Value: data.StartTimeOperator},
		{Value: data.DurationOperator},
		{Value: data.TrafficOperator},

		{Value: data.RequestMethodOperator},
		{Value: data.RequestHostOperator},
		{Value: data.RequestPathOperator},
		{Value: data.RequestProtocolOperator},

		{Value: data.ResponseStatusCodeOperator},
		{Value: data.StatusFlagsOperator},
		{Value: data.ResponseBytesSentOperator},
	}
)

func init() {
	log.InitEgressOperators(config)

}

func Example_No_Wrapper() {
	req, _ := http.NewRequest("GET", googleUrl, nil)

	// Testing - check for a nil wrapper or round tripper
	w := wrapper{}
	resp, err := w.RoundTrip(req)
	fmt.Printf("test: RoundTrip(wrapper:nil) -> [resp:%v] [err:%v]\n", resp, err)

	// Testing - no wrapper, calling Google search
	resp, err = http.DefaultClient.Do(req)
	fmt.Printf("test: RoundTrip(handler:false) -> [status_code:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: RoundTrip(wrapper:nil) -> [resp:<nil>] [err:invalid handler round tripper configuration : http.RoundTripper is nil]
	//test: RoundTrip(handler:false) -> [status_code:200] [err:<nil>]

}

func Example_Default() {
	req, _ := http.NewRequest("GET", instagramUrl, nil)

	if !isEnabled {
		isEnabled = true
		WrapTransport(nil)
	}
	resp, err := http.DefaultClient.Do(req)
	fmt.Printf("test: RoundTrip(handler:true) -> [status_code:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: RoundTrip(handler:true) -> [status_code:200] [err:<nil>]

}
