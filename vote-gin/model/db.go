package model

import (
	"fmt"
	"time"
	"vote-gin/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB
var err error
var sugar *zap.SugaredLogger

// 基础结构体
type Base struct {
	ID          uint      `db:"id" json:"id"`
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"update_time"`
	DeletedTime time.Time `db:"deleted_time" json:"deleted_time"`
}

// TODO 更改成单例
func InitDB() {
	sugar = utils.Logger.Sugar()
	defer sugar.Sync()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser,
		utils.DBPassWord,
		utils.DBHost,
		utils.DBPort,
		utils.DBName,
	)

	db, err = sqlx.Open(utils.DB, dsn)
	if err != nil {
		sugar.Errorf("参数格式有误:%w\n", err)
		panic("datasourcename error")
	}

	err = db.Ping()
	if err != nil {
		sugar.Errorf("数据库连接失败:%w\n", err)
		panic("database connect failed")
	}

	db.SetMaxOpenConns(20) // 设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) // 设置连接池中最大的闲置连接数
}

func Close() {
	db.Close()
}
