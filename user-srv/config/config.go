package config

type ServerConfig struct {
	ServerInfo ServerCfg    `mapstructure:"server"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul"`
}

// server
type ServerCfg struct {
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
	Host string   `mapstructure:"host"`
	Port int      `mapstructure:"port"`
	Name string   `mapstructure:"name"`
	Tag  []string `mapstructure:"tag"`
}
