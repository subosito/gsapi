package gsapi

import (
	check "gopkg.in/check.v1"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

func init() {
	check.Suite(&APISuite{})
}

type APISuite struct {
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
}

func (s *APISuite) SetUpTest(c *check.C) {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)
	s.client = NewClient(nil)
	s.client.url, _ = url.Parse(s.server.URL)
}

func (s *APISuite) TearDownTest(c *check.C) {
	s.server.Close()
}

func (s *APISuite) TestNewClient_customClient(c *check.C) {
	x := &http.Client{}
	t := NewClient(x)
	c.Assert(t.client, check.Equals, x)
}

func (s *APISuite) TestNewClient(c *check.C) {
	t := NewClient(nil)
	c.Assert(t.client, check.Equals, http.DefaultClient)
	c.Assert(t.url.Scheme, check.Equals, "http")
	c.Assert(t.url.Host, check.Equals, "go-search.org")
	c.Assert(t.url.Path, check.Equals, "/api")
}

func (s *APISuite) TestClient_Package(c *check.C) {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out := `
{
    "Package": "github.com/subosito/gotenv",
    "Name": "gotenv",
    "StarCount": 31,
    "Synopsis": "Package gotenv provides functionality to dynamically load the environment variables",
    "Description": "Package gotenv provides functionality to dynamically load the environment variables",
    "Imported": [
        "github.com/KevDog/go-stormpath",
        "github.com/MongoHQ/forego",
        "github.com/davidpelaez/forego",
        "github.com/ddollar/forego",
        "github.com/jweslley/forego",
        "github.com/jwilder/forego",
        "github.com/vendion/ssh-manage"
    ],
    "Imports": null,
    "ProjectURL": "https://github.com/subosito/gotenv",
    "StaticRank": 2373
}`

		io.WriteString(w, out)
	})

	pkg, err := s.client.Package("github.com/subosito/gotenv")
	c.Assert(err, check.IsNil)
	c.Assert(pkg.Package, check.Equals, "github.com/subosito/gotenv")
	c.Assert(pkg.Name, check.Equals, "gotenv")
	c.Assert(pkg.StarCount, check.Equals, 31)
	c.Assert(pkg.StaticRank, check.Equals, 2373)
}
