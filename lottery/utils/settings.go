package utils

import "github.com/BurntSushi/toml"

type Config struct {
	Mysql  *Mysql `toml:"mysql"`
	Server *Server `toml:"server"`
}

type Mysql struct {
	DBHost     string `toml:"db_host"`
	DBPort     string `toml:"db_port"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
	DBName     string `toml:"db_name"`
}

type Server struct {
	AppMode  string `toml:"app_mode"`
	HttpPort string	`toml:"http_port"`
	JwtKey   string `toml:"jwt_key"`
}

// 初始化配置
func init() {
	var config Config
	_, err := toml.DecodeFile("./config/config.toml", &config)


	if err != nil {
		panic(err)
	}
}
