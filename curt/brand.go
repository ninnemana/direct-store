package curt

import "net/url"

// BrandInteractor defines the available service methods for
// the Brand API.
type BrandInteractor interface {
	// List returns a list of brands
	List() ([]Brand, error)
	// Get returns a specific brand by identifier.
	Get(id int) (*Brand, error)
}

// Brand defines the object structure for a manufacturer brand.
type Brand struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Code          string    `json:"code"`
	Logo          *url.URL  `json:"logo"`
	LogoAlternate *url.URL  `json:"logo_alternate"`
	FormalName    string    `json:"formal_name"`
	LongName      string    `json:"long_name"`
	PrimaryColor  string    `json:"primary_color"`
	AutocareID    string    `json:"autocareId"`
	Websites      []Website `json:"websites"`
}

// Website defines the object definition for a Brand's Websites.
type Website struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	URL         *url.URL `json:"url"`
	BrandID     int      `json:"brand_id"`
}
