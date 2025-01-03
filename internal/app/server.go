package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/game/api/v1/mod"
	"github.com/dstgo/game/api/v1/server"
	"github.com/dstgo/game/internal/conf"
)

type Server struct {
	server.UnimplementedServerServer
}

func NewServer(c *conf.Config) *Server {
	return &Server{}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewMod(c)
		mod.RegisterModServer(gs, srv)
		mod.RegisterModHTTPServer(hs, srv)
	})
}
