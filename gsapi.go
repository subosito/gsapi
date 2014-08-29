package gsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const BaseURL = "http://go-search.org/api"

type Client struct {
	url    *url.URL
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	return &Client{url: u, client: httpClient}
}

func (c *Client) Package(pkg string) (Package, error) {
	p := Package{}
	d := url.Values{}
	d.Set("action", "package")
	d.Set("id", pkg)

	u := c.url
	u.RawQuery = d.Encode()

	resp, err := c.client.Get(u.String())
	if err != nil {
		return p, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

type Package struct {
	Package     string
	Name        string
	StarCount   int
	Synopsis    string
	Description string
	Imported    []string
	Imports     []string
	ProjectURL  string
	StaticRank  int
}
