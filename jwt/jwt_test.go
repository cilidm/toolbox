package jwt

import (
	"testing"
)

func TestNewClaims(t *testing.T) {
	c := NewClaims(1, "cilidm")
	token, err := c.Generate()
	t.Log(token,err)
	p,err := c.Parse(token)
	t.Log(p)
}
