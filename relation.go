package kanka

import "time"

// Relations wraps a list of relationships.
// Relations exists to satisfy the API's JSON structure.
type Relations struct {
	Data []Relation `json:"data"`
	Sync time.Time  `json:"sync"`
}

// Relation represents a relationship between two entities.
type Relation struct {
	OwnerID   int       `json:"owner_id"`
	TargetID  int       `json:"target_id"`
	Relation  string    `json:"relation"`
	IsPrivate bool      `json:"is_private"`
	Attitude  int       `json:"attitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
