package global

import (
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"mxshop-srvs/user-srv/config"
)

var (
	ServerConf *config.ServerConfig = &config.ServerConfig{}
	DB         *gorm.DB
	Consul     *api.Client
)
