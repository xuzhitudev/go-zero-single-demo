package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-single-demo/greet/internal/config"
	"go-zero-single-demo/greet/model"
)

type ServiceContext struct {
	Config     config.Config
	TUserModel model.TUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TUserModel: model.NewTUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
