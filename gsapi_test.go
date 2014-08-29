package gsapi

import (
	"bytes"
	"encoding/json"
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
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
	out := s.loadFixture("package.json")

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(out)
	})

	pkg, err := s.client.Package("github.com/subosito/gotenv")
	c.Assert(err, check.IsNil)

	want := Package{}
	err = s.jsonDecode(out, &want)
	c.Assert(err, check.IsNil)
	c.Assert(pkg, check.DeepEquals, want)
}

func (s *APISuite) TestClient_Tops(c *check.C) {
	out := s.loadFixture("tops.json")

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(out)
	})

	pkg, err := s.client.Tops()
	c.Assert(err, check.IsNil)

	want := []Top{}
	err = s.jsonDecode(out, &want)
	c.Assert(err, check.IsNil)
	c.Assert(pkg, check.DeepEquals, want)
}

func (s *APISuite) jsonDecode(b []byte, v interface{}) error {
	return json.NewDecoder(bytes.NewReader(b)).Decode(v)
}

func (s *APISuite) loadFixture(name string) []byte {
	b, _ := ioutil.ReadFile(filepath.Join("fixtures", name))
	return b
}
