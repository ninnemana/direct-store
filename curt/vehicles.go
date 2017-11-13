package curt

// Vehicles defines the service methods used
// to look up vehicles for the CURT brand.
type Vehicles interface {
	Lookup(Vehicle) (interface{}, error)
}

// Vehicle represents a CURT Vehicle definition.
type Vehicle struct {
	Year  *string `json:"year,omitempty"`
	Make  *string `json:"make,omitempty"`
	Model *string `json:"model,omitempty"`
	Style *string `json:"style,omitempty"`
}

// AriesVehicles defines the service methods used
// to look up vehicles for the ARIES brand.
type AriesVehicles interface {
	GetYears() (interface{}, error)
	GetMakes(year string) (interface{}, error)
	GetModels(year, make string) (interface{}, error)
	GetCategories(year, make, model string) (interface{}, error)
	GetStyles(year, make, model, category string) (interface{}, error)
}

// AriesVehicle represents an ARIES Vehicle definition.
type AriesVehicle struct {
	Year     *string `json:"year,omitempty"`
	Make     *string `json:"make,omitempty"`
	Model    *string `json:"model,omitempty"`
	Style    *string `json:"style,omitempty"`
	Category *string `json:"category,omitempty"`
}

// LuverneVehicles defines the service methods used
// to lookup vehicles for the Luverne brand.
type LuverneVehicles interface {
	GetYears() (interface{}, error)
	GetMakes(year string) (interface{}, error)
	GetModels(year, make string) (interface{}, error)
	GetCategories(year, make, model string) (interface{}, error)
	GetStyles(year, make, model, category string) (interface{}, error)
}

// LuverneVehicle represents a Luverne Vehicle definition.
type LuverneVehicle struct {
	Year     *string `json:"year,omitempty"`
	Make     *string `json:"make,omitempty"`
	Model    *string `json:"model,omitempty"`
	Style    *string `json:"style,omitempty"`
	Category *string `json:"category,omitempty"`
}
