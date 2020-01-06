package kanka

import "time"

// EntityFiles wraps a list of entity files.
// EntityFiles exists to satisfy the API's JSON structure.
type EntityFiles struct {
	Data []*EntityFile `json:"data"`
	Sync time.Time     `json:"sync"`
}

// EntityFile represents a specific calendar event relating to the parent
// entity.
type EntityFile struct {
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	EntityID   int       `json:"entity_id"`
	ID         int       `json:"id"`
	IsPrivate  bool      `json:"is_private"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Size       int       `json:"size"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
	Visibility string    `json:"visibility"`
}
