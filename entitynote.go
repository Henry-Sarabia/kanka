package kanka

import "time"

// EntityNotes wraps a list of entity notes.
// EntityNotes exists to satisfy the API's JSON structure.
type EntityNotes struct {
	Data []EntityNote `json:"data"`
	Sync time.Time    `json:"sync"`
}

// EntityNote represents a note relating to the parent entity.
type EntityNote struct {
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	EntityID   int       `json:"entity_id"`
	Entry      string    `json:"entry"`
	ID         int       `json:"id"`
	IsPrivate  bool      `json:"is_private"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
	Visibility string    `json:"visibility"`
}
