package cookie

import (
	"github.com/cilidm/toolbox/session"
	gs "github.com/gorilla/sessions"
)

type Store interface {
	session.Store
}

func NewStore(keyPairs ...[]byte) Store {
	return &store{gs.NewCookieStore(keyPairs...)}
}

type store struct {
	*gs.CookieStore
}

func (c *store) Options(options session.Options) {
	c.CookieStore.Options = options.ToGorillaOptions()
}
