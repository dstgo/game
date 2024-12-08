package app

import (
	"context"

	"github.com/dstgo/kratosx"
	"github.com/dstgo/kratosx/pkg/valx"
	ktypes "github.com/dstgo/kratosx/types"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/game/api/game/errors"

	pb "github.com/dstgo/game/api/game/feedback/v1"

	"github.com/dstgo/game/internal/conf"
	"github.com/dstgo/game/internal/domain/entity"
	"github.com/dstgo/game/internal/domain/service"
	"github.com/dstgo/game/internal/infra/dbs"
	"github.com/dstgo/game/internal/infra/rpc"

	"github.com/dstgo/game/internal/types"
)

type Feedback struct {
	pb.UnimplementedFeedbackServer
	srv *service.Feedback
}

func NewFeedback(conf *conf.Config) *Feedback {
	return &Feedback{
		srv: service.NewFeedback(
			conf,
			dbs.NewFeedback(),
			rpc.NewPermission(),
			rpc.NewFile(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewFeedback(c)
		pb.RegisterFeedbackHTTPServer(hs, srv)
		pb.RegisterFeedbackServer(gs, srv)
	})
}

// ListFeedbackCategory 获取反馈建议分类列表
func (fb *Feedback) ListFeedbackCategory(c context.Context, req *pb.ListFeedbackCategoryRequest) (*pb.ListFeedbackCategoryReply, error) {
	list, total, err := fb.srv.ListFeedbackCategory(kratosx.MustContext(c), &types.ListFeedbackCategoryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListFeedbackCategoryReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListFeedbackCategoryReply_FeedbackCategory{
			Id:        item.Id,
			Name:      item.Name,
			CreatedAt: uint32(item.CreatedAt),
		})
	}
	return &reply, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (fb *Feedback) CreateFeedbackCategory(c context.Context, req *pb.CreateFeedbackCategoryRequest) (*pb.CreateFeedbackCategoryReply, error) {
	id, err := fb.srv.CreateFeedbackCategory(kratosx.MustContext(c), &entity.FeedbackCategory{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedbackCategoryReply{Id: id}, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (fb *Feedback) UpdateFeedbackCategory(c context.Context, req *pb.UpdateFeedbackCategoryRequest) (*pb.UpdateFeedbackCategoryReply, error) {
	if err := fb.srv.UpdateFeedbackCategory(kratosx.MustContext(c), &entity.FeedbackCategory{
		CreateModel: ktypes.CreateModel{Id: req.Id},
		Name:        req.Name,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateFeedbackCategoryReply{}, nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (fb *Feedback) DeleteFeedbackCategory(c context.Context, req *pb.DeleteFeedbackCategoryRequest) (*pb.DeleteFeedbackCategoryReply, error) {
	if err := fb.srv.DeleteFeedbackCategory(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteFeedbackCategoryReply{}, nil
}

// ListFeedback 获取反馈建议列表
func (fb *Feedback) ListFeedback(c context.Context, req *pb.ListFeedbackRequest) (*pb.ListFeedbackReply, error) {
	list, total, err := fb.srv.ListFeedback(kratosx.MustContext(c), &types.ListFeedbackRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Order:      req.Order,
		OrderBy:    req.OrderBy,
		AppId:      req.AppId,
		CategoryId: req.CategoryId,
		Status:     req.Status,
		Platform:   req.Platform,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListFeedbackReply{Total: total}
	if err := valx.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateFeedback 创建反馈建议
func (fb *Feedback) CreateFeedback(c context.Context, req *pb.CreateFeedbackRequest) (*pb.CreateFeedbackReply, error) {
	id, err := fb.srv.CreateFeedback(kratosx.MustContext(c), &entity.Feedback{
		AppId:      req.AppId,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Content:    req.Content,
		Images:     req.Images,
		Contact:    req.Contact,
		Device:     req.Device,
		Platform:   req.Platform,
		Version:    req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedbackReply{Id: id}, nil
}

// DeleteFeedback 删除反馈建议
func (fb *Feedback) DeleteFeedback(c context.Context, req *pb.DeleteFeedbackRequest) (*pb.DeleteFeedbackReply, error) {
	if err := fb.srv.DeleteFeedback(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteFeedbackReply{}, nil
}

// UpdateFeedback 更新反馈建议
func (fb *Feedback) UpdateFeedback(c context.Context, req *pb.UpdateFeedbackRequest) (*pb.UpdateFeedbackReply, error) {
	if err := fb.srv.UpdateFeedback(kratosx.MustContext(c), &entity.Feedback{
		BaseModel:       ktypes.BaseModel{Id: req.Id},
		Status:          req.Status,
		ProcessedResult: req.ProcessedResult,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateFeedbackReply{}, nil
}
