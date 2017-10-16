package client

// Options defines per-request parameters for
// the client to use on execution.
type Options struct {
	Endpoint    string
	Result      interface{}
	Body        interface{}
	QueryString map[string]string
}
