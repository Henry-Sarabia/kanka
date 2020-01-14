package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Family contains information about a family.
// For more information, visit: https://kanka.io/en-US/dofs/1.0/families
type Family struct {
	SimpleFamily
	ID             int       `json:"id"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
	HasCustomImage bool      `json:"has_custom_image"`
	EntityID       int       `json:"entity_id"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      int       `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      int       `json:"updated_by"`
	Members        []int     `json:"members"`

	Attributes   Attributes   `json:"attributes"`
	EntityEvents EntityEvents `json:"entity_events"`
	EntityFiles  EntityFiles  `json:"entity_files"`
	EntityNotes  EntityNotes  `json:"entity_notes"`
	Relations    Relations    `json:"relations"`
	Inventory    Inventory    `json:"inventory"`
}

// SimpleFamily contains only the simple information about a family.
// SimpleFamily is primarily used to create new families for posting to
// Kanka.
type SimpleFamily struct {
	Name       string `json:"name"`
	Entry      string `json:"entry,omitempty"`
	Type       string `json:"type,omitempty"`
	LocationID int    `json:"location_id,omitempty"`
	FamilyID   int    `json:"family_id,omitempty"`
	Tags       []int  `json:"tags,omitempty"`
	IsPrivate  bool   `json:"is_private,omitempty"`
	Image      string `json:"image,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleFamily into its JSON-encoded form if it
// has the required populated fields.
func (sf SimpleFamily) MarshalJSON() ([]byte, error) {
	if blank.Is(sf.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleFamily into JSON with a missing Name")
	}

	type alias SimpleFamily
	return json.Marshal(alias(sf))
}

// FamilyService handles communication with the Family endpoint.
type FamilyService service

// Index returns the list of all Families in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Families that have
// been changed since that time.
func (fs *FamilyService) Index(campID int, sync *time.Time) ([]*Family, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(fs.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Family `json:"data"`
	}

	err = fs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Family Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Family associated with famID from the Campaign
// associated with campID.
func (fs *FamilyService) Get(campID int, famID int) (*Family, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(fs.end)

	end, err = end.ID(famID)
	if err != nil {
		return nil, fmt.Errorf("invalid Family ID: %w", err)
	}

	var wrap struct {
		Data *Family `json:"data"`
	}

	err = fs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Family (ID: %d) from Campaign (ID: %d): %w", famID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Family in the Campaign associated with campID using
// the provided SimpleFamily data.
// Create returns the newly created Family.
func (fs *FamilyService) Create(campID int, fam SimpleFamily) (*Family, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(fs.end)

	b, err := json.Marshal(fam)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleFamily (Name: %s): %w", fam.Name, err)
	}

	var wrap struct {
		Data *Family `json:"data"`
	}

	err = fs.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Family (Name: %s) for Campaign (ID: %d): %w", fam.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Family associated with famID from the
// Campaign associated with campID using the provided SimpleFamily data.
// Update returns the newly updated Family.
func (fs *FamilyService) Update(campID int, famID int, fam SimpleFamily) (*Family, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(fs.end)

	end, err = end.ID(famID)
	if err != nil {
		return nil, fmt.Errorf("invalid Family ID: %w", err)
	}

	b, err := json.Marshal(fam)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleFamily (Name: %s): %w", fam.Name, err)
	}

	var wrap struct {
		Data *Family `json:"data"`
	}

	err = fs.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Family (Name: %s) for Campaign (ID: %d): '%w'", fam.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Family associated with famID from the
// Campaign associated with campID.
func (fs *FamilyService) Delete(campID int, famID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(fs.end)

	end, err = end.ID(famID)
	if err != nil {
		return fmt.Errorf("invalid Family ID: %w", err)
	}

	err = fs.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Family (ID: %d) for Campaign (ID: %d): %w", famID, campID, err)
	}

	return nil
}
