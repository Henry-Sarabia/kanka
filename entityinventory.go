package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// EntityInventory contains information about a specific entity inventory.
// For more information, visit: https://kanka.io/en-US/docs/1.0/entity-inventory
type EntityInventory struct {
	SimpleEntityInventory
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleEntityInventory contains only the simple information about an entity inventory.
// SimpleEntityInventory is primarily used to create new entity inventories for posting to Kanka.
type SimpleEntityInventory struct {
	EntityID   int    `json:"entity_id"`
	ItemID     int    `json:"item_id"`
	Amount     int    `json:"amount"`
	Position   string `json:"position,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	IsPrivate  bool   `json:"is_private,omitempty"`
}

// EntityInventoryService handles communication with the EntityInventory endpoint.
type EntityInventoryService service

// Index returns the list of all EntityInventories for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return EntityInventories that have
// been changed since that time.
func (es *EntityInventoryService) Index(campID int, entID int, sync *time.Time) ([]*EntityInventory, error) {
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
		Data []*EntityInventory `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityInventory Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new EntityInventory for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleEntityInventory data.
// Create returns the newly created EntityInventory.
func (es *EntityInventoryService) Create(campID int, entID int, inv SimpleEntityInventory) (*EntityInventory, error) {
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

	b, err := json.Marshal(inv)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityInventory: %w", err)
	}

	var wrap struct {
		Data *EntityInventory `json:"data"`
	}

	if err = es.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create EntityInventory for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing EntityInventory associated with invID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleEntityInventory data.
// Update returns the newly updated EntityInventory.
func (es *EntityInventoryService) Update(campID int, entID int, invID int, inv SimpleEntityInventory) (*EntityInventory, error) {
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

	if end, err = end.id(invID); err != nil {
		return nil, fmt.Errorf("invalid EntityInventory ID: %w", err)
	}

	b, err := json.Marshal(inv)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityInventory: %w", err)
	}

	var wrap struct {
		Data *EntityInventory `json:"data"`
	}

	if err = es.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update EntityInventory for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing EntityInventory associated with invID from the
// Campaign associated with campID.
func (es *EntityInventoryService) Delete(campID int, entID int, invID int) error {
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

	if end, err = end.id(invID); err != nil {
		return fmt.Errorf("invalid EntityInventory ID: %w", err)
	}

	if err = es.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete EntityInventory (ID: %d) for Campaign (ID: %d): %w", invID, campID, err)
	}

	return nil
}
