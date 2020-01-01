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
	Traits     Traits      `json:"traits"`
}

// Traits wraps a list of character traits.
// Traits exists to satisfy the API's JSON structure.
type Traits struct {
	Data []*Trait `json:"data"`
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

// Index returns a list of all characters in the Campaign corresponding with
// the provided id.
func (cs *CharacterService) Index(campID int) ([]*Character, error) {
	var wrap struct {
		Data []*Character `json:"data"`
	}

	end := EndpointCampaign.ID(campID)
	end = end.Append("/" + string(cs.end))

	err := cs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Character index from Campaign with ID '%d': %w", campID, err)
	}

	return wrap.Data, nil
}

// Get returns the Character corresponding with the provided ID from the
// Campaign corresponding with the other provided ID.
func (cs *CharacterService) Get(campID int, charID int) (*Character, error) {
	var wrap struct {
		Data *Character `json:"data"`
	}

	end := EndpointCampaign.ID(campID)
	end = end.Append("/" + string(cs.end))
	end = end.ID(charID)

	err := cs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Character with ID '%d' from Campaign with ID '%d': %w", charID, campID, err)
	}

	return wrap.Data, nil
}
