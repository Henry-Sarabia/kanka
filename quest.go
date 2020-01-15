package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Quest contains information about a specific quest.
// For more information, visit: https://kanka.io/en-US/docs/1.0/quests
type Quest struct {
	SimpleQuest
	ID             int       `json:"id"`
	ImageFull      string    `json:"image_full"`
	ImageThumb     string    `json:"image_thumb"`
	HasCustomImage bool      `json:"has_custom_image"`
	EntityID       int       `json:"entity_id"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      int       `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      int       `json:"updated_by"`
	Characters     int       `json:"characters"`
	Locations      int       `json:"locations"`

	Attributes   Attributes   `json:"attributes"`
	EntityEvents EntityEvents `json:"entity_events"`
	EntityFiles  EntityFiles  `json:"entity_files"`
	EntityNotes  EntityNotes  `json:"entity_notes"`
	Relations    Relations    `json:"relations"`
	Inventory    Inventory    `json:"inventory"`
}

// SimpleQuest contains only the simple information about a quest.
// SimpleQuest is primarily used to create new quests for posting to Kanka.
type SimpleQuest struct {
	Name        string `json:"name"`
	Entry       string `json:"entry,omitempty"`
	Type        string `json:"type,omitempty"`
	QuestID     int    `json:"quest_id,omitempty"`
	CharacterID int    `json:"character_id,omitempty"`
	Tags        []int  `json:"tags,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
	IsCompleted bool   `json:"is_completed,omitempty"`
	Image       string `json:"image,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleQuest into its JSON-encoded form if it
// has the required populated fields.
func (sq SimpleQuest) MarshalJSON() ([]byte, error) {
	if blank.Is(sq.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleQuest into JSON with a missing Name")
	}

	type alias SimpleQuest
	return json.Marshal(alias(sq))
}

// QuestService handles communication with the Quest endpoint.
type QuestService service

// Index returns the list of all Quests in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Quests that have
// been changed since that time.
func (qs *QuestService) Index(campID int, sync *time.Time) ([]*Quest, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(qs.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Quest `json:"data"`
	}

	err = qs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Quest Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Quest associated with qstID from the Campaign
// associated with campID.
func (qs *QuestService) Get(campID int, qstID int) (*Quest, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(qs.end)

	end, err = end.ID(qstID)
	if err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}

	var wrap struct {
		Data *Quest `json:"data"`
	}

	err = qs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Quest (ID: %d) from Campaign (ID: %d): %w", qstID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Quest in the Campaign associated with campID using
// the provided SimpleQuest data.
// Create returns the newly created Quest.
func (qs *QuestService) Create(campID int, qst SimpleQuest) (*Quest, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(qs.end)

	b, err := json.Marshal(qst)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuest (Name: %s): %w", qst.Name, err)
	}

	var wrap struct {
		Data *Quest `json:"data"`
	}

	err = qs.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Quest (Name: %s) for Campaign (ID: %d): %w", qst.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Quest associated with qstID from the
// Campaign associated with campID using the provided SimpleQuest data.
// Update returns the newly updated Quest.
func (qs *QuestService) Update(campID int, qstID int, qst SimpleQuest) (*Quest, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(qs.end)

	end, err = end.ID(qstID)
	if err != nil {
		return nil, fmt.Errorf("invalid Quest ID: %w", err)
	}

	b, err := json.Marshal(qst)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleQuest (Name: %s): %w", qst.Name, err)
	}

	var wrap struct {
		Data *Quest `json:"data"`
	}

	err = qs.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Quest (Name: %s) for Campaign (ID: %d): '%w'", qst.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Quest associated with qstID from the
// Campaign associated with campID.
func (qs *QuestService) Delete(campID int, qstID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(qs.end)

	end, err = end.ID(qstID)
	if err != nil {
		return fmt.Errorf("invalid Quest ID: %w", err)
	}

	err = qs.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Quest (ID: %d) for Campaign (ID: %d): %w", qstID, campID, err)
	}

	return nil
}
