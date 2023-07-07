package main

import (
	"flag"
	"fmt"
	"github.com/super667/user/user/api/internal/config"
	"github.com/super667/user/user/api/internal/handler"
	"github.com/super667/user/user/api/internal/middleware"
	"github.com/super667/user/user/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var configFile = flag.String("f", "user/api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Use(middleware.TokenBlackList(ctx.AuthRpc))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6068", nil))
	}()

	server.Start()
}
