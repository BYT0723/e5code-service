// Code generated by goctl. DO NOT EDIT!
// Source: deploy.proto

package server

import (
	"context"

	"e5Code-Service/service/project/rpc/internal/logic"
	"e5Code-Service/service/project/rpc/internal/svc"
	"e5Code-Service/service/project/rpc/project"
)

type DeployServer struct {
	svcCtx *svc.ServiceContext
}

func NewDeployServer(svcCtx *svc.ServiceContext) *DeployServer {
	return &DeployServer{
		svcCtx: svcCtx,
	}
}

func (s *DeployServer) GetDepoly(ctx context.Context, in *project.GetDeployReq) (*project.GetDeployRsp, error) {
	l := logic.NewGetDepolyLogic(ctx, s.svcCtx)
	return l.GetDepoly(in)
}

func (s *DeployServer) AddDeploy(ctx context.Context, in *project.AddDeployReq) (*project.AddDeployRsp, error) {
	l := logic.NewAddDeployLogic(ctx, s.svcCtx)
	return l.AddDeploy(in)
}

func (s *DeployServer) UpdateDeploy(ctx context.Context, in *project.UpdateDeployReq) (*project.UpdateDeployRsp, error) {
	l := logic.NewUpdateDeployLogic(ctx, s.svcCtx)
	return l.UpdateDeploy(in)
}

func (s *DeployServer) DeleteDeploy(ctx context.Context, in *project.DeleteDeployReq) (*project.DeleteDeployRsp, error) {
	l := logic.NewDeleteDeployLogic(ctx, s.svcCtx)
	return l.DeleteDeploy(in)
}

func (s *DeployServer) Deploy(ctx context.Context, in *project.DeployReq) (*project.DeployRsp, error) {
	l := logic.NewDeployLogic(ctx, s.svcCtx)
	return l.Deploy(in)
}