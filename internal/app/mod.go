package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/game/api/v1/mod"
	"github.com/dstgo/game/internal/conf"
)

type Mod struct {
	mod.UnimplementedModServer
}

func NewMod(c *conf.Config) *Mod {
	return &Mod{}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewMod(c)
		mod.RegisterModServer(gs, srv)
		mod.RegisterModHTTPServer(hs, srv)
	})
}
