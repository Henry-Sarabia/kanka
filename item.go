package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Item contains information about a specific item.
// For more information, visit: https://kanka.io/en-US/docs/1.0/items
type Item struct {
	SimpleItem
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

// SimpleItem contains only the simple information about an item.
// SimpleItem is primarily used to create new items for posting to Kanka.
type SimpleItem struct {
	Name        string `json:"name"`
	Entry       string `json:"entry,omitempty"`
	Type        string `json:"type,omitempty"`
	Price       string `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	LocationID  int    `json:"location_id,omitempty"`
	CharacterID int    `json:"character_id,omitempty"`
	Tags        []int  `json:"tags,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
	Image       string `json:"image,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleItem into its JSON-encoded form if it
// has the required populated fields.
func (si SimpleItem) MarshalJSON() ([]byte, error) {
	if blank.Is(si.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleItem into JSON with a missing Name")
	}

	type alias SimpleItem
	return json.Marshal(alias(si))
}

// ItemService handles communication with the Item endpoint.
type ItemService service

// Index returns the list of all Items in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Items that have
// been changed since that time.
func (is *ItemService) Index(campID int, sync *time.Time) ([]*Item, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(is.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Item `json:"data"`
	}

	err = is.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Item Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Item associated with itemID from the Campaign
// associated with campID.
func (is *ItemService) Get(campID int, itemID int) (*Item, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(is.end)

	end, err = end.ID(itemID)
	if err != nil {
		return nil, fmt.Errorf("invalid Item ID: %w", err)
	}

	var wrap struct {
		Data *Item `json:"data"`
	}

	err = is.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Item (ID: %d) from Campaign (ID: %d): %w", itemID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Item in the Campaign associated with campID using
// the provided SimpleItem data.
// Create returns the newly created Item.
func (is *ItemService) Create(campID int, item SimpleItem) (*Item, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(is.end)

	b, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleItem (Name: %s): %w", item.Name, err)
	}

	var wrap struct {
		Data *Item `json:"data"`
	}

	err = is.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Item (Name: %s) for Campaign (ID: %d): %w", item.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Item associated with itemID from the
// Campaign associated with campID using the provided SimpleItem data.
// Update returns the newly updated Item.
func (is *ItemService) Update(campID int, itemID int, item SimpleItem) (*Item, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(is.end)

	end, err = end.ID(itemID)
	if err != nil {
		return nil, fmt.Errorf("invalid Item ID: %w", err)
	}

	b, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleItem (Name: %s): %w", item.Name, err)
	}

	var wrap struct {
		Data *Item `json:"data"`
	}

	err = is.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Item (Name: %s) for Campaign (ID: %d): '%w'", item.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Item associated with itemID from the
// Campaign associated with campID.
func (is *ItemService) Delete(campID int, itemID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(is.end)

	end, err = end.ID(itemID)
	if err != nil {
		return fmt.Errorf("invalid Item ID: %w", err)
	}

	err = is.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Item (ID: %d) for Campaign (ID: %d): %w", itemID, campID, err)
	}

	return nil
}
