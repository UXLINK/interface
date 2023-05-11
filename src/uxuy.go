package main

import (
	"flag"
	"fmt"
	"net/http"

	"uxuy/src/internal/config"
	"uxuy/src/internal/handler"
	"uxuy/src/internal/svc"
	bizresponse "uxuy/src/util/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var configFile = flag.String("f", "etc/uxuy.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(&c)
	handler.RegisterHandlers(server, ctx)

	// 关闭心跳检测的日志
	logx.DisableStat()

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, any) {
		if e, ok := err.(*bizresponse.CodeError); ok {
			return http.StatusOK, e.Data()
		}
		if s, ok := status.FromError(err); ok {
			// grpc 通常都是100以内
			if s.Code() < 100 {
				switch s.Code() {
				case codes.Canceled:
					return http.StatusOK, bizresponse.ErrClientCancel.Data()
				}
				return http.StatusOK, bizresponse.ErrInternalFailed.Data()
			}
			return http.StatusOK, bizresponse.NewCodeError(int(s.Code()), s.Message()).Data()
		}
		return http.StatusOK, bizresponse.ErrUnknown.Data()
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
