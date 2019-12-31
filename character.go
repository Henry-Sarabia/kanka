package kanka

import "fmt"

// Character provides simple data about a character.
// For more information, visit: https://kanka.io/en-US/docs/1.0/characters
type Character struct {
	Entity
	LocationID int         `json:"location_id"`
	Title      interface{} `json:"title"`
	Age        string      `json:"age"`
	Sex        string      `json:"sex"`
	RaceID     int         `json:"race_id"`
	Type       interface{} `json:"type"`
	FamilyID   int         `json:"family_id"`
	IsDead     bool        `json:"is_dead"`
	Traits     []Trait     `json:"traits"`
}

// Trait represents a character's personality or appearance detail.
type Trait struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Entry        string `json:"entry"`
	Section      string `json:"section"`
	IsPrivate    bool   `json:"is_private"`
	DefaultOrder int    `json:"default_order"`
}

// CharacterService handles communication with the Character endpoint.
type CharacterService service

// Index returns a list of all characters in the current campaign.
func (cs *CharacterService) Index() ([]*Character, error) {
	var wrap struct {
		Data []*Character `json:"data"`
	}

	err := cs.client.get(cs.end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Character index: %w", err)
	}

	return wrap.Data, nil
}
