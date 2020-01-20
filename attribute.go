package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Attribute represents a distinct detail relating to the parent entity.
// For more information, visit: https://kanka.io/en-US/docs/1.0/attributes
type Attribute struct {
	SimpleAttribute
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleAttribute contains only the simple information about a attribute.
// SimpleAttribute is primarily used to create new attributes for posting to Kanka.
type SimpleAttribute struct {
	Name         string `json:"name"`
	Value        string `json:"value"`
	DefaultOrder int    `json:"default_order"`
	Type         string `json:"type"`
	EntityID     int    `json:"entity_id"`
	IsPrivate    bool   `json:"is_private"`
	APIKey       string `json:"api_key"`
}

// MarshalJSON marshals the SimpleAttribute into its JSON-encoded form if it
// has the required populated fields.
func (sa SimpleAttribute) MarshalJSON() ([]byte, error) {
	if blank.Is(sa.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleAttribute into JSON with a missing Name")
	}

	type alias SimpleAttribute
	return json.Marshal(alias(sa))
}

// Attributes wraps a list of attributes.
// Attributes exists to satisfy the API's JSON structure.
type Attributes struct {
	Data []*Attribute `json:"data"`
	Sync time.Time    `json:"sync"`
}

// AttributeService handles communication with the Attribute endpoint.
type AttributeService service

// Index returns the list of all Attributes for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Attributes that have
// been changed since that time.
func (as *AttributeService) Index(campID int, entID int, sync *time.Time) ([]*Attribute, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.ID(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(endpointEntity)

	if end, err = end.ID(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.Concat(as.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Attribute `json:"data"`
	}

	if err = as.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get Attribute Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Attribute associated with atrID for the entity associated
// with entID from the Campaign associated with campID.
func (as *AttributeService) Get(campID int, entID int, atrID int) (*Attribute, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.ID(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(endpointEntity)

	if end, err = end.ID(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.Concat(as.end)

	if end, err = end.ID(atrID); err != nil {
		return nil, fmt.Errorf("invalid Attribute ID: %w", err)
	}

	var wrap struct {
		Data *Attribute `json:"data"`
	}

	if err = as.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get Attribute (ID: %d) from Campaign (ID: %d): %w", atrID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Attribute for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleAttribute data.
// Create returns the newly created Attribute.
func (as *AttributeService) Create(campID int, entID int, atr SimpleAttribute) (*Attribute, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.ID(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(endpointEntity)

	if end, err = end.ID(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.Concat(as.end)

	b, err := json.Marshal(atr)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleAttribute (Name: %s): %w", atr.Name, err)
	}

	var wrap struct {
		Data *Attribute `json:"data"`
	}

	if err = as.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create Attribute (Name: %s) for Campaign (ID: %d): %w", atr.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Attribute associated with atrID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleAttribute data.
// Update returns the newly updated Attribute.
func (as *AttributeService) Update(campID int, entID int, atrID int, atr SimpleAttribute) (*Attribute, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.ID(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(endpointEntity)

	if end, err = end.ID(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.Concat(as.end)

	if end, err = end.ID(atrID); err != nil {
		return nil, fmt.Errorf("invalid Attribute ID: %w", err)
	}

	b, err := json.Marshal(atr)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleAttribute (Name: %s): %w", atr.Name, err)
	}

	var wrap struct {
		Data *Attribute `json:"data"`
	}

	if err = as.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update Attribute (Name: %s) for Campaign (ID: %d): '%w'", atr.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Attribute associated with atrID from the
// Campaign associated with campID.
func (as *AttributeService) Delete(campID int, entID int, atrID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.ID(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(endpointEntity)

	if end, err = end.ID(entID); err != nil {
		return fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.Concat(as.end)

	if end, err = end.ID(atrID); err != nil {
		return fmt.Errorf("invalid Attribute ID: %w", err)
	}

	if err = as.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete Attribute (ID: %d) for Campaign (ID: %d): %w", atrID, campID, err)
	}

	return nil
}
