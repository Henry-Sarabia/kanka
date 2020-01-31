package kanka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Henry-Sarabia/blank"
)

// EntityNote contains information about a specific entity note.
// For more information, visit: https://kanka.io/en-US/docs/1.0/entity-notes
// EntityNote represents a note relating to the parent entity.
type EntityNote struct {
	SimpleEntityNote
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	ID        int       `json:"id"`
	IsPrivate bool      `json:"is_private"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

// SimpleEntityNote contains only the simple information about an entity note.
// SimpleEntityNote is primarily used to create new entity notes for posting to Kanka.
type SimpleEntityNote struct {
	Name       string `json:"name"`
	EntityID   int    `json:"entity_id"`
	Entry      string `json:"entry,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

// MarshalJSON marshals the SimpleEntityNote into its JSON-encoded form if it
// has the required populated fields.
func (se SimpleEntityNote) MarshalJSON() ([]byte, error) {
	if blank.Is(se.Name) {
		return nil, fmt.Errorf("cannot marshal SimpleEntityNote into JSON with a missing Name")
	}

	type alias SimpleEntityNote
	return json.Marshal(alias(se))
}

// EntityNotes wraps a list of entity notes.
// EntityNotes exists to satisfy the API's JSON structure.
type EntityNotes struct {
	Data []EntityNote `json:"data"`
	Sync time.Time    `json:"sync"`
}

// EntityNoteService handles communication with the EntityNote endpoint.
type EntityNoteService service

// Index returns the list of all EntityNotes for the entity associated with
// entID in the Campaign associated with campID.
// If a non-nil time is provided, Index will only return EntityNotes that have
// been changed since that time.
func (es *EntityNoteService) Index(campID int, entID int, sync *time.Time) ([]*EntityNote, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if sync != nil {
		end = end.sync(*sync)
	}

	var wrap struct {
		Data []*EntityNote `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityNote Index from Campaign (ID: %d): %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the EntityNote associated with evtID for the entity associated
// with entID from the Campaign associated with campID.
func (es *EntityNoteService) Get(campID int, entID int, evtID int) (*EntityNote, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(evtID); err != nil {
		return nil, fmt.Errorf("invalid EntityNote ID: %w", err)
	}

	var wrap struct {
		Data *EntityNote `json:"data"`
	}

	if err = es.client.get(end, &wrap); err != nil {
		return nil, fmt.Errorf("cannot get EntityNote (ID: %d) from Campaign (ID: %d): %w", evtID, campID, err)
	}

	return wrap.Data, nil
}

// Create creates a new EntityNote for the entity associated with entID in the
// Campaign associated with campID using the provided SimpleEntityNote data.
// Create returns the newly created EntityNote.
func (es *EntityNoteService) Create(campID int, entID int, note SimpleEntityNote) (*EntityNote, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	b, err := json.Marshal(note)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityNote (Name: %s): %w", note.Name, err)
	}

	var wrap struct {
		Data *EntityNote `json:"data"`
	}

	if err = es.client.post(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot create EntityNote (Name: %s) for Campaign (ID: %d): %w", note.Name, campID, err)
	}

	return wrap.Data, nil
}

// Update updates an existing EntityNote associated with noteID for the entity
// associated with entID from the Campaign associated with campID using the
// provided SimpleEntityNote data.
// Update returns the newly updated EntityNote.
func (es *EntityNoteService) Update(campID int, entID int, noteID int, note SimpleEntityNote) (*EntityNote, error) {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return nil, fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(noteID); err != nil {
		return nil, fmt.Errorf("invalid EntityNote ID: %w", err)
	}

	b, err := json.Marshal(note)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal SimpleEntityNote (Name: %s): %w", note.Name, err)
	}

	var wrap struct {
		Data *EntityNote `json:"data"`
	}

	if err = es.client.put(end, bytes.NewReader(b), &wrap); err != nil {
		return nil, fmt.Errorf("cannot update EntityNote (Name: %s) for Campaign (ID: %d): '%w'", note.Name, campID, err)
	}

	return wrap.Data, nil
}

// Delete deletes an existing EntityNote associated with noteID from the
// Campaign associated with campID.
func (es *EntityNoteService) Delete(campID int, entID int, noteID int) error {
	var err error
	end := EndpointCampaign

	if end, err = end.id(campID); err != nil {
		return fmt.Errorf("invalid Campaign ID: %w", err)
	}
	end = end.concat(endpointEntity)

	if end, err = end.id(entID); err != nil {
		return fmt.Errorf("invalid Entity ID: %w", err)
	}
	end = end.concat(es.end)

	if end, err = end.id(noteID); err != nil {
		return fmt.Errorf("invalid EntityNote ID: %w", err)
	}

	if err = es.client.delete(end); err != nil {
		return fmt.Errorf("cannot delete EntityNote (ID: %d) for Campaign (ID: %d): %w", noteID, campID, err)
	}

	return nil
}
