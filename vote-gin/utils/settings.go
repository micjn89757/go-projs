package utils

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

var Logger *zap.Logger

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
	Server   *server `toml:"server"`
	Database *database `toml:"database"`
}

type server struct {
	AppMode  string `toml:"AppMode"`
	HttpPort string	`toml:"HttpPort"`
	JwtKey   string	`toml:"JwtKey"`
}

type database struct {
	DB         string `toml:"DB"`
	DBHost     string `toml:"DBHost"`
	DBPort     int	`toml:"DBPort"`
	DBUser     string `toml:"DBUser"`
	DBPassWord string `toml:"DBPassword"`
	DBName     string `toml:"DBName"`
}

// 初始化读取config.toml
func init() {
	var err error
	conf := &config{}
	_, err = toml.DecodeFile("config/config.toml", conf)
	if err != nil {
		fmt.Printf("parse toml failed:%v", err)
		panic("配置文件读取错误")
	}
	LoadLogger(conf)
	LoadData(conf)
	LoadServer(conf)
}

// 加载数据库信息
func LoadData(conf *config) {
	DB = conf.Database.DB
	DBHost = conf.Database.DBHost
	DBPort = conf.Database.DBPort
	DBUser = conf.Database.DBUser
	DBPassWord = conf.Database.DBPassWord
	DBName = conf.Database.DBName
}

// LoadServer 加载服务器信息
func LoadServer(conf *config) {
	AppMode = conf.Server.AppMode
	HttpPort = conf.Server.HttpPort
	JwtKey = conf.Server.JwtKey
}

// LoadLogger 加载日志
func LoadLogger(conf *config) {
	var err error 
	Logger, err = zap.NewDevelopment()
	if err != nil {
		fmt.Printf("load logger failed: %v", err)
		panic("日志加载失败")
	}
}