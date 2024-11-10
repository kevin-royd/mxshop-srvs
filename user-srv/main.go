package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop-srvs/user-srv/global"
	"mxshop-srvs/user-srv/handler"
	"mxshop-srvs/user-srv/initialize"
	"mxshop-srvs/user-srv/proto"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1、初始化zap
	initialize.InitLogger()

	// 2.获取配置
	initialize.InitConfig()
	// 3、mysql初始化
	initialize.InitMysql()

	// 4、grpc注册
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.ServerConf.ServerInfo.Host, global.ServerConf.ServerInfo.Port))
	if err != nil {
		panic(err)
	}

	// 5、注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 6、注册consul
	initialize.InitConsul()

	// 7、 服务发现
	go func() {
		zap.S().Infof("gRPC 服务器成功启动在 %s:%d", global.ServerConf.ServerInfo.Host, global.ServerConf.ServerInfo.Port)
		err = server.Serve(listen)
		if err != nil {
			zap.S().Panicw("grpc server start error", zap.Error(err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// 退出时注销服务
	err = global.Consul.Agent().ServiceDeregister(global.ServerConf.ConsulInfo.Name)
	if err != nil {
		zap.S().Errorw("Failed to deregister service", zap.Error(err))
	} else {
		zap.S().Info("Service deregistered successfully")
	}
}
