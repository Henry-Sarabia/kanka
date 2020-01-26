package kanka

import (
	"fmt"
	"time"
)

// Result contains the response to a search query.
// For more information, visit: https://kanka.io/en-US/docs/1.0/search
type Result struct {
	ID                  int       `json:"id"`
	EntityID            int       `json:"entity_id"`
	Name                string    `json:"name"`
	Image               string    `json:"image"`
	ImageThumb          string    `json:"image_thumb"`
	HasCustomImage      bool      `json:"has_custom_image"`
	Type                string    `json:"type"`
	Tooltip             string    `json:"tooltip"`
	URL                 string    `json:"url"`
	IsAttributesPrivate int       `json:"is_attributes_private"`
	IsPrivate           bool      `json:"is_private"`
	CreatedAt           time.Time `json:"created_at"`
	CreatedBy           int       `json:"created_by"`
	UpdatedAt           time.Time `json:"updated_at"`
	UpdatedBy           int       `json:"updated_by"`
}

// Results wraps a list of results.
// Results exists to satisfy the API's JSON structure.
type Results struct {
	Data []*Result `json:"data"`
	Sync time.Time `json:"sync"`
}

// Search searches the Campaign associated with campID for the provided query.
func (c *Client) Search(campID int, qry string, sync *time.Time) ([]*Result, error) {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointSearch)
	end = end.append("/" + qry)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap Results

	if err = c.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get Search results from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}
