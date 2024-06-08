package utils

import (
	"fmt"
	"os"
	"path/filepath"

	// "path/filepath"

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


type Config struct {
	Server   *Server `toml:"server"`
	Database *Database `toml:"database"`
}

type Server struct {
	AppMode  string `toml:"AppMode"`
	HttpPort string	`toml:"HttpPort"`
	JwtKey   string	`toml:"JwtKey"`
}

type Database struct {
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
	conf := &Config{}
	var configDir string

	// 获取当前项目根目录路径
	if ex, err := os.Getwd(); err == nil {
		if filepath.Base(ex) == "vote-gin" {
			configDir = filepath.Join(ex, "config/config.toml")
		}else{
			for filepath.Base(configDir) != "vote-gin"{
				configDir = filepath.Dir(ex)  
				fmt.Println(configDir)
			}

			if configDir == "" || configDir == "." {
				panic("cannot find config file path")
			}else {
				configDir = filepath.Join(configDir, "config/config.toml")
			}
		}
	}

	_, err = toml.DecodeFile(configDir, conf)
	if err != nil {
		fmt.Printf("parse toml failed:%v", err)
		panic("config file decode failed")
	}
	LoadLogger(conf)
	LoadData(conf)
	LoadServer(conf)
}

// 加载数据库信息
func LoadData(conf *Config) {
	DB = conf.Database.DB
	DBHost = conf.Database.DBHost
	DBPort = conf.Database.DBPort
	DBUser = conf.Database.DBUser
	DBPassWord = conf.Database.DBPassWord
	DBName = conf.Database.DBName
}

// LoadServer 加载服务器信息
func LoadServer(conf *Config) {
	AppMode = conf.Server.AppMode
	HttpPort = conf.Server.HttpPort
	JwtKey = conf.Server.JwtKey
}

// LoadLogger 加载日志
func LoadLogger(conf *Config) {
	var err error 
	Logger, err = zap.NewDevelopment()
	if err != nil {
		fmt.Printf("load logger failed: %v", err)
		panic("日志加载失败")
	}
}