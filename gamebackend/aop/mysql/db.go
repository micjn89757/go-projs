package mysql

import (
	"fmt"
	"gamebackend/aop/logger"
	"gamebackend/common"
	"sync"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mysqlClient *gorm.DB
	mysqlClientOnce sync.Once

	// redisClient	*redis.Client
	// redisClientOnce sync.Once
)


func createMysqlDB(dbName, host, user, password string, port int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)	// mb4兼容emoji表情符合

	var err error 
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DontSupportRenameIndex: true,  // 重命名索引时采用删除并新建的方式
	}), &gorm.Config{
		PrepareStmt: true,	// 启用PrepareStmt, sql预编译，提高查询效率
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,	// 单数表名
		},
	})

	if err != nil {
		logger.Logger.Panic("connect to mysql failed", zap.String("dsn", dsn), zap.Error(err))
	}
	// 设置数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	logger.Logger.Info("connect to mysql db", zap.String("dbname", dbName))
	return db
}


// 单例
func GetMysqlDBConnection() *gorm.DB {
	mysqlClientOnce.Do(func ()  {
		if mysqlClient == nil {
			dbName := "database"
			viper := common.CreateConfig("mysql")
			host := viper.GetString(dbName + ".host")
			port := viper.GetInt(dbName + ".port")
			user := viper.GetString(dbName + ".user")
			password := viper.GetString(dbName + ".password")
			mysqlClient = createMysqlDB(dbName, host, user, password, port)
		}
	})
	return mysqlClient
}