package kanka

import "time"

// Inventories wraps a list of inventories.
// Inventories exists to satisfy the API's JSON structure.
type Inventories struct {
	Data []*Inventory `json:"data"`
	Sync time.Time    `json:"sync"`
}

// Inventory represents a single inventory belonging to the parent entity.
type Inventory struct {
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	EntityID   int       `json:"entity_id"`
	ID         int       `json:"id"`
	IsPrivate  bool      `json:"is_private"`
	ItemID     int       `json:"item_id"`
	Position   string    `json:"position"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
	Visibility string    `json:"visibility"`
}
