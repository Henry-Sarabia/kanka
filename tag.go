package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Tag contains information about a specific tag.
// For more information, visit: https://kanka.io/en-US/docs/1.0/tags
type Tag struct {
	SimpleTag
	ID             int       `json:"id"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
	HasCustomImage bool      `json:"has_custom_image"`
	EntityID       int       `json:"entity_id"`
	Entities       []int     `json:"entities"`
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

// SimpleTag contains only the simple information about a tag.
// SimpleTag is primarily used to create new tags for posting to Kanka.
type SimpleTag struct {
	Name      string `json:"name"`
	Entry     string `json:"entry,omitempty"`
	Type      string `json:"type,omitempty"`
	TagID     int    `json:"tag_id,omitempty"`
	Color     string `json:"colour,omitempty"`
	Tags      []int  `json:"tags,omitempty"`
	IsPrivate bool   `json:"is_private,omitempty"`
	Image     string `json:"image,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleTag into its JSON-encoded form if it
// has the required populated fields.
func (st SimpleTag) MarshalJSON() ([]byte, error) {
	if blank.Is(st.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleTag into JSON with a missing Name")
	}

	type alias SimpleTag
	return json.Marshal(alias(st))
}

// TagService handles communication with the Tag endpoint.
type TagService service

// Index returns the list of all Tags in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Tags that have
// been changed since that time.
func (ts *TagService) Index(campID int, sync *time.Time) ([]*Tag, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ts.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Tag `json:"data"`
	}

	err = ts.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Tag Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Tag associated with tagID from the Campaign
// associated with campID.
func (ts *TagService) Get(campID int, tagID int) (*Tag, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ts.end)

	end, err = end.ID(tagID)
	if err != nil {
		return nil, fmt.Errorf("invalid Tag ID: %w", err)
	}

	var wrap struct {
		Data *Tag `json:"data"`
	}

	err = ts.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Tag (ID: %d) from Campaign (ID: %d): %w", tagID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Tag in the Campaign associated with campID using
// the provided SimpleTag data.
// Create returns the newly created Tag.
func (ts *TagService) Create(campID int, tag SimpleTag) (*Tag, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ts.end)

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleTag (Name: %s): %w", tag.Name, err)
	}

	var wrap struct {
		Data *Tag `json:"data"`
	}

	err = ts.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Tag (Name: %s) for Campaign (ID: %d): %w", tag.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Tag associated with tagID from the
// Campaign associated with campID using the provided SimpleTag data.
// Update returns the newly updated Tag.
func (ts *TagService) Update(campID int, tagID int, tag SimpleTag) (*Tag, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ts.end)

	end, err = end.ID(tagID)
	if err != nil {
		return nil, fmt.Errorf("invalid Tag ID: %w", err)
	}

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleTag (Name: %s): %w", tag.Name, err)
	}

	var wrap struct {
		Data *Tag `json:"data"`
	}

	err = ts.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Tag (Name: %s) for Campaign (ID: %d): '%w'", tag.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Tag associated with tagID from the
// Campaign associated with campID.
func (ts *TagService) Delete(campID int, tagID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ts.end)

	end, err = end.ID(tagID)
	if err != nil {
		return fmt.Errorf("invalid Tag ID: %w", err)
	}

	err = ts.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Tag (ID: %d) for Campaign (ID: %d): %w", tagID, campID, err)
	}

	return nil
}
