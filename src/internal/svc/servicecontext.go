package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"log"
	"uxuy/src/internal/config"
	"uxuy/src/internal/middleware"
	"uxuy/src/internal/model"
	"uxuy/src/util/jwttoken"
)

type ServiceContext struct {
	Config    *config.Config
	JwtClient *jwttoken.JwtClient
	DappModel *model.DappModel
	UserModel *model.UserModel

	ApiKey rest.Middleware
}

func NewServiceContext(c *config.Config) *ServiceContext {

	dappModel, err := model.NewDappModel()
	if err != nil {
		log.Fatalf("NewDappModel init fail. err:%v", err)
		return nil
	}

	userModel, err := model.NewUserModel()
	if err != nil {
		log.Fatalf("NewUserModel init fail. err:%v", err)
		return nil
	}

	return &ServiceContext{
		Config:    c,
		JwtClient: jwttoken.NewJwtClient(c.JwtAuth.AccessSecret, c.JwtAuth.AccessExpire, c.JwtAuth.RefreshExpire),
		DappModel: dappModel,
		UserModel: userModel,
		ApiKey:    middleware.NewApiKeyMiddleware(dappModel).Handle,
	}
}
