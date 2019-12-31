package kanka

import "fmt"

// Profile provides simple data about the current user.
// For more information, visit: https://kanka.io/en-US/docs/1.0/profile
type Profile struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	AvatarThumb       string `json:"avatar_thumb"`
	Locale            string `json:"locale"`
	Timezone          string `json:"timezone"`
	DateFormat        string `json:"date_format"`
	DefaultPagination int    `json:"default_pagination"`
	LastCampaignID    int    `json:"last_campaign_id"`
	IsPatreon         bool   `json:"is_patreon"`
}

// ProfileService handles communication with the Profile endpoint.
type ProfileService service

// Get returns the Profile of the current user.
func (ps *ProfileService) Get() (*Profile, error) {
	var wrap struct {
		Data Profile `json:"data"`
	}

	err := ps.client.get(ps.end, &wrap)
	if err != nil {
		return nil, fmt.Errorf("cannot get Profile: %w", err)
	}

	return &wrap.Data, nil
}
