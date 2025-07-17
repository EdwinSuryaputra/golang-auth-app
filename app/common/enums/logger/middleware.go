package loggerenum

type Middleware string

var (
	APIRequest  Middleware = "API-REQUEST"
	APIResponse Middleware = "API-RESPONSE"
)

func (m Middleware) ToString() string {
	return string(m)
}
