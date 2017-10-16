package curt

import (
	"net/url"
	"time"
)

type Categories interface {
	List() (interface{}, error)
	Get() (interface{}, error)
}

type Category struct {
	ID       int        `json:"id"`
	ParentID int        `json:"parent_id"`
	Children []Category `json:"children"`

	Sort            int                      `json:"sort"`
	DateAdded       time.Time                `json:"date_added"`
	Title           string                   `json:"title"`
	ShortDesc       string                   `json:"short_description"`
	LongDesc        string                   `json:"long_description"`
	ColorCode       string                   `json:"color_code"`
	FontCode        string                   `json:"font_code"`
	Image           *url.URL                 `json:"image"`
	Icon            *url.URL                 `json:"icon"`
	IsLifestyle     bool                     `json:"lifestyle"`
	VehicleSpecific bool                     `json:"vehicle_specific"`
	VehicleRequired bool                     `json:"vehicle_required"`
	MetaTitle       string                   `json:"meta_title"`
	MetaDescription string                   `json:"meta_description"`
	MetaKeywords    string                   `json:"meta_keywords"`
	Content         []Content                `json:"content"`
	Videos          []Video                  `json:"videos"`
	PartIDs         []int                    `json:"part_ids"`
	Brand           Brand                    `json:"brand"`
	ProductListing  *PaginatedProductListing `json:"product_listing"`
	PDFpath         *url.URL                 `json:"pdf_path"`
	XLSpath         *url.URL                 `json:"xls_path"`
}
