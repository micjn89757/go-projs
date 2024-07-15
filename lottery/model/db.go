package model

import (
	"fmt"
	"lottery/utils"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	lotteryDB *gorm.DB
	lotteryDBOnce sync.Once 


	lotteryRedis *redis.Client
	lotteryRedisOnce sync.Once
	
)

func InitDB() {
	getLotteryDBConnection()
	
	// 将Mysql中的库存同步到redis
	InitInventory()
}


// TODO: 使用日志
func createMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		utils.Conf.Mysql.DBUser,
		utils.Conf.Mysql.DBPassword,
		utils.Conf.Mysql.DBHost,
		utils.Conf.Mysql.DBPort,
		utils.Conf.Mysql.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,

		// 禁用默认事务(提高运行速度)
		SkipDefaultTransaction: true,

		// 使用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})


	// TODO： 换成日志
	if err != nil {
		fmt.Printf("connect mysql failed: %s", err.Error())
		os.Exit(1)
	}


	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("init connection pool failed: %s", err.Error())
		os.Exit(1)
	}

	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置数据库最大连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db
}

// 单例
func getLotteryDBConnection() {
	lotteryDBOnce.Do(func ()  {
		if lotteryDB == nil {
			lotteryDB = createMysqlDB()
		}
	})
}


func createRedisClient(addr, passwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: passwd,
		DB:	db,
	})


	if err := client.Ping().Err(); err != nil {
		panic("connect to redis failed")
	} else {
		fmt.Printf("connect to redis %d", db)
	}

	return client
}


func GetRedisClient() *redis.Client {
	lotteryRedisOnce.Do(func ()  {
		if lotteryRedis == nil {
			lotteryRedis = createRedisClient(utils.Conf.Redis.Addr, "", utils.Conf.Redis.DB)
		}
	})
}