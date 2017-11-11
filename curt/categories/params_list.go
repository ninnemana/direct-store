package categories

import "strconv"

// ListParams define the different ways to retrieve a list of
// categories.
type ListParams struct {
	BrandID *int `json:"branchID"`
}

// Build converts the ListParams into a map.
func (p *ListParams) Build() (map[string]string, error) {
	params := map[string]string{}
	if p == nil {
		return params, nil
	}

	if p.BrandID != nil {
		params["brandID"] = strconv.Itoa(*p.BrandID)
	}

	return params, nil
}
