package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// Note contains information about a specific note.
// For more information, visit: https://kanka.io/en-US/docs/1.0/notes
type Note struct {
	SimpleNote
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

// SimpleNote contains only the simple information about a note.
// SimpleNote is primarily used to create new notes for posting to Kanka.
type SimpleNote struct {
	Name      string `json:"name"`
	Entry     string `json:"entry,omitempty"`
	Type      string `json:"type,omitempty"`
	Tags      []int  `json:"tags,omitempty"`
	IsPrivate bool   `json:"is_private,omitempty"`
	Image     string `json:"image,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

// MarshalJSON marshals the SimpleNote into its JSON-encoded form if it
// has the required populated fields.
func (sn SimpleNote) MarshalJSON() ([]byte, error) {
	if blank.Is(sn.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleNote into JSON with a missing Name")
	}

	type alias SimpleNote
	return json.Marshal(alias(sn))
}

// NoteService handles communication with the Note endpoint.
type NoteService service

// Index returns the list of all Notes in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return Notes that have
// been changed since that time.
func (ns *NoteService) Index(campID int, sync *time.Time) ([]*Note, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ns.end)

	if sync != nil {
		end = end.Sync(*sync)
	}

	var wrap struct {
		Data []*Note `json:"data"`
	}

	err = ns.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Note Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Note associated with noteID from the Campaign
// associated with campID.
func (ns *NoteService) Get(campID int, noteID int) (*Note, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ns.end)

	end, err = end.ID(noteID)
	if err != nil {
		return nil, fmt.Errorf("invalid Note ID: %w", err)
	}

	var wrap struct {
		Data *Note `json:"data"`
	}

	err = ns.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Note (ID: %d) from Campaign (ID: %d): %w", noteID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new Note in the Campaign associated with campID using
// the provided SimpleNote data.
// Create returns the newly created Note.
func (ns *NoteService) Create(campID int, note SimpleNote) (*Note, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ns.end)

	b, err := json.Marshal(note)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleNote (Name: %s): %w", note.Name, err)
	}

	var wrap struct {
		Data *Note `json:"data"`
	}

	err = ns.client.post(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot create Note (Name: %s) for Campaign (ID: %d): %w", note.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing Note associated with noteID from the
// Campaign associated with campID using the provided SimpleNote data.
// Update returns the newly updated Note.
func (ns *NoteService) Update(campID int, noteID int, note SimpleNote) (*Note, error) {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ns.end)

	end, err = end.ID(noteID)
	if err != nil {
		return nil, fmt.Errorf("invalid Note ID: %w", err)
	}

	b, err := json.Marshal(note)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleNote (Name: %s): %w", note.Name, err)
	}

	var wrap struct {
		Data *Note `json:"data"`
	}

	err = ns.client.put(end, bytes.NewReader(b), &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot update Note (Name: %s) for Campaign (ID: %d): '%w'", note.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing Note associated with noteID from the
// Campaign associated with campID.
func (ns *NoteService) Delete(campID int, noteID int) error {
	end, err := EndpointCampaign.ID(campID)
	if err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.Concat(ns.end)

	end, err = end.ID(noteID)
	if err != nil {
		return fmt.Errorf("invalid Note ID: %w", err)
	}

	err = ns.client.delete(end)
	if err != nil {
		return fmt.Errorf("cannot delete Note (ID: %d) for Campaign (ID: %d): %w", noteID, campID, err)
	}

	return nil
}
