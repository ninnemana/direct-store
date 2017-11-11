package curt

import (
	"net/url"
	"time"
)

type PaginatedProductListing struct{}

type PartResponse struct {
	Parts      []Part `json:"parts"`
	Page       int    `json:"page"`
	TotalPages int    `json:"total_pages"`
}

type Part struct {
	ID         int    `json:"id" xml:"id,attr" bson:"id"`
	PartNumber string `bson:"part_number" json:"part_number" xml:"part_number,attr"`
	// Brand             brand.Brand          `json:"brand" xml:"brand,attr" bson:"brand"`
	Status        int       `json:"status" xml:"status,attr" bson:"status"`
	PriceCode     int       `json:"price_code" xml:"price_code,attr" bson:"price_code"`
	RelatedCount  int       `json:"related_count" xml:"related_count,attr" bson:"related_count"`
	AverageReview float64   `json:"average_review" xml:"average_review,attr" bson:"average_review"`
	DateModified  time.Time `json:"date_modified" xml:"date_modified,attr" bson:"date_modified"`
	DateAdded     time.Time `json:"date_added" xml:"date_added,attr" bson:"date_added"`
	ShortDesc     string    `json:"short_description" xml:"short_description,attr" bson:"short_description"`
	InstallSheet  *url.URL  `json:"install_sheet" xml:"install_sheet" bson:"install_sheet"`
	// Attributes        []Attribute          `json:"attributes" xml:"attributes" bson:"attributes"`
	// AcesVehicles      []AcesVehicle        `bson:"aces_vehicles" json:"aces_vehicles" xml:"aces_vehicles"`
	VehicleAttributes []string `json:"vehicle_atttributes" xml:"vehicle_attributes" bson:"vehicle_attributes"`
	// Vehicles          []VehicleApplication `json:"vehicle_applications,omitempty" xml:"vehicle_applications,omitempty" bson:"vehicle_applications"`
	// LuverneVehicles   []LuverneApplication `json:"luverne_applications,omitempty" xml:"luverne_applications,omitempty" bson:"luverne_applications"`
	// Content           []Content            `json:"content" xml:"content" bson:"content"`
	// Pricing           []Price              `json:"pricing" xml:"pricing" bson:"pricing"`
	// Reviews           []Review             `json:"reviews" xml:"reviews" bson:"reviews"`
	// Images            []Image              `json:"images" xml:"images" bson:"images"`
	Related    []int `json:"related" xml:"related" bson:"related" bson:"related"`
	ReplacedBy int   `bson:"replaced_by" json:"replaced_by,omitempty" xml:"replaced_by,omitempty"`
	// Categories        []Category           `json:"categories" xml:"categories" bson:"categories"`
	// Videos            []video.Video        `json:"videos" xml:"videos" bson:"videos"`
	// Packages          []Package            `json:"packages" xml:"packages" bson:"packages"`
	// Customer          CustomerPart         `json:"customer,omitempty" xml:"customer,omitempty" bson:"v"`
	// Class             Class                `json:"class,omitempty" xml:"class,omitempty" bson:"class"`
	Featured       bool `json:"featured,omitempty" xml:"featured,omitempty" bson:"featured"`
	AcesPartTypeID int  `json:"acesPartTypeId,omitempty" xml:"acesPartTypeId,omitempty" bson:"acesPartTypeId"`
	// Inventory         PartInventory        `json:"inventory,omitempty" xml:"inventory,omitempty" bson:"inventory"`
	UPC             string `json:"upc,omitempty" xml:"upc,omitempty" bson:"upc"`
	Layer           string `json:"iconLayer" xml:"iconLayer" bson:"iconLayer"`
	MappedToVehicle bool   `json:"mappedToVehicle" xml:"mappedToVehicle" bson:"mappedToVehicle,omitempty"`
	// ComplexPart       *ComplexPart         `bson:"complex_part" json:"complex_part,omitempty" xml:"complex_part,omitempty"`
}
