package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// QuestLocation contains information about a specific quest location.
// For more information, visit: https://kanka.io/en-US/docs/1.0/quests#quest-locations
type QuestLocation struct {
	SimpleQuestLocation
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleQuestLocation contains only the simple information about a quest location.
// SimpleQuestLocation is primarily used to create new quest locations for posting to Kanka.
type SimpleQuestLocation struct {
	QuestID     int    `json:"quest_id"`
	LocationID  int    `json:"location_id"`
	Description string `json:"description,omitempty"`
	Role        string `json:"role,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
}

// QuestLocationService handles communication with the QuestLocation endpoint.
type QuestLocationService service

// Index returns the list of all QuestLocations for the quest associated with
// qstID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return QuestLocations that have
// been changed since that time.
func (qs *QuestLocationService) Index(campID int, qstID int, sync *time.Time) ([]*QuestLocation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuestLocation)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*QuestLocation `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestLocation Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the QuestLocation associated with qlocID for the quest associated
// with qstID from the Campaign associated with campID.
func (qs *QuestLocationService) Get(campID int, qstID int, qlocID int) (*QuestLocation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuestLocation)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(qlocID); err != nil {
		return nil, fmt.Errorf("invalid QuestLocation ID: %w", err)
	}

	var wrap struct {
		Data *QuestLocation `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestLocation (ID: %d) from Campaign (ID: %d): %w", qlocID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new QuestLocation for the quest associated with qstID in the
// Campaign associated with campID using the provided SimpleQuestLocation data.
// Create returns the newly created QuestLocation.
func (qs *QuestLocationService) Create(campID int, qstID int, qloc SimpleQuestLocation) (*QuestLocation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuestLocation)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	b, err := json.Marshal(qloc)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestLocation: %w", err)
	}

	var wrap struct {
		Data *QuestLocation `json:"data"`
	}

	if err = qs.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create QuestLocation for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing QuestLocation associated with qlocID for the quest
// associated with qstID from the Campaign associated with campID using the
// provided SimpleQuestLocation data.
// Update returns the newly updated QuestLocation.
func (qs *QuestLocationService) Update(campID int, qstID int, qlocID int, qloc SimpleQuestLocation) (*QuestLocation, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuestLocation)

	if end, err = end.id(qstID); err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(qlocID); err != nil {
		return nil, fmt.Errorf("invalid QuestLocation ID: %w", err)
	}

	b, err := json.Marshal(qloc)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestLocation: %w", err)
	}

	var wrap struct {
		Data *QuestLocation `json:"data"`
	}

	if err = qs.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update QuestLocation for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing QuestLocation associated with qlocID from the
// Campaign associated with campID.
func (qs *QuestLocationService) Delete(campID int, qstID int, qlocID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(EndpointQuestLocation)

	if end, err = end.id(qstID); err != nil {
		return fmt.Errorf("invalid Quest ID: %w", err)
	}
	end = end.concat(qs.end)

	if end, err = end.id(qlocID); err != nil {
		return fmt.Errorf("invalid QuestLocation ID: %w", err)
	}

	if err = qs.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete QuestLocation (ID: %d) for Campaign (ID: %d): %w", qlocID, campID, err)
	}

	return nil
}
