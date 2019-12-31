package kanka

// Location provides simple data about a location within the campaign.
// For more information, visit: https://kanka.io/en-US/docs/1.0/temp
type Location struct {
	Entity
	ParentLocationID int    `json:"parent_location_id"`
	Map              string `json:"map"`
	Type             string `json:"type"`
}

// LocationService handles communication with the Location endpoint.
type LocationService service
