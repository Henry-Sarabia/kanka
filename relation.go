package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Relation contains information about a specific relation.
// For more information, visit: https://kanka.io/en-US/docs/1.0/relations
// Relation represents a relationship between two entities.
type Relation struct {
	SimpleRelation
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SimpleRelation contains only the simple information about a relation.
// SimpleRelation is primarily used to create new relations for posting to Kanka.
type SimpleRelation struct {
	Relation  string `json:"relation"`
	OwnerID   int    `json:"owner_id"`
	TargetID  int    `json:"target_id"`
	Attitude  int    `json:"attitude"`
	TwoWay    bool   `json:"two_way,omitempty"`
	IsPrivate bool   `json:"is_private,omitempty"`
}

// For more information, visit: https://kanka.io/en-US/docs/1.0/relations#create-relation
const (
	relationLengthMax int = 255
	attitudeMin       int = -100
	attitudeMax       int = 100
)

// MarshalJSON marshals the SimpleRelation into its JSON-encoded form if it
// has the required populated fields.
func (sr SimpleRelation) MarshalJSON() ([]byte, error) {
	if blank.Is(sr.Relation) {
		return nil, fmt.Errorf("cannot marshal SimpleRelation into JSON with a missing Relation")
	}

	if len(sr.Relation) > relationLengthMax {
		return nil, fmt.Errorf("length of Relation string must not exceed %d characters", relationLengthMax)
	}

	if sr.Attitude < attitudeMin || sr.Attitude > attitudeMax {
		return nil, fmt.Errorf("value of Attitude must be between %d and %d", attitudeMin, attitudeMax)
	}

	type alias SimpleRelation
	return json.Marshal(alias(sr))
}

// Relations wraps a list of relationships.
// Relations exists to satisfy the API's JSON structure.
type Relations struct {
	Data []Relation `json:"data"`
	Sync time.Time  `json:"sync"`
}

// RelationService handles communication with the Relation endpoint.
type RelationService service

// Index returns the list of all Relations for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Relations that have
// been changed since that time.
func (rs *RelationService) Index(campID int, entID int, sync *time.Time) ([]*Relation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(rs.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*Relation `json:"data"`
	}

	if err = rs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get Relation Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Relation associated with relID for the entity associated
// with entID from the Campaign associated with campID.
func (rs *RelationService) Get(campID int, entID int, relID int) (*Relation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(rs.end)

	if end, err = end.id(relID); err != nil {
		return nil, fmt.Errorf("invalid Relation ID: %w", err)
	}

	var wrap struct {
		Data *Relation `json:"data"`
	}

	if err = rs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get Relation (ID: %d) from Campaign (ID: %d): %w", relID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Relation for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleRelation data.
// Create returns the newly created Relation.
func (rs *RelationService) Create(campID int, entID int, rel SimpleRelation) (*Relation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(rs.end)

	b, err := json.Marshal(rel)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleRelation (Relation: %s): %w", rel.Relation, err)
	}

	var wrap struct {
		Data *Relation `json:"data"`
	}

	if err = rs.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create Relation (Relation: %s) for Campaign (ID: %d): %w", rel.Relation, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Relation associated with relID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleRelation data.
// Update returns the newly updated Relation.
func (rs *RelationService) Update(campID int, entID int, relID int, rel SimpleRelation) (*Relation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(rs.end)

	if end, err = end.id(relID); err != nil {
		return nil, fmt.Errorf("invalid Relation ID: %w", err)
	}

	b, err := json.Marshal(rel)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleRelation (Relation: %s): %w", rel.Relation, err)
	}

	var wrap struct {
		Data *Relation `json:"data"`
	}

	if err = rs.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update Relation (Relation: %s) for Campaign (ID: %d): '%w'", rel.Relation, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Relation associated with relID from the
// Campaign associated with campID.
func (rs *RelationService) Delete(campID int, entID int, relID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(rs.end)

	if end, err = end.id(relID); err != nil {
		return fmt.Errorf("invalid Relation ID: %w", err)
	}

	if err = rs.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete Relation (ID: %d) for Campaign (ID: %d): %w", relID, campID, err)
	}

	return nil
}
