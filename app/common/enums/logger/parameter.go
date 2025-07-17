package loggerenum

type Parameter string

var (
	StatusCode   Parameter = "statusCode"
	RequestId    Parameter = "requestId"
	Path         Parameter = "path"
	Method       Parameter = "method"
	Ip           Parameter = "ip"
	Panic        Parameter = "panic"
	Stacktrace   Parameter = "stacktrace"
	QueryParams  Parameter = "queryParams"
	RequestBody  Parameter = "requestBody"
	ResponseBody Parameter = "responseBody"
	Latency      Parameter = "latency"
)

func (p Parameter) ToString() string {
	return string(p)
}
