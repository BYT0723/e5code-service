package logic

import (
	"context"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/rpc/internal/svc"
	"e5Code-Service/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserReq) (*user.DeleteUserRsp, error) {
	u, _ := l.svcCtx.UserModel.FindOne(in.Id)
	err := l.svcCtx.UserModel.Delete(in.Id)
	if err != nil {
		logx.Errorf("Fail to delete user(%s), err: %v", in.Id, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, status.Error(codesx.NotFound, "UserNotFound")
		}
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	if res, err := l.svcCtx.GitCli.DestoryUser(u.Name); err != nil {
		logx.Error("Fail to DestoryGitUser on DeleteUser: ", err.Error())
		return nil, status.Error(codesx.GitError, res)
	}

	return &user.DeleteUserRsp{}, nil
}
