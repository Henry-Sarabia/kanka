package kanka

import (
	"fmt"
	"time"
)

// Campaign provides simple data about a campaign.
// For more information, visit: https://kanka.io/en-US/docs/1.0/campaigns
type Campaign struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Entry      string    `json:"entry"`
	Image      string    `json:"image"`
	ImageFull  string    `json:"image_full"`
	ImageThumb string    `json:"image_thumb"`
	IsPrivate  bool      `json:"is_private"`
	Visibility string    `json:"visibility"`
	Locale     string    `json:"locale"`
	EntityID   int       `json:"entity_id"`
	Tags       []int     `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
	Members    Members   `json:"members"`
}

// Members wraps a list of campaign members.
// Members exists to satisfy the API's JSON structure.
type Members struct {
	Data []*Member `json:"data"`
	Sync time.Time `json:"sync"`
}

// Member provides simple data about a member of a campaign.
type Member struct {
	ID   int  `json:"id"`
	User User `json:"user"`
}

// User provides simple data about a user.
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// Links provides paging data.
type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  interface{} `json:"next"`
}

// Meta provides basic information about its query.
type Meta struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	LastPage    int    `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

// CampaignService handles communication with the Campaign endpoint.
type CampaignService service

// Index returns a list of all the campaigns the user has access to.
func (cs *CampaignService) Index() ([]*Campaign, error) {
	var wrap struct {
		Data  []*Campaign `json:"data"`
		Links Links       `json:"links"`
		Meta  Meta        `json:"meta"`
		Sync  time.Time   `json:"sync"`
		//TODO: Implement paging.
	}

	err := cs.client.get(cs.end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Campaign index: %w", err)
	}

	return wrap.Data, nil
}

// Get returns the Campaign corresponding with the provided ID.
func (cs *CampaignService) Get(campID int) (*Campaign, error) {
	var wrap struct {
		Data *Campaign `json:"data"`
	}

	end, err := cs.end.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}

	err = cs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Campaign with ID '%d': %w", campID, err)
	}

	return wrap.Data, nil
}

const pathUsers string = "/users"

// Members returns a list of all members of the Campaign corresponding with the
// provided id.
func (cs *CampaignService) Members(campID int) ([]*Member, error) {
	var wrap Members

	end, err := cs.end.id(campID)
	if err != nil {
		return nil, fmt.Errorf("invalid Campaign ID: %w", err)
	}

	end = end.append(pathUsers)
	err = cs.client.get(end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Members from Campaign with ID '%d': %w", campID, err)
	}

	return wrap.Data, nil
}
