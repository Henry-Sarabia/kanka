package kanka

// EntityTag represents a specific tag relating to the parent entity.
type EntityTag struct {
	ID       int `json:"id"`
	EntityID int `json:"entity_id"`
	TagID    int `json:"tag_id"`
}
