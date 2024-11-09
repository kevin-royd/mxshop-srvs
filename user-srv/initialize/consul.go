package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"mxshop-srvs/user-srv/global"
)

// InitConsul 创建Consul客户端
func InitConsul() {
	// 实例化consul对象
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.ServerConf.ConsulInfo.Host, global.ServerConf.ConsulInfo.Port)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panicw("[InitConsul] init consul fail")

	}
	global.Consul = client
	registerService()
}

// registerService 将gRPC服务注册到consul
func registerService() {
	// 健康检查
	check := &api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d", global.ServerConf.ServerInfo.Host, global.ServerConf.ServerInfo.Port), //
		Timeout:  "5s",                                                                                       // 超时时间
		Interval: "5s",                                                                                       // 运行检查的频率
		// 指定时间后自动注销不健康的服务节点
		DeregisterCriticalServiceAfter: "15s",
	}

	// 注册consul中的信息
	registration := &api.AgentServiceRegistration{
		ID:      global.ServerConf.ConsulInfo.Name, // 服务唯一ID
		Name:    global.ServerConf.ConsulInfo.Name, // 服务名称
		Tags:    global.ServerConf.ConsulInfo.Tag,  // 为服务打标签
		Address: global.ServerConf.ServerInfo.Host,
		Port:    global.ServerConf.ServerInfo.Port,
		Check:   check,
	}
	err := global.Consul.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Panicw("[InitConsul] register service fail", err)
	}

}
