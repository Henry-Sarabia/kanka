package kanka

import "time"

// Attributes wraps a list of attributes.
// Attributes exists to satisfy the API's JSON structure.
type Attributes struct {
	Data []*Attribute `json:"data"`
	Sync time.Time    `json:"sync"`
}

// Attribute represents a distinct detail relating to the parent entity.
type Attribute struct {
	APIKey       string    `json:"api_key"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int       `json:"created_by"`
	DefaultOrder int       `json:"default_order"`
	EntityID     int       `json:"entity_id"`
	ID           int       `json:"id"`
	IsPrivate    bool      `json:"is_private"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int       `json:"updated_by"`
	Value        string    `json:"value"`
}
