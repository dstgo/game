package auth

import (
	"github.com/dstgo/kratosx"
)

type Auth struct {
	UserId     uint32 `json:"userId"`
	AppKeyword string `json:"appKeyword"`
	AppId      uint32 `json:"appId"`
}

func Get(ctx kratosx.Context) (*Auth, error) {
	data := Auth{}
	return &data, ctx.Authentication().ParseAuth(ctx, &data)
}
