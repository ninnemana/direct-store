package categories

import (
	"strconv"

	"github.com/pkg/errors"
)

// GetParams define the different ways to retrieve a particular category.
type GetParams struct {
	ID    *int `json:"id"`    // Category Identifier
	Page  *int `json:"page"`  // page of product results to include
	Count *int `json:"count"` // number of product results to include
}

// Build converts the ListParams into a map.
func (p *GetParams) Build() (map[string]string, error) {
	params := map[string]string{}
	if p == nil {
		return params, nil
	}

	if p.ID != nil {
		params["id"] = strconv.Itoa(*p.ID)
	}

	params["count"] = "1"

	return params, nil
}

// Validate checks the get parameters to make sure they are properly configured
// to run a query on.
func (p *GetParams) Validate() error {
	if p == nil {
		return errors.New("query parameters cannot be nil")
	}

	if p.ID == nil {
		return errors.New("category identifier cannot be nil")
	}

	return nil
}
