package user

import (
	"context"
	"errors"
	"go-zero-single-demo/greet/internal/svc"
	"go-zero-single-demo/greet/internal/types"
	"go-zero-single-demo/greet/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetUserInfoLogic {
	return &GreetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetUserInfoLogic) GreetUserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {

	user, err := l.svcCtx.TUserModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败！")
	}
	if user == nil {
		return nil, errors.New("用户不存在！")
	}

	return &types.UserInfoResponse{
		Id:      user.UserId,
		Name:    user.NickName,
		Age:     20,
		Address: "jinanshi",
	}, nil
}
