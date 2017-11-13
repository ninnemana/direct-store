package service

import "github.com/pkg/errors"

// Service ...
type Service struct{}

// GetYears ...
func (s *Service) GetYears() (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetMakes ...
func (s *Service) GetMakes(year string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetModels ...
func (s *Service) GetModels(year string, make string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetCategories ...
func (s *Service) GetCategories(year string, make string, model string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetStyles ...
func (s *Service) GetStyles(year string, make string, model string, category string) (interface{}, error) {
	panic("not implemented")
}
