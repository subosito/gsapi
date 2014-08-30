// Package gsapi provides simple wrapper for Go Search API (http://go-search.org/)
package gsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// BaseURL is the base endpoint for Go Search API
const BaseURL = "http://go-search.org/api"

// Client manages communication with Go Search API
type Client struct {
	url    *url.URL
	client *http.Client
}

// NewClient returns a new Go Search Client. This will load http.DefaultClient if httpClient is nil.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	return &Client{url: u, client: httpClient}
}

// GetRequest constructs http.Request which points to Go Search API URL.
func (c *Client) GetRequest(v url.Values) (*http.Request, error) {
	u := c.url
	u.RawQuery = v.Encode()
	return http.NewRequest("GET", u.String(), nil)
}

// Do performs actual request and decode json output as proper struct object.
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func (c *Client) perform(uv url.Values, v interface{}) error {
	req, _ := c.GetRequest(uv)
	_, err := c.Do(req, v)
	return err
}

// Package returns information for particular package.
func (c *Client) Package(pkg string) (Package, error) {
	p := Package{}
	d := url.Values{}
	d.Set("action", "package")
	d.Set("id", pkg)

	err := c.perform(d, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// Tops returns best packages based on categories.
func (c *Client) Tops() ([]Top, error) {
	t := []Top{}
	d := url.Values{}
	d.Set("action", "tops")
	d.Set("len", "100")

	err := c.perform(d, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// Packages returns ID of all registered packages.
func (c *Client) Packages() (Packages, error) {
	s := Packages{}
	d := url.Values{}
	d.Set("action", "packages")

	err := c.perform(d, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}

// Search returns Result based on q search results.
func (c *Client) Search(q string) (Result, error) {
	r := Result{}
	d := url.Values{}
	d.Set("action", "search")
	d.Set("q", q)

	err := c.perform(d, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
