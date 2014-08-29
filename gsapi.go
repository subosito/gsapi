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

func (c *Client) GetRequest(v url.Values) (*http.Request, error) {
	u := c.url
	u.RawQuery = v.Encode()
	return http.NewRequest("GET", u.String(), nil)
}

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

func (c *Client) Package(pkg string) (Package, error) {
	p := Package{}
	d := url.Values{}
	d.Set("action", "package")
	d.Set("id", pkg)

	req, _ := c.GetRequest(d)
	_, err := c.Do(req, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (c *Client) Tops() ([]Top, error) {
	t := []Top{}
	d := url.Values{}
	d.Set("action", "tops")
	d.Set("len", "100")

	req, _ := c.GetRequest(d)
	_, err := c.Do(req, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
