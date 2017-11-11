package categories

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

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
func New(cfg Config) (*Service, error) {
	gob.Register(&curt.Category{})
	gob.Register([]curt.Category{})

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

	return &s, nil
}

// List implements the CategoryInteractors List requirement.
func (s Service) List(params *ListParams) ([]curt.Category, error) {
	res := []curt.Category{}
	options := client.Options{
		Method:      http.MethodGet,
		Endpoint:    "/category",
		Result:      &res,
		QueryString: params,
	}

	s.Log.Log(logging.Entry{
		Severity: logging.Debug,
		Payload: map[string]interface{}{
			"options": options,
		},
	})

	err := s.Client.Do(options)
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
func (s Service) Get(params *GetParams) (*curt.Category, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	res := curt.Category{}
	options := client.Options{
		Method:      http.MethodGet,
		Endpoint:    fmt.Sprintf("/category/%d", *params.ID),
		Result:      &res,
		QueryString: params,
	}

	s.Log.Log(logging.Entry{
		Severity: logging.Debug,
		Payload: map[string]interface{}{
			"options": options,
		},
	})

	err := s.Client.Do(options)
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

	return &res, nil
}

// GetParts implements the CategoryInteractors GetParts requirement.
func (s Service) GetParts(params *GetParams) (*curt.PartResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	res := curt.PartResponse{}
	options := client.Options{
		Method:      http.MethodGet,
		Endpoint:    fmt.Sprintf("/category/%d/parts", *params.ID),
		Result:      &res,
		QueryString: params,
	}

	s.Log.Log(logging.Entry{
		Severity: logging.Debug,
		Payload: map[string]interface{}{
			"options": options,
		},
	})

	err := s.Client.Do(options)
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

	return &res, nil
}
