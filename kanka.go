package kanka

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const kankaURL string = "https://kanka.io/api/1.0/"

// service handles communication with a specific endpoint.
type service struct {
	client *Client
	end    endpoint
}

// Client handles communication between the user and the Kanka API.
// Client requires a valid Kanka user's OAuth token to authenticate each
// request. Client contains separate services for each endpoint.
type Client struct {
	http    *http.Client
	rootURL string
	token   string

	// Services
	Profiles      *ProfileService
	Campaigns     *CampaignService
	Characters    *CharacterService
	Locations     *LocationService
	Families      *FamilyService
	Organizations *OrganizationService
	Items         *ItemService
	Notes         *NoteService
	Events        *EventService
	Races         *RaceService
	Quests        *QuestService
	Journals      *JournalService
	Tags          *TagService

	Attributes        *AttributeService
	EntityEvents      *EntityEventService
	EntityInventories *EntityInventoryService
}

// NewClient returns an appropriately configured Client using the provided
// OAuth token. A provided custom HTTP client can be used to make the API
// requests otherwise a default HTTP client will be used instead.
func NewClient(token string, custom *http.Client) *Client {
	if custom == nil {
		custom = http.DefaultClient
	}

	c := &Client{
		http:    custom,
		rootURL: kankaURL,
		token:   token,
	}

	c.Profiles = &ProfileService{client: c, end: EndpointProfile}
	c.Campaigns = &CampaignService{client: c, end: EndpointCampaign}
	c.Characters = &CharacterService{client: c, end: EndpointCharacter}
	c.Locations = &LocationService{client: c, end: EndpointLocation}
	c.Families = &FamilyService{client: c, end: EndpointFamily}
	c.Organizations = &OrganizationService{client: c, end: EndpointOrganization}
	c.Items = &ItemService{client: c, end: EndpointItem}
	c.Notes = &NoteService{client: c, end: EndpointNote}
	c.Events = &EventService{client: c, end: EndpointEvent}
	c.Races = &RaceService{client: c, end: EndpointRace}
	c.Quests = &QuestService{client: c, end: EndpointQuest}
	c.Journals = &JournalService{client: c, end: EndpointJournal}
	c.Tags = &TagService{client: c, end: EndpointTag}

	c.Attributes = &AttributeService{client: c, end: EndpointAttribute}
	c.EntityEvents = &EntityEventService{client: c, end: EndpointEntityEvent}
	c.EntityInventories = &EntityInventoryService{client: c, end: EndpointEntityInventory}

	return c
}

// request returns an appropriately configured HTTP request with the provided
// method, endpoint, and body.
func (c *Client) request(method string, end endpoint, body io.Reader) (*http.Request, error) {
	url := c.rootURL + string(end)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("cannot create request with method '%s' for url '%s': %w", method, url, err)
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// send executes the provided request and stores the unmarshaled JSON result in
// the provided empty interface.
func (c *Client) send(req *http.Request, result interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("http client cannot send request with method '%s' to url '%s': %w", req.Method, req.URL.String(), err)
	}
	defer resp.Body.Close()

	if !isSuccess(resp.StatusCode) {
		return &ServerError{code: resp.StatusCode, status: resp.Status, temporary: isTemporary(resp.StatusCode)}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read response body: %w", err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return fmt.Errorf("cannot unmarshal body data: %w", err)
	}

	return nil
}

const paramRelated string = "?related=1"

// get executes a GET request to the provided endpoint and stores the
// unmarshaled JSON result in the provided empty interface.
func (c *Client) get(end endpoint, result interface{}) error {
	end = end.append(paramRelated)

	req, err := c.request("GET", end, nil)
	if err != nil {
		return err
	}

	err = c.send(req, result)
	if err != nil {
		return err
	}

	return nil
}

// post executes a POST request to the provided endpoint with the provided body
// and stores the unmarshaled JSON result in the provided empty interface.
func (c *Client) post(end endpoint, body io.Reader, result interface{}) error {
	req, err := c.request("POST", end, body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	err = c.send(req, result)
	if err != nil {
		return err
	}

	return nil
}

// put executes a PUT request to the provided endpoint with the provided body
// and stores the unmarshaled JSON result in the provided empty interface.
func (c *Client) put(end endpoint, body io.Reader, result interface{}) error {
	req, err := c.request("PUT", end, body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	err = c.send(req, result)
	if err != nil {
		return err
	}

	return nil
}

// delete executes a DELETE request to the provided endpoint.
func (c *Client) delete(end endpoint) error {
	req, err := c.request("DELETE", end, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("http client cannot send request with method '%s' to url '%s': %w", req.Method, req.URL.String(), err)
	}
	defer resp.Body.Close()

	if !isSuccess(resp.StatusCode) {
		return &ServerError{code: resp.StatusCode, status: resp.Status, temporary: isTemporary(resp.StatusCode)}
	}

	return nil
}
