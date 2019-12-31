package kanka

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const kankaURL string = "https://kanka.io/api/1.0/"

// service handles communication with a specific endpoint.
type service struct {
	client *Client
	end    endpoint
}

// Client wraps an HTTP client used to communicate with the Kanka API,
// the root URL of the Kanka API, and the user's Kanka OAuth key.
// Client holds all the separate services for each endpoint.
type Client struct {
	http    *http.Client
	rootURL string
	token   string

	// Services
	Profiles   *ProfileService
	Campaigns  *CampaignService
	Characters *CharacterService
}

// NewClient returns a new Client configured to communicate with the Kanka API.
// The provided token is used to make API calls on your behalf. If no custom
// HTTP client is provided, a default HTTP client is used instead.
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

	return c
}

// request returns an appropriately configured HTTP request with the provided
// method and endpoint.
func (c *Client) request(method string, end endpoint) (*http.Request, error) {
	url := c.rootURL + string(end)

	req, err := http.NewRequest(method, url, nil)
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

// get executes a GET request to the provided endpoint and stores the
// unmarshaled JSON result in the provided empty interface.
func (c *Client) get(end endpoint, result interface{}) error {
	req, err := c.request("GET", end)
	if err != nil {
		return err
	}

	err = c.send(req, result)
	if err != nil {
		return err
	}

	return nil
}
