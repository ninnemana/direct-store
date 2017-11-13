package brands

import (
	"context"
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

// Service implements the BrandInteractor.
type Service struct {
	Client client.Provider
	Trace  *trace.Span
	Log    *logging.Logger
}

// New initiates a version of the curt.BrandInteractor as defined by the provided configuration.
func New(cfg Config) (*Service, error) {

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

func (s *Service) List() ([]curt.Brand, error) {
	res := []curt.Brand{}
	options := client.Options{
		Method:   http.MethodGet,
		Endpoint: "/brands",
		Result:   &res,
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

// Get implements the BrandInteractor Get requirement.
func (s *Service) Get(id int) (*curt.Brand, error) {
	res := curt.Brand{}
	options := client.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/brands/%d", id),
		Result:   &res,
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
