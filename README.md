# accessevents

Motif was inspired by the challenges of developing Go applications. Determining the patterns, or motifs, that need to be employed, is critical for writing clear idiomatic Go code. This YouTube video [Edward Muller - Go Anti-Patterns][emuller], does an excellent job of framing idiomatic go. 
[Robert Griesemer - The Evolution of Go][rgriesemer] also presents an important analogy between Go packages and LEGO® bricks. Reviewing the Go standard
library packaging structure provides a blueprint for an application architecture, and underscores how essential package design is for idiomatic Go. 

With the release of Go generics, a new paradigm has emerged: [templates][tutorialspoint]. Templates are not new, having been available in  C++ since 1991, and have become a standard through the work of teams like [boost][boost]. I prefer the term templates over generics, as templates are a paradigm, and generics connotes a class of implementations. What follows is a description of the packages in Motif, highlighting specific patterns and template implementations.  


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
The log function is a templated function, allowing for selection of output and formatting:
~~~
func Log[O OutputHandler, F accessdata.Formatter](entry *accessdata.Entry) {
    // implementation details
}
~~~

## middleware

[Middleware][middlewarepkg] provides functionality to invoke logging on ingress and egress traffic.


[datapkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accessdata>
[extractpkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accesslog>
[logpkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accesslog>
[middlewarepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accesslog>
