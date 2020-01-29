package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// MapPoint contains information about a specific map point.
// For more information, visit: https://kanka.io/en-US/docs/1.0/locations#location-map-points
type MapPoint struct {
	SimpleMapPoint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SimpleMapPoint contains only the simple information about a map point.
// SimpleMapPoint is primarily used to create new mappoints for posting to Kanka.
type SimpleMapPoint struct {
	LocationID     int    `json:"location_id"`
	TargetEntityID int    `json:"target_entity_id,omitempty"`
	Name           string `json:"name,omitempty"`
	AxisX          int    `json:"axis_x"`
	AxisY          int    `json:"axis_y"`
	Color          string `json:"colour"`
	Icon           string `json:"icon"`
	Shape          string `json:"shape"`
	Size           string `json:"size"`
}

// MarshalJSON marshals the SimpleMapPoint into its JSON-encoded form if it
// has the required populated fields.
func (sm SimpleMapPoint) MarshalJSON() ([]byte, error) {
	if blank.Is(sm.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint into JSON with a missing Name")
	}
	if blank.Is(sm.Color) {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint into JSON with a missing Color")
	}
	if blank.Is(sm.Icon) {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint into JSON with a missing Icon")
	}
	if blank.Is(sm.Shape) {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint into JSON with a missing Shape")
	}
	if blank.Is(sm.Size) {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint into JSON with a missing Size")
	}

	type alias SimpleMapPoint
	return json.Marshal(alias(sm))
}

// MapPointService handles communication with the MapPoint endpoint.
type MapPointService service

// Index returns the list of all MapPoints for the location associated with
// locID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return MapPoints that have
// been changed since that time.
func (ms *MapPointService) Index(campID int, locID int, sync *time.Time) ([]*MapPoint, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointLocation)

	if end, err = end.id(locID); err != nil {
		return nil, fmt.Errorf("invalid Location ID: %w", err)
	}
	end = end.concat(ms.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*MapPoint `json:"data"`
	}

	if err = ms.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get MapPoint Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new MapPoint for the location associated with locID in the
// Campaign associated with campID using the provided SimpleMapPoint data.
// Create returns the newly created MapPoint.
func (ms *MapPointService) Create(campID int, locID int, mp SimpleMapPoint) (*MapPoint, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointLocation)

	if end, err = end.id(locID); err != nil {
		return nil, fmt.Errorf("invalid Location ID: %w", err)
	}
	end = end.concat(ms.end)

	b, err := json.Marshal(mp)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint (Name: %s): %w", mp.Name, err)
	}

	var wrap struct {
		Data *MapPoint `json:"data"`
	}

	if err = ms.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create MapPoint (Name: %s) for Campaign (ID: %d): %w", mp.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing MapPoint associated with mpID for the location
// associated with locID from the Campaign associated with campID using the
// provided SimpleMapPoint data.
// Update returns the newly updated MapPoint.
func (ms *MapPointService) Update(campID int, locID int, mpID int, mp SimpleMapPoint) (*MapPoint, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointLocation)

	if end, err = end.id(locID); err != nil {
		return nil, fmt.Errorf("invalid Location ID: %w", err)
	}
	end = end.concat(ms.end)

	if end, err = end.id(mpID); err != nil {
		return nil, fmt.Errorf("invalid MapPoint ID: %w", err)
	}

	b, err := json.Marshal(mp)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleMapPoint (Name: %s): %w", mp.Name, err)
	}

	var wrap struct {
		Data *MapPoint `json:"data"`
	}

	if err = ms.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update MapPoint (Name: %s) for Campaign (ID: %d): '%w'", mp.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing MapPoint associated with mpID from the
// Campaign associated with campID.
func (ms *MapPointService) Delete(campID int, locID int, mpID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointLocation)

	if end, err = end.id(locID); err != nil {
		return fmt.Errorf("invalid Location ID: %w", err)
	}
	end = end.concat(ms.end)

	if end, err = end.id(mpID); err != nil {
		return fmt.Errorf("invalid MapPoint ID: %w", err)
	}

	if err = ms.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete MapPoint (ID: %d) for Campaign (ID: %d): %w", mpID, campID, err)
	}

	return nil
}
