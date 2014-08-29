package gsapi

import (
	check "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

func init() {
	check.Suite(&APISuite{})
}

type APISuite struct{}

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
