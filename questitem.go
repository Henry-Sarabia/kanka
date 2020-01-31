package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// QuestItem contains information about a specific quest item.
// For more information, visit: https://kanka.io/en-US/docs/1.0/quests#quest-items
type QuestItem struct {
	SimpleQuestItem
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleQuestItem contains only the simple information about a quest item.
// SimpleQuestItem is primarily used to create new quest items for posting to Kanka.
type SimpleQuestItem struct {
	QuestID     int    `json:"quest_id"`
	ItemID      int    `json:"item_id"`
	Description string `json:"description,omitempty"`
	Role        string `json:"role,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
}

// QuestItemService handles communication with the QuestItem endpoint.
type QuestItemService service

// Index returns the list of all QuestItems for the quest associated with
// qstID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return QuestItems that have
// been changed since that time.
func (qs *QuestItemService) Index(campID int, qstID int, sync *time.Time) ([]*QuestItem, error) {
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
		Data []*QuestItem `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestItem Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the QuestItem associated with itemID for the quest associated
// with qstID from the Campaign associated with campID.
func (qs *QuestItemService) Get(campID int, qstID int, itemID int) (*QuestItem, error) {
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

	if end, err = end.id(itemID); err != nil {
		return nil, fmt.Errorf("invalid QuestItem ID: %w", err)
	}

	var wrap struct {
		Data *QuestItem `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestItem (ID: %d) from Campaign (ID: %d): %w", itemID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new QuestItem for the quest associated with qstID in the
// Campaign associated with campID using the provided SimpleQuestItem data.
// Create returns the newly created QuestItem.
func (qs *QuestItemService) Create(campID int, qstID int, item SimpleQuestItem) (*QuestItem, error) {
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

	b, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestItem: %w", err)
	}

	var wrap struct {
		Data *QuestItem `json:"data"`
	}

	if err = qs.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create QuestItem for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing QuestItem associated with itemID for the quest
// associated with qstID from the Campaign associated with campID using the
// provided SimpleQuestItem data.
// Update returns the newly updated QuestItem.
func (qs *QuestItemService) Update(campID int, qstID int, itemID int, item SimpleQuestItem) (*QuestItem, error) {
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

	if end, err = end.id(itemID); err != nil {
		return nil, fmt.Errorf("invalid QuestItem ID: %w", err)
	}

	b, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestItem: %w", err)
	}

	var wrap struct {
		Data *QuestItem `json:"data"`
	}

	if err = qs.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update QuestItem for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing QuestItem associated with itemID from the
// Campaign associated with campID.
func (qs *QuestItemService) Delete(campID int, qstID int, itemID int) error {
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

	if end, err = end.id(itemID); err != nil {
		return fmt.Errorf("invalid QuestItem ID: %w", err)
	}

	if err = qs.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete QuestItem (ID: %d) for Campaign (ID: %d): %w", itemID, campID, err)
	}

	return nil
}
