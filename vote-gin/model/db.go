package model

import (
	"fmt"
	"vote-gin/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB
var err error 
var sugar *zap.SugaredLogger

func InitDB() {
	sugar = utils.Logger.Sugar()
	defer sugar.Sync()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser,
		utils.DBPassWord,
		utils.DBHost,
		utils.DBPort,
		utils.DBName,
	)

	db, err = sqlx.Open(utils.DB, dsn)
	if err != nil {
		sugar.Errorf("参数格式有误:%w\n", err)
	}

	err = db.Ping()
	if err != nil {
		sugar.Errorf("数据库连接失败:%w\n", err)
	}

	db.SetMaxOpenConns(20) // 设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) // 设置连接池中最大的闲置连接数
}