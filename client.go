package store

type ClientProvider interface {
	Get() (interface{}, error)
	Post() (interface{}, error)
	Put() (interface{}, error)
	Patch() (interface{}, error)
	Delete() (interface{}, error)
	Options() (interface{}, error)
}

type Client struct{}
