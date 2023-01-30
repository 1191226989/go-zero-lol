package svc

import (
	"lol/internal/config"
	"lol/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	CheckPath rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		CheckPath: middleware.NewCheckPathMiddleware().Handle,
	}
}
