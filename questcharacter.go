package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// QuestCharacter contains information about a specific questcharacter.
// For more information, visit: https://kanka.io/en-US/docs/1.0/quests#quest-characters
type QuestCharacter struct {
	SimpleQuestCharacter
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleQuestCharacter contains only the simple information about a questcharacter.
// SimpleQuestCharacter is primarily used to create new questcharacters for posting to Kanka.
type SimpleQuestCharacter struct {
	QuestID     int    `json:"quest_id"`
	CharacterID int    `json:"character_id"`
	Description string `json:"description,omitempty"`
	Role        string `json:"role,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
}

// QuestCharacterService handles communication with the QuestCharacter endpoint.
type QuestCharacterService service

// Index returns the list of all QuestCharacters for the quest associated with
// qstID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return QuestCharacters that have
// been changed since that time.
func (qs *QuestCharacterService) Index(campID int, qstID int, sync *time.Time) ([]*QuestCharacter, error) {
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
		Data []*QuestCharacter `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestCharacter Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the QuestCharacter associated with qchID for the quest associated
// with qstID from the Campaign associated with campID.
func (qs *QuestCharacterService) Get(campID int, qstID int, qchID int) (*QuestCharacter, error) {
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

	if end, err = end.id(qchID); err != nil {
		return nil, fmt.Errorf("invalid QuestCharacter ID: %w", err)
	}

	var wrap struct {
		Data *QuestCharacter `json:"data"`
	}

	if err = qs.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get QuestCharacter (ID: %d) from Campaign (ID: %d): %w", qchID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new QuestCharacter for the quest associated with qstID in the
// Campaign associated with campID using the provided SimpleQuestCharacter data.
// Create returns the newly created QuestCharacter.
func (qs *QuestCharacterService) Create(campID int, qstID int, qch SimpleQuestCharacter) (*QuestCharacter, error) {
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

	b, err := json.Marshal(qch)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestCharacter: %w", err)
	}

	var wrap struct {
		Data *QuestCharacter `json:"data"`
	}

	if err = qs.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create QuestCharacter for Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing QuestCharacter associated with qchID for the quest
// associated with qstID from the Campaign associated with campID using the
// provided SimpleQuestCharacter data.
// Update returns the newly updated QuestCharacter.
func (qs *QuestCharacterService) Update(campID int, qstID int, qchID int, qch SimpleQuestCharacter) (*QuestCharacter, error) {
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

	if end, err = end.id(qchID); err != nil {
		return nil, fmt.Errorf("invalid QuestCharacter ID: %w", err)
	}

	b, err := json.Marshal(qch)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuestCharacter: %w", err)
	}

	var wrap struct {
		Data *QuestCharacter `json:"data"`
	}

	if err = qs.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update QuestCharacter for Campaign (ID: %d): '%w'", campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing QuestCharacter associated with qchID from the
// Campaign associated with campID.
func (qs *QuestCharacterService) Delete(campID int, qstID int, qchID int) error {
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

	if end, err = end.id(qchID); err != nil {
		return fmt.Errorf("invalid QuestCharacter ID: %w", err)
	}

	if err = qs.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete QuestCharacter (ID: %d) for Campaign (ID: %d): %w", qchID, campID, err)
	}

	return nil
}
