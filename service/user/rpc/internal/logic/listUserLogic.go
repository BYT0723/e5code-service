package logic

import (
	"context"

	"e5Code-Service/api/pb/user"
	"e5Code-Service/common/errorx/codesx"
	"e5Code-Service/service/user/model"
	"e5Code-Service/service/user/rpc/internal/svc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *user.ListUserReq) (*user.ListUserRsp, error) {
	us := []model.User{}
	if err := l.svcCtx.Db.Find(&us, "id in ?", in.Ids).Error; err != nil {
		logx.Error("Fail to Find User: ", err.Error())
		return nil, status.Error(codesx.SQLError, err.Error())
	}
	res := []*user.UserModel{}
	copier.Copy(&res, us)

	return &user.ListUserRsp{
		Count:  int64(len(us)),
		Result: res,
	}, nil
}
