package user

import (
	"context"
	"errors"
	"fmt"
	"go-zero-single-demo/greet/model"
	"strconv"

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
	userId, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, errors.New("传参失败！")
	}
	user, err := l.svcCtx.TUserModel.FindOne(l.ctx, userId)
	fmt.Println("dddddddddd===", user, err)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败！" + string(userId))
	}
	if user == nil {
		return nil, errors.New("用户不存在！")
	}

	return &types.UserInfoResponse{
		Id:      string(user.UserId),
		Name:    user.NickName,
		Age:     20,
		Address: "jinanshi",
	}, nil
}
