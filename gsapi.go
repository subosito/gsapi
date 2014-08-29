package gsapi

import (
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
