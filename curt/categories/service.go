package categories

import (
	"context"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/trace"
	"github.com/ninnemana/direct-store/curt"
	"github.com/ninnemana/direct-store/curt/client"
	"github.com/pkg/errors"
)

const (
	logPrefix = "categories"
)

// Config defines the available configuration parameters for the
// category service to be initialized.
type Config struct {
	Client  client.Provider
	Log     interface{}
	Span    *trace.Span
	Context context.Context
}

// Service implements the CategoryInteractor.
type Service struct {
	Client client.Provider
	Trace  *trace.Span
	Log    *logging.Logger
}

// New initiates a version of the categories.Service as defined by the provided configuration.
func New(cfg Config) (curt.Categories, error) {
	if cfg.Client == nil {
		return nil, errors.New("client.Client was not provided")
	}

	s := Service{
		Client: cfg.Client,
		Trace:  cfg.Span,
	}

	switch cfg.Log.(type) {
	case nil:
		return nil, errors.New("logging.Logger was not provided")
	case logging.Logger:
		lg := cfg.Log.(logging.Logger)
		s.Log = &lg
	case *logging.Logger:
		s.Log = cfg.Log.(*logging.Logger)
	case *logging.Client:
		s.Log = cfg.Log.(*logging.Client).Logger(logPrefix)
	}

	return s, nil
}

// List implements the CategoryInteractors List requirement.
func (s Service) List() (interface{}, error) {
	res := []curt.Category{}
	options := client.Options{
		Endpoint: "/category",
		Result:   &res,
	}

	err := s.Client.Get(options)
	if err != nil {
		s.Log.Log(logging.Entry{
			Severity: logging.Error,
			Payload: map[string]interface{}{
				"error":   err.Error(),
				"options": options,
			},
		})
		return nil, err
	}

	return res, nil
}

// Get implements the CategoryInteractors Get requirement.
func (s Service) Get() (interface{}, error) {
	return nil, errors.New("not implemented")
}
