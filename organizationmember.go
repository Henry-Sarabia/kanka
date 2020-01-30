package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// OrganizationMember contains information about a specific organization member.
// For more information, visit: https://kanka.io/en-US/docs/1.0/organisations#organisation-members
type OrganizationMember struct {
	SimpleOrganizationMember
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	ID        int       `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleOrganizationMember contains only the simple information about an
// organization member.
// SimpleOrganizationMember is primarily used to create new organization
// members for posting to Kanka.
type SimpleOrganizationMember struct {
	CharacterID    int    `json:"character_id"`
	OrganizationID int    `json:"organisation_id"`
	Role           string `json:"role,omitempty"`
	IsPrivate      bool   `json:"is_private,omitempty"`
}

// OrganizationMemberService handles communication with the OrganizationMember endpoint.
type OrganizationMemberService service

// Index returns the list of all OrganizationMembers for the organization
// associated with orgID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return OrganizationMembers
// that have been changed since that time.
func (os *OrganizationMemberService) Index(campID int, orgID int, sync *time.Time) ([]*OrganizationMember, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointOrganization)

	if end, err = end.id(orgID); err != nil {
		return nil, fmt.Errorf("invalid Organization ID: %w", err)
	}
	end = end.concat(os.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*OrganizationMember `json:"data"`
	}

	if err = os.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get OrganizationMember Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the OrganizationMember associated with memID for the organization
// associated with orgID from the Campaign associated with campID.
func (os *OrganizationMemberService) Get(campID int, orgID int, memID int) (*OrganizationMember, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointOrganization)

	if end, err = end.id(orgID); err != nil {
		return nil, fmt.Errorf("invalid Organization ID: %w", err)
	}
	end = end.concat(os.end)

	if end, err = end.id(memID); err != nil {
		return nil, fmt.Errorf("invalid OrganizationMember ID: %w", err)
	}

	var wrap struct {
		Data *OrganizationMember `json:"data"`
	}

	if err = os.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get OrganizationMember (ID: %d) from Campaign (ID: %d): %w", memID, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing OrganizationMember associated with memID for the
// organization associated with orgID from the Campaign associated with campID
// using the provided SimpleOrganizationMember data.
// Update returns the newly updated OrganizationMember.
func (os *OrganizationMemberService) Update(campID int, orgID int, memID int, mem SimpleOrganizationMember) (*OrganizationMember, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointOrganization)

	if end, err = end.id(orgID); err != nil {
		return nil, fmt.Errorf("invalid Organization ID: %w", err)
	}
	end = end.concat(os.end)

	if end, err = end.id(memID); err != nil {
		return nil, fmt.Errorf("invalid OrganizationMember ID: %w", err)
	}

	b, err := json.Marshal(mem)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleOrganizationMember: %w", err)
	}

	var wrap struct {
		Data *OrganizationMember `json:"data"`
	}

	if err = os.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update OrganizationMember for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing OrganizationMember associated with memID from the
// Campaign associated with campID.
func (os *OrganizationMemberService) Delete(campID int, orgID int, memID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointOrganization)

	if end, err = end.id(orgID); err != nil {
		return fmt.Errorf("invalid Organization ID: %w", err)
	}
	end = end.concat(os.end)

	if end, err = end.id(memID); err != nil {
		return fmt.Errorf("invalid OrganizationMember ID: %w", err)
	}

	if err = os.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete OrganizationMember (ID: %d) for Campaign (ID: %d): %w", memID, campID, err)
	}

	return nil
}
