package user

import (
	"context"

	"go-zero-single-demo/greet/internal/svc"
	"go-zero-single-demo/greet/internal/types"

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
	// todo: add your logic here and delete this line

	return &types.UserInfoResponse{
		Id:      "1",
		Name:    "xuzhitu",
		Age:     20,
		Address: "jinanshi",
	}, nil
}
