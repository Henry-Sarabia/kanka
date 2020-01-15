package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Event contains information about a specific event.
// For more information, visit: https://kanka.io/en-US/docs/1.0/events
type Event struct {
	SimpleEvent
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

// SimpleEvent contains only the simple information about an event.
// SimpleEvent is primarily used to create new events for posting to Kanka.
type SimpleEvent struct {
	Name       string `json:"name"`
	Entry      string `json:"entry,omitempty"`
	Type       string `json:"type,omitempty"`
	Date       string `json:"date,omitempty"`
	LocationID int    `json:"location_id,omitempty"`
	Tags       []int  `json:"tags,omitempty"`
	IsPrivate  bool   `json:"is_private,omitempty"`
	Image      string `json:"image,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleEvent into its JSON-encoded form if it
// has the required populated fields.
func (se SimpleEvent) MarshalJSON() ([]byte, error) {
	if blank.Is(se.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleEvent into JSON with a missing Name")
	}

	type alias SimpleEvent
	return json.Marshal(alias(se))
}

// EventService handles communication with the Event endpoint.
type EventService service

// Index returns the list of all Events in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Events that have
// been changed since that time.
func (es *EventService) Index(campID int, sync *time.Time) ([]*Event, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(es.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Event `json:"data"`
	}

	err = es.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Event Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Event associated with evtID from the Campaign
// associated with campID.
func (es *EventService) Get(campID int, evtID int) (*Event, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(es.end)

	end, err = end.ID(evtID)
	if err != nil {
		return nil, fmt.Errorf("invalid Event ID: %w", err)
	}

	var wrap struct {
		Data *Event `json:"data"`
	}

	err = es.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Event (ID: %d) from Campaign (ID: %d): %w", evtID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Event in the Campaign associated with campID using
// the provided SimpleEvent data.
// Create returns the newly created Event.
func (es *EventService) Create(campID int, evt SimpleEvent) (*Event, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(es.end)

	b, err := json.Marshal(evt)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEvent (Name: %s): %w", evt.Name, err)
	}

	var wrap struct {
		Data *Event `json:"data"`
	}

	err = es.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Event (Name: %s) for Campaign (ID: %d): %w", evt.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Event associated with evtID from the
// Campaign associated with campID using the provided SimpleEvent data.
// Update returns the newly updated Event.
func (es *EventService) Update(campID int, evtID int, evt SimpleEvent) (*Event, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(es.end)

	end, err = end.ID(evtID)
	if err != nil {
		return nil, fmt.Errorf("invalid Event ID: %w", err)
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEvent (Name: %s): %w", evt.Name, err)
	}

	var wrap struct {
		Data *Event `json:"data"`
	}

	err = es.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Event (Name: %s) for Campaign (ID: %d): '%w'", evt.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Event associated with evtID from the
// Campaign associated with campID.
func (es *EventService) Delete(campID int, evtID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(es.end)

	end, err = end.ID(evtID)
	if err != nil {
		return fmt.Errorf("invalid Event ID: %w", err)
	}

	err = es.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Event (ID: %d) for Campaign (ID: %d): %w", evtID, campID, err)
	}

	return nil
}
