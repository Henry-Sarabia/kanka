package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// QuestOrganization contains information about a specific quest organization.
// For more information, visit: https://kanka.io/en-US/docs/1.0/quests#quest-organisations
type QuestOrganization struct {
	SimpleQuestOrganization
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleQuestOrganization contains only the simple information about a quest organization.
// SimpleQuestOrganization is primarily used to create new quest organizations for posting to Kanka.
type SimpleQuestOrganization struct {
	QuestID        int    `json:"quest_id"`
	OrganizationID int    `json:"organisation_id"`
	Description    string `json:"description,omitempty"`
	Role           string `json:"role,omitempty"`
	IsPrivate      bool   `json:"is_private,omitempty"`
}

// QuestOrganizationService handles communication with the QuestOrganization endpoint.
type QuestOrganizationService service

// Index returns the list of all QuestOrganizations for the quest associated with
// qstID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return QuestOrganizations that have
// been changed since that time.
func (qs *QuestOrganizationService) Index(campID int, qstID int, sync *time.Time) ([]*QuestOrganization, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuest)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*QuestOrganization `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestOrganization Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the QuestOrganization associated with orgID for the quest associated
// with qstID from the Campaign associated with campID.
func (qs *QuestOrganizationService) Get(campID int, qstID int, orgID int) (*QuestOrganization, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuest)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(orgID); err != nil {
		return nil, fmt.Errorf("invalid QuestOrganization ID: %w", err)
	}

	var wrap struct {
		Data *QuestOrganization `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestOrganization (ID: %d) from Campaign (ID: %d): %w", orgID, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing QuestOrganization associated with orgID for the quest
// associated with qstID from the Campaign associated with campID using the
// provided SimpleQuestOrganization data.
// Update returns the newly updated QuestOrganization.
func (qs *QuestOrganizationService) Update(campID int, qstID int, orgID int, org SimpleQuestOrganization) (*QuestOrganization, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuest)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(orgID); err != nil {
		return nil, fmt.Errorf("invalid QuestOrganization ID: %w", err)
	}

	b, err := json.Marshal(org)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestOrganization: %w", err)
	}

	var wrap struct {
		Data *QuestOrganization `json:"data"`
	}

	if err = qs.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update QuestOrganization for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing QuestOrganization associated with orgID from the
// Campaign associated with campID.
func (qs *QuestOrganizationService) Delete(campID int, qstID int, orgID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuest)

	if end, err = end.id(qstID); err != nil {
		return fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(orgID); err != nil {
		return fmt.Errorf("invalid QuestOrganization ID: %w", err)
	}

	if err = qs.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete QuestOrganization (ID: %d) for Campaign (ID: %d): %w", orgID, campID, err)
	}

	return nil
}
