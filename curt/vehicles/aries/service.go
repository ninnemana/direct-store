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
func (s *Service) GetModels(year, make string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetCategories ...
func (s *Service) GetCategories(year, make, model string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// GetStyles ...
func (s *Service) GetStyles(year, make, model, category string) (interface{}, error) {
	return nil, errors.New("not implemented")
}
