package kanka

import (
	"fmt"
	"strconv"
)

type endpoint string

// Available Kanka API Endpoints
const (
	// Core Objects
	EndpointProfile      endpoint = "profile"
	EndpointCampaign     endpoint = "campaigns"
	EndpointCharacter    endpoint = "characters"
	EndpointLocation     endpoint = "locations"
	EndpointFamily       endpoint = "families"
	EndpointOrganisation endpoint = "organisations"
	EndpointItem         endpoint = "items"
	EndpointNote         endpoint = "notes"
	EndpointEvent        endpoint = "events"
	EndpointCalendar     endpoint = "calendars"
	EndpointRace         endpoint = "races"
	EndpointQuest        endpoint = "quests"
	EndpointJournal      endpoint = "journals"
	EndpointTag          endpoint = "tags"
	EndpointConversation endpoint = "conversations"
	EndpointDiceRoll     endpoint = "dice_rolls"

	// Entities
	EndpointAttribute         endpoint = "attributes"
	EndpointEntityEvent       endpoint = "entity_events"
	EndpointEntityFile        endpoint = "entity_files"
	EndpointEntityInventories endpoint = "inventories"
	EndpointEntityInventory   endpoint = "inventory"
	EndpointEntityNote        endpoint = "entity_notes"
	EndpointEntityTag         endpoint = "entity_tags"
	EndpointEntityRelation    endpoint = "relations"

	// Search
	EndpointSearch endpoint = "search"
)

// Append returns an endpoint appended with the provided string.
func (e endpoint) Append(s string) endpoint {
	return endpoint(string(e) + s)
}

// ID returns an endpoint appropriately formatted with the provided id.
func (e endpoint) ID(id int) (endpoint, error) {
	if id < 0 {
		return "", fmt.Errorf("provided ID (%d) cannot be negative", id)
	}

	return e.Append("/" + strconv.Itoa(id)), nil
}
