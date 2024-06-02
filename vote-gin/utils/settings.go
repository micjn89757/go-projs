package utils

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

var logger *zap.Logger

var (
	AppMode string 
	HttpPort string 
	JwtKey string 

	DB string 
	DBHost string 
	DBPort int 
	DBUser string 
	DBPassWord string 
	DBName string 
)


type config struct {
	server   *server
	database *database
}

type server struct {
	appMode  string `toml:"AppMode"`
	httpPort string	`toml:"HttpPort"`
	jwtKey   string	`toml:"JwtKey"`
}

type database struct {
	db         string `toml:"DB"`
	dbHost     string `toml:"DBHost"`
	dbPort     int	`toml:"DBPort"`
	dbUser     string `toml:"DBUser"`
	dbPassWord string `toml:"DBPassword"`
	dbName     string `toml:"DBName"`
}

// 初始化读取config.toml
func init() {
	var err error
	conf := &config{}
	_, err = toml.DecodeFile("config/config.toml", conf)
	if err != nil {
		fmt.Errorf("配置文件读取错误: %w", err)
	}
	LoadLogger(conf)
	LoadData(conf)
	LoadServer(conf)
}

// 加载数据库信息
func LoadData(conf *config) {
	DB = conf.database.db
	DBHost = conf.database.dbHost
	DBPort = conf.database.dbPort
	DBUser = conf.database.dbUser
	DBPassWord = conf.database.dbPassWord
	DBName = conf.database.dbName
}

// LoadServer 加载服务器信息
func LoadServer(conf *config) {
	AppMode = conf.server.appMode
	HttpPort = conf.server.httpPort
	JwtKey = conf.server.jwtKey
}

// LoadLogger 加载日志
func LoadLogger(conf *config) {
	var err error 
	logger, err = zap.NewDevelopment()
	if err != nil {
		fmt.Errorf("日志加载失败：%w", err)
	}
}