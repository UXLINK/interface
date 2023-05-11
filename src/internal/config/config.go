package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	NodeId  int64 `json:"NodeId,default=1"`
	JwtAuth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}
}
