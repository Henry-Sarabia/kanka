package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Journal contains information about a specific journal.
// For more information, visit: https://kanka.io/en-US/docs/1.0/journals
type Journal struct {
	SimpleJournal
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

// SimpleJournal contains only the simple information about a journal.
// SimpleJournal is primarily used to create new journals for posting to Kanka.
type SimpleJournal struct {
	Name        string `json:"name"`
	Entry       string `json:"entry,omitempty"`
	Type        string `json:"type,omitempty"`
	Date        string `json:"date,omitempty"`
	LocationID  int    `json:"location_id,omitempty"`
	CharacterID int    `json:"character_id,omitempty"`
	Tags        []int  `json:"tags,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
	Image       string `json:"image,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleJournal into its JSON-encoded form if it
// has the required populated fields.
func (sj SimpleJournal) MarshalJSON() ([]byte, error) {
	if blank.Is(sj.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleJournal into JSON with a missing Name")
	}

	type alias SimpleJournal
	return json.Marshal(alias(sj))
}

// JournalService handles communication with the Journal endpoint.
type JournalService service

// Index returns the list of all Journals in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Journals that have
// been changed since that time.
func (js *JournalService) Index(campID int, sync *time.Time) ([]*Journal, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(js.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Journal `json:"data"`
	}

	err = js.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Journal Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Journal associated with jrnID from the Campaign
// associated with campID.
func (js *JournalService) Get(campID int, jrnID int) (*Journal, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(js.end)

	end, err = end.ID(jrnID)
	if err != nil {
		return nil, fmt.Errorf("invalid Journal ID: %w", err)
	}

	var wrap struct {
		Data *Journal `json:"data"`
	}

	err = js.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Journal (ID: %d) from Campaign (ID: %d): %w", jrnID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Journal in the Campaign associated with campID using
// the provided SimpleJournal data.
// Create returns the newly created Journal.
func (js *JournalService) Create(campID int, jrn SimpleJournal) (*Journal, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(js.end)

	b, err := json.Marshal(jrn)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleJournal (Name: %s): %w", jrn.Name, err)
	}

	var wrap struct {
		Data *Journal `json:"data"`
	}

	err = js.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Journal (Name: %s) for Campaign (ID: %d): %w", jrn.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Journal associated with jrnID from the
// Campaign associated with campID using the provided SimpleJournal data.
// Update returns the newly updated Journal.
func (js *JournalService) Update(campID int, jrnID int, jrn SimpleJournal) (*Journal, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(js.end)

	end, err = end.ID(jrnID)
	if err != nil {
		return nil, fmt.Errorf("invalid Journal ID: %w", err)
	}

	b, err := json.Marshal(jrn)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleJournal (Name: %s): %w", jrn.Name, err)
	}

	var wrap struct {
		Data *Journal `json:"data"`
	}

	err = js.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Journal (Name: %s) for Campaign (ID: %d): '%w'", jrn.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Journal associated with jrnID from the
// Campaign associated with campID.
func (js *JournalService) Delete(campID int, jrnID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(js.end)

	end, err = end.ID(jrnID)
	if err != nil {
		return fmt.Errorf("invalid Journal ID: %w", err)
	}

	err = js.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Journal (ID: %d) for Campaign (ID: %d): %w", jrnID, campID, err)
	}

	return nil
}
