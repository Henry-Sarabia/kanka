package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Organization contains informations about an organization.
// For more information. visit: https://kanka.io/en-US/docs/1.0/organisations
type Organization struct {
	SimpleOrganization
	ID             int       `json:"id"`
	Entry          string    `json:"entry"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
	HasCustomImage bool      `json:"has_custom_image"`
	EntityID       int       `json:"entity_id"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      int       `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      int       `json:"updated_by"`
	Members        int       `json:"members"`

	Attributes   Attributes   `json:"attributes"`
	EntityEvents EntityEvents `json:"entity_events"`
	EntityFiles  EntityFiles  `json:"entity_files"`
	EntityNotes  EntityNotes  `json:"entity_notes"`
	Relations    Relations    `json:"relations"`
	Inventory    Inventory    `json:"inventory"`
}

// SimpleOrganization contains only the simple information about an organization.
// SimpleOrganization is primarily used to create new organizations for posting
// to Kanka.
type SimpleOrganization struct {
	Name           string `json:"name"`
	Type           string `json:"type,omitempty"`
	OrganizationID int    `json:"organisation_id,omitempty"`
	LocationID     int    `json:"location_id,omitempty"`
	Tags           []int  `json:"tags,omitempty"`
	IsPrivate      bool   `json:"is_private,omitempty"`
	Image          string `json:"image,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleOrganization into its JSON-encoded form if it
// has the required populated fields.
func (so SimpleOrganization) MarshalJSON() ([]byte, error) {
	if blank.Is(so.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleOrganization into JSON with a missing Name")
	}

	type alias SimpleOrganization
	return json.Marshal(alias(so))
}

// OrganizationService handles communication with the Organization endpoint.
type OrganizationService service

// Index returns the list of all Organizations in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Organizations that have
// been changed since that time.
func (os *OrganizationService) Index(campID int, sync *time.Time) ([]*Organization, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(os.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Organization `json:"data"`
	}

	err = os.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Organization Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Organization associated with orgID from the Campaign
// associated with campID.
func (os *OrganizationService) Get(campID int, orgID int) (*Organization, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(os.end)

	end, err = end.ID(orgID)
	if err != nil {
		return nil, fmt.Errorf("invalid Organization ID: %w", err)
	}

	var wrap struct {
		Data *Organization `json:"data"`
	}

	err = os.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Organization (ID: %d) from Campaign (ID: %d): %w", orgID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Organization in the Campaign associated with campID using
// the provided SimpleOrganization data.
// Create returns the newly created Organization.
func (os *OrganizationService) Create(campID int, org SimpleOrganization) (*Organization, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(os.end)

	b, err := json.Marshal(org)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleOrganization (Name: %s): %w", org.Name, err)
	}

	var wrap struct {
		Data *Organization `json:"data"`
	}

	err = os.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Organization (Name: %s) for Campaign (ID: %d): %w", org.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Organization associated with orgID from the
// Campaign associated with campID using the provided SimpleOrganization data.
// Update returns the newly updated Organization.
func (os *OrganizationService) Update(campID int, orgID int, org SimpleOrganization) (*Organization, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(os.end)

	end, err = end.ID(orgID)
	if err != nil {
		return nil, fmt.Errorf("invalid Organization ID: %w", err)
	}

	b, err := json.Marshal(org)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleOrganization (Name: %s): %w", org.Name, err)
	}

	var wrap struct {
		Data *Organization `json:"data"`
	}

	err = os.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Organization (Name: %s) for Campaign (ID: %d): '%w'", org.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Organization associated with orgID from the
// Campaign associated with campID.
func (os *OrganizationService) Delete(campID int, orgID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(os.end)

	end, err = end.ID(orgID)
	if err != nil {
		return fmt.Errorf("invalid Organization ID: %w", err)
	}

	err = os.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Organization (ID: %d) for Campaign (ID: %d): %w", orgID, campID, err)
	}

	return nil
}
