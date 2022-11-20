package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-single-demo/greet/internal/svc"
	"go-zero-single-demo/greet/internal/types"
	"go-zero-single-demo/greet/model"
	"google.golang.org/grpc/status"
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
	tUser := model.TUser{
		UserId:   req.Id,
		NickName: req.Name,
	}

	res, err := l.svcCtx.TUserModel.Insert(l.ctx, &tUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	tUser.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UserRegisterRequest{
		Id:      req.Id,
		Name:    req.Name,
		Age:     30,
		Address: "济南市三盛国际公园",
	}, nil
}
