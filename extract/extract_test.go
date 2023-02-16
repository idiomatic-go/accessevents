package extract

/*
func setTestErrorHandler() {
	opt.handler = func(err error) {
		fmt.Printf("test: extract(logd) -> [err:%v]\n", err)
	}
}

*/

/*
func Example_Initialize_Url() {
	status := Initialize[template.DebugError]("", nil)
	fmt.Printf("test: initialize(\"\") -> [%v] [url:%v]\n", status, url)

	status = Initialize[template.DebugError]("test", nil)
	fmt.Printf("test: initialize(\"\") -> [%v] [url:%v]\n", status, url)

	//Output:
	//[[] github.com/idiomatic-go/resiliency/extract/initialize [invalid argument: uri is empty]]
	//test: initialize("") -> [Internal] [url:]
	//test: initialize("") -> [OK] [url:test]

}

func Example_Handler_NotProcessed() {
	url = "http://localhost:8080/accesslog"

	status := handler(nil)
	fmt.Printf("test: handler(nil) -> [%v]\n", status)

	req, _ := http.NewRequest("post", "http://localhost:8080/accesslog", nil)
	data := new(accessdata.Entry)
	data.AddRequest(req)
	status = handler(data)
	fmt.Printf("test: handler(data) -> [%v]\n", status)

	//Output:
	//[[] github.com/idiomatic-go/resiliency/extract/do [invalid argument: access log data is nil]]
	//test: handler(nil) -> [false]
	//test: handler(data) -> [false]

}

func Example_Handler_ConnectFailure() {
	url = "http://localhost:8080/accesslog"

	req, _ := http.NewRequest("post", "localhost:8081/accesslog", nil)
	data := new(accessdata.Entry)
	data.AddRequest(req)
	status := handler(data)
	fmt.Printf("test: handler(data) -> [%v]\n", status)

	//Output:
	//[[] github.com/idiomatic-go/resiliency/extract/do [Post "http://localhost:8080/accesslog": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]]
	//test: handler(data) -> [false]

}

func _Example_Handler_Processed() {
	// Override the message handler
	handler = func(l *accessdata.Entry) bool {
		fmt.Printf("test: handler(logd) -> [%v]\n", accessdata.WriteJson(operators, l))
		return true
	}

	status := Initialize[template.DebugError]("http://localhost:8080/accesslog", nil)
	fmt.Printf("test: initialize() -> [%v]\n", status)

	//r0, _ := route.NewRoute("route-data-0")
	//r1, _ := route.NewRoute("route-data-1")
	//r2, _ := route.NewRoute("route-data-2")
	//r3, _ := route.NewRoute("route-data-3")

	data0 := accessdata.Entry{Origin: &accessdata.Origin{Region: "region-1"}, ActState: map[string]string{accessdata.ActName: "route-data-0"}}
	data1 := accessdata.Entry{Origin: &accessdata.Origin{Region: "region-2"}, ActState: map[string]string{accessdata.ActName: "route-data-1"}}
	data2 := accessdata.Entry{Origin: &accessdata.Origin{Region: "region-3"}, ActState: map[string]string{accessdata.ActName: "route-data-2"}}
	data3 := accessdata.Entry{Origin: &accessdata.Origin{Region: "region-4"}, ActState: map[string]string{accessdata.ActName: "route-data-3"}}
	extract(&data0)
	extract(&data1)
	extract(&data2)
	extract(&data3)
	time.Sleep(time.Second * 2)
	Shutdown()

	//Output:
	//test: initialize() -> [0 Successful]
	//test: handler(logd) -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":null,"route_name":"route-data-0","region":"region-1","zone":null,"sub_zone":null,"service":null,"instance_id":null,"method":null,"host":null,"path":null,"protocol":null,"request_id":null,"forwarded":null,"status_code":"0","status_flags":null,"timeout_ms":null,"rate_limit":null,"rate_burst":null,"retry":null,"retry_rate_limit":null,"retry_rate_burst":null,"failover":null}]
	//test: handler(logd) -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":null,"route_name":"route-data-2","region":"region-3","zone":null,"sub_zone":null,"service":null,"instance_id":null,"method":null,"host":null,"path":null,"protocol":null,"request_id":null,"forwarded":null,"status_code":"0","status_flags":null,"timeout_ms":null,"rate_limit":null,"rate_burst":null,"retry":null,"retry_rate_limit":null,"retry_rate_burst":null,"failover":null}]
	//test: handler(logd) -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":null,"route_name":"route-data-3","region":"region-4","zone":null,"sub_zone":null,"service":null,"instance_id":null,"method":null,"host":null,"path":null,"protocol":null,"request_id":null,"forwarded":null,"status_code":"0","status_flags":null,"timeout_ms":null,"rate_limit":null,"rate_burst":null,"retry":null,"retry_rate_limit":null,"retry_rate_burst":null,"failover":null}]
	//test: handler(logd) -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":null,"route_name":"route-data-1","region":"region-2","zone":null,"sub_zone":null,"service":null,"instance_id":null,"method":null,"host":null,"path":null,"protocol":null,"request_id":null,"forwarded":null,"status_code":"0","status_flags":null,"timeout_ms":null,"rate_limit":null,"rate_burst":null,"retry":null,"retry_rate_limit":null,"retry_rate_burst":null,"failover":null}]

}


*/
