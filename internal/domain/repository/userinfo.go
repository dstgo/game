package repository

import (
	"github.com/dstgo/kratosx"

	"github.com/dstgo/game/internal/domain/entity"
	"github.com/dstgo/game/internal/types"
)

type Userinfo interface {
	// ListUserinfo 获取用户扩展信息列表
	ListUserinfo(ctx kratosx.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, uint32, error)

	// UpdateUserinfo 更新用户扩展信息
	UpdateUserinfo(ctx kratosx.Context, req *entity.Userinfo) error

	// CheckKeywords 检查keyword是否合法
	CheckKeywords(ctx kratosx.Context, appId uint32, keywords []string) error

	// CreateUserinfo 创建用户扩展信息
	CreateUserinfo(ctx kratosx.Context, req *entity.Userinfo) (uint32, error)

	// DeleteUserinfo 删除用户扩展信息
	DeleteUserinfo(ctx kratosx.Context, id uint32) error
}
