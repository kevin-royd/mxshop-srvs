package config

// server
type ServerCfg struct {
	ServerInfo ServerConfig `mapstructure:"server"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

// MysqlConfig MysqlInfo
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbName"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
	Id   string
	Tag  []string `mapstructure:"tag"`
}
