package model

import (
	"database/sql/driver"
	"strings"
	"vote-gin/utils/msgcode"

	"github.com/jmoiron/sqlx"
)

type VoteOpt struct {
	Base
	VoteID int    `db:"vote_id" json:"vote_id"`
	Name   string `db:"name" json:"name"`
	Count  int    `db:"count" json:"count"`
}

// 以便使用sqlx.In批量插入
func (vo VoteOpt) Value() (driver.Value, error) {
	return []any{vo.VoteID, vo.Name, vo.Count}, nil
}

// CreateVoteOpt 批量创建投票选项
func CreateVoteOpts(vos []any) int {
	var err error

	valueString := make([]string, 0, len(vos))

	sqlStr := "insert into vote_opt(vote_id, name, count, created_time, updated_time) values"
	for i := 0; i < len(vos); i++ {
		valueString = append(valueString, "(?)")
	}

	sqlStr = sqlStr + strings.Join(valueString, ",")

	query, args, err := sqlx.In(sqlStr, vos...)

	if err != nil {
		sugar.Errorf("concat insert string failed:%s", err.Error())
		return msgcode.ERROR
	}

	_, err = db.Exec(query, args...)

	if err != nil {
		sugar.Errorf("batch insert vote_opt failed:%s", err.Error())
		return msgcode.ERROR
	}

	sugar.Info("Create vote opts success!")

	return msgcode.SUCCESS
}

// CreateVoteOpt 创建投票选项
func CreateVoteOpt(vo VoteOpt) (int64, int) {
	sqlStr := "insert into vote_opt(name, count, vote_id, created_time, updated_time) values(?, ?, ?, ?, ?)"
	ret, err := db.NamedExec(sqlStr, vo)

	if err != nil {
		sugar.Errorf("insert vote opt failed: %s", err.Error())
		return 0, msgcode.ERROR
	}

	insertId, err := ret.LastInsertId()
	if err != nil {
		sugar.Errorf("get insertID failed: %s", err.Error())
		return 0, msgcode.ERROR
	}

	sugar.Infof("insert vote opt success, insertID is: %d", insertId)

	return insertId, msgcode.SUCCESS
}

// GetVoteOpts  根据投票id获取选项内容
func GetVoteOpts(id uint) ([]VoteOpt, error) {
	var err error
	var vos []VoteOpt

	sqlStr := "select id, name, count, vote_id, created_time, updated_time from vote_opt where vote_id = ?"

	err = db.Select(&vos, sqlStr, id)

	if err != nil {
		sugar.Errorf("get vote_opts error:%s", err.Error())
		return vos, err
	}

	return vos, err
}

// EditVoteOpt 编辑投票选项
func EditVoteOpt(id uint, data VoteOpt) (VoteOpt, error) {
	var vo VoteOpt
	var maps = make(map[string]any, 0)

	maps["count"] = data.Count

	sqlStr := "update vote_opt set count = ? where id = ?"

}
