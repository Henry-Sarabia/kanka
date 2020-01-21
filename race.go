package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Race contains information about a specific race.
// For more information, visit: https://kanka.io/en-US/docs/1.0/races
type Race struct {
	SimpleRace
	ID             int       `json:"id"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
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

// SimpleRace contains only the simple information about a race.
// SimpleRace is primarily used to create new races for posting to Kanka.
type SimpleRace struct {
	Name      string `json:"name"`
	Entry     string `json:"entry,omitempty"`
	Type      string `json:"type,omitempty"`
	RaceID    int    `json:"race_id,omitempty"`
	Tags      []int  `json:"tags,omitempty"`
	IsPrivate bool   `json:"is_private,omitempty"`
	Image     string `json:"image,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleRace into its JSON-encoded form if it
// has the required populated fields.
func (sr SimpleRace) MarshalJSON() ([]byte, error) {
	if blank.Is(sr.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleRace into JSON with a missing Name")
	}

	type alias SimpleRace
	return json.Marshal(alias(sr))
}

// RaceService handles communication with the Race endpoint.
type RaceService service

// Index returns the list of all Races in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Races that have
// been changed since that time.
func (rs *RaceService) Index(campID int, sync *time.Time) ([]*Race, error) {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(rs.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*Race `json:"data"`
	}

	err = rs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Race Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Race associated with raceID from the Campaign
// associated with campID.
func (rs *RaceService) Get(campID int, raceID int) (*Race, error) {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(rs.end)

	end, err = end.id(raceID)
	if err != nil {
		return nil, fmt.Errorf("invalid Race ID: %w", err)
	}

	var wrap struct {
		Data *Race `json:"data"`
	}

	err = rs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Race (ID: %d) from Campaign (ID: %d): %w", raceID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Race in the Campaign associated with campID using
// the provided SimpleRace data.
// Create returns the newly created Race.
func (rs *RaceService) Create(campID int, race SimpleRace) (*Race, error) {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(rs.end)

	b, err := json.Marshal(race)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleRace (Name: %s): %w", race.Name, err)
	}

	var wrap struct {
		Data *Race `json:"data"`
	}

	err = rs.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Race (Name: %s) for Campaign (ID: %d): %w", race.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Race associated with raceID from the
// Campaign associated with campID using the provided SimpleRace data.
// Update returns the newly updated Race.
func (rs *RaceService) Update(campID int, raceID int, race SimpleRace) (*Race, error) {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(rs.end)

	end, err = end.id(raceID)
	if err != nil {
		return nil, fmt.Errorf("invalid Race ID: %w", err)
	}

	b, err := json.Marshal(race)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleRace (Name: %s): %w", race.Name, err)
	}

	var wrap struct {
		Data *Race `json:"data"`
	}

	err = rs.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Race (Name: %s) for Campaign (ID: %d): '%w'", race.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Race associated with raceID from the
// Campaign associated with campID.
func (rs *RaceService) Delete(campID int, raceID int) error {
	end, err := EndpointCampaign.id(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(rs.end)

	end, err = end.id(raceID)
	if err != nil {
		return fmt.Errorf("invalid Race ID: %w", err)
	}

	err = rs.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Race (ID: %d) for Campaign (ID: %d): %w", raceID, campID, err)
	}

	return nil
}
