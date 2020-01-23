package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// EntityTag contains information about a specific entitytag.
// For more information, visit: https://kanka.io/en-US/docs/1.0/entity-tags
// EntityTag represents a specific tag relating to the parent entity.
type EntityTag struct {
	SimpleEntityTag
	ID int `json:"id"`
}

// SimpleEntityTag contains only the simple information about an entity tag.
// SimpleEntityTag is primarily used to create new entitytags for posting to Kanka.
type SimpleEntityTag struct {
	EntityID int `json:"entity_id"`
	TagID    int `json:"tag_id"`
}

// EntityTagService handles communication with the EntityTag endpoint.
type EntityTagService service

// Index returns the list of all EntityTags for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return EntityTags that have
// been changed since that time.
func (es *EntityTagService) Index(campID int, entID int, sync *time.Time) ([]*EntityTag, error) {
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
		Data []*EntityTag `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityTag Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the EntityTag associated with tagID for the entity associated
// with entID from the Campaign associated with campID.
func (es *EntityTagService) Get(campID int, entID int, tagID int) (*EntityTag, error) {
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

	if end, err = end.id(tagID); err != nil {
		return nil, fmt.Errorf("invalid EntityTag ID: %w", err)
	}

	var wrap struct {
		Data *EntityTag `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityTag (ID: %d) from Campaign (ID: %d): %w", tagID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new EntityTag for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleEntityTag data.
// Create returns the newly created EntityTag.
func (es *EntityTagService) Create(campID int, entID int, tag SimpleEntityTag) (*EntityTag, error) {
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

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityTag: %w", err)
	}

	var wrap struct {
		Data *EntityTag `json:"data"`
	}

	if err = es.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create EntityTag for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing EntityTag associated with tagID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleEntityTag data.
// Update returns the newly updated EntityTag.
func (es *EntityTagService) Update(campID int, entID int, tagID int, tag SimpleEntityTag) (*EntityTag, error) {
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

	if end, err = end.id(tagID); err != nil {
		return nil, fmt.Errorf("invalid EntityTag ID: %w", err)
	}

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityTag: %w", err)
	}

	var wrap struct {
		Data *EntityTag `json:"data"`
	}

	if err = es.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update EntityTag for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing EntityTag associated with tagID from the
// Campaign associated with campID.
func (es *EntityTagService) Delete(campID int, entID int, tagID int) error {
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

	if end, err = end.id(tagID); err != nil {
		return fmt.Errorf("invalid EntityTag ID: %w", err)
	}

	if err = es.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete EntityTag (ID: %d) for Campaign (ID: %d): %w", tagID, campID, err)
	}

	return nil
}
