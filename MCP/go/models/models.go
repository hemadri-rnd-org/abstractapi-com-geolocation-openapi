package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Inlineresponse200 represents the Inlineresponse200 schema from the OpenAPI specification
type Inlineresponse200 struct {
	City string `json:"city,omitempty"`
	Security map[string]interface{} `json:"security,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Postal_code string `json:"postal_code,omitempty"`
	City_geoname_id int `json:"city_geoname_id,omitempty"`
	Continent string `json:"continent,omitempty"`
	Continent_code string `json:"continent_code,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Country_code string `json:"country_code,omitempty"`
	Region_geoname_id int `json:"region_geoname_id,omitempty"`
	Region_iso_code string `json:"region_iso_code,omitempty"`
	Connection map[string]interface{} `json:"connection,omitempty"`
	Flag map[string]interface{} `json:"flag,omitempty"`
	Continent_geoname_id int `json:"continent_geoname_id,omitempty"`
	Country_geoname_id int `json:"country_geoname_id,omitempty"`
	Country string `json:"country,omitempty"`
	Country_is_eu bool `json:"country_is_eu,omitempty"`
	Currency map[string]interface{} `json:"currency,omitempty"`
	Region string `json:"region,omitempty"`
	Ip_address string `json:"ip_address,omitempty"`
	Timezone map[string]interface{} `json:"timezone,omitempty"`
}
