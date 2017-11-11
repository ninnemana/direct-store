package client

// Options defines per-request parameters for
// the client to use on execution.
type Options struct {
	Method      string
	Endpoint    string
	Result      interface{}
	Body        interface{}
	QueryString Params
}

// Params specifies the requirement for parsing a series
// of parameters into a Query String.
type Params interface {
	Build() (map[string]string, error)
}
