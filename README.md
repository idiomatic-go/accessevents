# accessevents

Accessevents provides 4 packages that work with access events, access logging, and access event extract.

## data 

[Data][datapkg] provides the Entry type, which contains all of the data needed for access logging. Also provided are functions and types that define command operators which 
allow the extraction and formatting of Entry data. The formatting of Entry data is implemented as a template parameter: 
~~~
// Formatter - template parameter for formatting
type Formatter interface {
	Format(items []Operator, data *Entry) string
}
~~~
Configurable items, specific to a package, are defined in an options.go file.

## extract

[Extract][extractpkg] provides functionality to initialize and process access log extract.


## log

[Log][logpkg] encompasses access logging functionality. Seperate operators, and runtime initialization of those operators, are provided for ingress and egress traffic. An output template parameter allows redirection of the access logging: 
~~~
// OutputHandler - template parameter for log output
type OutputHandler interface {
	Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter)
}
~~~
The log.Write function is a templated function, allowing for selection of output and formatting:
~~~
func Write[O OutputHandler, F accessdata.Formatter](entry *accessdata.Entry) {
    // implementation details
}
~~~

## middleware

[Middleware][middlewarepkg] provides functionality to invoke logging on ingress and egress traffic. An http.Handler implements ingress logging, while an
http.RoundTrip interface logs egress.


[datapkg]: <https://pkg.go.dev/github.com/idiomatic-go/accessevents/data>
[extractpkg]: <https://pkg.go.dev/github.com/idiomatic-go/accessevents/extract>
[logpkg]: <https://pkg.go.dev/github.com/idiomatic-go/accessevents/log>
[middlewarepkg]: <https://pkg.go.dev/github.com/idiomatic-go/accessevents/middleware>
