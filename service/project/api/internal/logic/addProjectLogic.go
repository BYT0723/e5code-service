package logic

import (
	"context"

	"e5Code-Service/common/errorx"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"e5Code-Service/service/project/rpc/project"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddProjectLogic {
	return AddProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProjectLogic) AddProject(req types.AddProjectReq) (*types.AddProjectReply, error) {
	rsp, err := l.svcCtx.ProjectRpc.AddProject(l.ctx, &project.AddProjectReq{
		Name:     req.Name,
		Desc:     req.Desc,
		Url:      req.Url,
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.Error("Fail to add Project, err: ", err.Error())
		return nil, errorx.NewCodeError(codesx.RPCError, err.Error())
	}
	return &types.AddProjectReply{
		ID: rsp.Id,
	}, nil
}
