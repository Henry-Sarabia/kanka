package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// EntityEvent contains information about a specific entityevent.
// For more information, visit: https://kanka.io/en-US/docs/1.0/entity-events
// EntityEvent represents a specific calendar event relating to the parent
// entity.
type EntityEvent struct {
	SimpleEntityEvent
	ID         int       `json:"id"`
	CalendarID int       `json:"calendar_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	Date       string    `json:"date"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
}

// SimpleEntityEvent contains only the simple information about a entityevent.
// SimpleEntityEvent is primarily used to create new entityevents for posting to Kanka.
type SimpleEntityEvent struct {
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Length         int    `json:"length"`
	EntityID       int    `json:"entity_id"`
	Colour         string `json:"colour,omitempty"`
	Comment        string `json:"comment,omitempty"`
	IsRecurring    bool   `json:"is_recurring,omitempty"`
	IsPrivate      bool   `json:"is_private,omitempty"`
	RecurringUntil int    `json:"recurring_until,omitempty"`
}

// EntityEvents wraps a list of entity events.
// EntityEvents exists to satisfy the API's JSON structure.
type EntityEvents struct {
	Data []*EntityEvent `json:"data"`
	Sync time.Time      `json:"sync"`
}

// EntityEventService handles communication with the EntityEvent endpoint.
type EntityEventService service

// Index returns the list of all EntityEvents for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return EntityEvents that have
// been changed since that time.
func (es *EntityEventService) Index(campID int, entID int, sync *time.Time) ([]*EntityEvent, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*EntityEvent `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityEvent Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the EntityEvent associated with evtID for the entity associated
// with entID from the Campaign associated with campID.
func (es *EntityEventService) Get(campID int, entID int, evtID int) (*EntityEvent, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(evtID); err != nil {
		return nil, fmt.Errorf("invalid EntityEvent ID: %w", err)
	}

	var wrap struct {
		Data *EntityEvent `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityEvent (ID: %d) from Campaign (ID: %d): %w", evtID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new EntityEvent for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleEntityEvent data.
// Create returns the newly created EntityEvent.
func (es *EntityEventService) Create(campID int, entID int, evt SimpleEntityEvent) (*EntityEvent, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	b, err := json.Marshal(evt)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityEvent: %w", err)
	}

	var wrap struct {
		Data *EntityEvent `json:"data"`
	}

	if err = es.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create EntityEvent for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing EntityEvent associated with evtID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleEntityEvent data.
// Update returns the newly updated EntityEvent.
func (es *EntityEventService) Update(campID int, entID int, evtID int, evt SimpleEntityEvent) (*EntityEvent, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(evtID); err != nil {
		return nil, fmt.Errorf("invalid EntityEvent ID: %w", err)
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityEvent: %w", err)
	}

	var wrap struct {
		Data *EntityEvent `json:"data"`
	}

	if err = es.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update EntityEvent for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing EntityEvent associated with evtID from the
// Campaign associated with campID.
func (es *EntityEventService) Delete(campID int, entID int, evtID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(evtID); err != nil {
		return fmt.Errorf("invalid EntityEvent ID: %w", err)
	}

	if err = es.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete EntityEvent (ID: %d) for Campaign (ID: %d): %w", evtID, campID, err)
	}

	return nil
}
