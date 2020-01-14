package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Location contains information about a specific location.
// For more information, visit: https://kanka.io/en-US/docs/1.0/locations
type Location struct {
	SimpleLocation
	ID             int       `json:"id"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
	IsMapPrivate   int       `json:"is_map_private"`
	HasCustomImage bool      `json:"has_custom_image"`
	EntityID       int       `json:"entity_id"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      int       `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      int       `json:"updated_by"`

	Attributes   Attributes   `json:"attributes"`
	EntityEvents EntityEvents `json:"entity_events"`
	EntityFiles  EntityFiles  `json:"entity_files"`
	EntityNotes  EntityNotes  `json:"entity_notes"`
	Relations    Relations    `json:"relations"`
	Inventory    Inventory    `json:"inventory"`
}

// SimpleLocation contains only the simple information about a Location.
// SimpleLocation is primarily used to create new Locations for posting to
// Kanka.
type SimpleLocation struct {
	Name             string `json:"name"`
	Entry            string `json:"entry,omitempty"`
	Type             string `json:"type,omitempty"`
	ParentLocationID int    `json:"parent_location_id,omitempty"`
	Tags             []int  `json:"tags,omitempty"`
	IsPrivate        bool   `json:"is_private,omitempty"`
	Image            string `json:"image,omitempty"`
	ImageURL         string `json:"image_url,omitempty"`
	Map              string `json:"map,omitempty"`
	MapURL           string `json:"map_url,omitempty"`
}

// MarshalJSON marshals the SimpleLocation into its JSON-encoded form if it has
// the required populated fields.
func (sl SimpleLocation) MarshalJSON() ([]byte, error) {
	if blank.Is(sl.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleLocation into JSON with a missing Name")
	}

	type alias SimpleLocation
	return json.Marshal(alias(sl))
}

// LocationService handles communication with the Location endpoint.
type LocationService service

// Index returns the list of all Locations in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Locations that have
// been changed since that time.
func (ls *LocationService) Index(campID int, sync *time.Time) ([]*Location, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ls.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Location `json:"data"`
	}

	err = ls.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Location Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Location associated with locID from the Campaign
// associated with campID.
func (ls *LocationService) Get(campID int, locID int) (*Location, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ls.end)

	end, err = end.ID(locID)
	if err != nil {
		return nil, fmt.Errorf("invalid Location ID: %w", err)
	}

	var wrap struct {
		Data *Location `json:"data"`
	}

	err = ls.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Location (ID: %d) from Campaign (ID: %d): %w", locID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Location in the Campaign associated with campID using
// the provided SimpleLocation data.
// Create returns the newly created Location.
func (ls *LocationService) Create(campID int, loc SimpleLocation) (*Location, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ls.end)

	b, err := json.Marshal(loc)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleLocation (Name: %s): %w", loc.Name, err)
	}

	var wrap struct {
		Data *Location `json:"data"`
	}

	err = ls.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Location (Name: %s) for Campaign (ID: %d): %w", loc.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Location associated with locID from the
// Campaign associated with campID using the provided SimpleLocation data.
// Update returns the newly updated Location.
func (ls *LocationService) Update(campID int, locID int, loc SimpleLocation) (*Location, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ls.end)

	end, err = end.ID(locID)
	if err != nil {
		return nil, fmt.Errorf("invalid Location ID: %w", err)
	}

	b, err := json.Marshal(loc)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleLocation (Name: %s): %w", loc.Name, err)
	}

	var wrap struct {
		Data *Location `json:"data"`
	}

	err = ls.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Location (Name: %s) for Campaign (ID: %d): '%w'", loc.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Location associated with locID from the
// Campaign associated with campID.
func (ls *LocationService) Delete(campID int, locID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ls.end)

	end, err = end.ID(locID)
	if err != nil {
		return fmt.Errorf("invalid Location ID: %w", err)
	}

	err = ls.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Location (ID: %d) for Campaign (ID: %d): %w", locID, campID, err)
	}

	return nil
}
