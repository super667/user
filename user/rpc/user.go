package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/super667/user/user/rpc/internal/mqs"

	"github.com/super667/user/user/rpc/internal/config"
	authserviceServer "github.com/super667/user/user/rpc/internal/server/authservice"
	permserviceServer "github.com/super667/user/user/rpc/internal/server/permservice"
	roleserviceServer "github.com/super667/user/user/rpc/internal/server/roleservice"
	strategyserviceServer "github.com/super667/user/user/rpc/internal/server/strategyservice"
	userroleserviceServer "github.com/super667/user/user/rpc/internal/server/userroleservice"
	userserviceServer "github.com/super667/user/user/rpc/internal/server/userservice"
	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "user/rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterAuthServiceServer(grpcServer, authserviceServer.NewAuthServiceServer(ctx))
		user.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		user.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		user.RegisterPermServiceServer(grpcServer, permserviceServer.NewPermServiceServer(ctx))
		user.RegisterStrategyServiceServer(grpcServer, strategyserviceServer.NewStrategyServiceServer(ctx))
		user.RegisterUserRoleServiceServer(grpcServer, userroleserviceServer.NewUserRoleServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
		serviceGroup.Add(mq)
	}
	go serviceGroup.Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
