package user

import (
	"context"

	"go-zero-single-demo/greet/internal/svc"
	"go-zero-single-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetUserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetUserRegisterLogic {
	return &GreetUserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetUserRegisterLogic) GreetUserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterRequest, err error) {
	// todo: add your logic here and delete this line

	return &types.UserRegisterRequest{
		Id:      "ww",
		Name:    "许志图",
		Age:     30,
		Address: "济南市高新区铭盛大厦",
	}, nil
}
