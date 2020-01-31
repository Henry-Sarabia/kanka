package kanka

import (
	"fmt"
	"strconv"
	"time"
)

type endpoint string

// Available Kanka API Endpoints
const (
	// Core Objects
	EndpointProfile            endpoint = "profile"
	EndpointCampaign           endpoint = "campaigns"
	EndpointCharacter          endpoint = "characters"
	EndpointLocation           endpoint = "locations"
	EndpointMapPoint           endpoint = "map_points"
	EndpointFamily             endpoint = "families"
	EndpointOrganization       endpoint = "organisations"
	EndpointOrganizationMember endpoint = "organisation_members"
	EndpointItem               endpoint = "items"
	EndpointNote               endpoint = "notes"
	EndpointEvent              endpoint = "events"
	EndpointCalendar           endpoint = "calendars"
	EndpointRace               endpoint = "races"
	EndpointQuest              endpoint = "quests"
	EndpointQuestCharacters    endpoint = "quest_characters"
	EndpointQuestLocation      endpoint = "quest_locations"
	EndpointQuestItem          endpoint = "quest_items"
	EndpointJournal            endpoint = "journals"
	EndpointTag                endpoint = "tags"
	EndpointConversation       endpoint = "conversations"
	EndpointDiceRoll           endpoint = "dice_rolls"

	// Entities
	EndpointAttribute         endpoint = "attributes"
	EndpointEntityEvent       endpoint = "entity_events"
	EndpointEntityFile        endpoint = "entity_files"
	EndpointEntityInventories endpoint = "inventories"
	EndpointEntityInventory   endpoint = "inventory"
	EndpointEntityNote        endpoint = "entity_notes"
	EndpointEntityTag         endpoint = "entity_tags"
	EndpointRelation          endpoint = "relations"
	endpointEntity            endpoint = "entities"

	// Search
	EndpointSearch endpoint = "search"
)

// append returns an endpoint appended with the provided string.
func (e endpoint) append(s string) endpoint {
	return endpoint(string(e) + s)
}

// concat returns an endpoint appropriately concatenated with the provided
// endpoint.
func (e endpoint) concat(end endpoint) endpoint {
	return e.append("/" + string(end))
}

// id returns an endpoint appropriately formatted with the provided id.
func (e endpoint) id(id int) (endpoint, error) {
	if id < 0 {
		return "", fmt.Errorf("provided ID (%d) cannot be negative", id)
	}

	return e.append("/" + strconv.Itoa(id)), nil
}

// sync returns an endpoint appropriately formatted with the provided lastSync
// time.
func (e endpoint) sync(t time.Time) endpoint {
	return e.append("/?lastSync=" + t.Format(time.RFC3339))
}
