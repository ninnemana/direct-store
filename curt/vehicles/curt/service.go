package service

import (
	"errors"

	"github.com/ninnemana/direct-store/curt"
)

// Service ...
type Service struct{}

// Lookup ...
func (s *Service) Lookup(v *curt.Vehicle) (interface{}, error) {
	return nil, errors.New("not implemented")
}
