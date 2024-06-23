package model

import (
	"time"
	"vote-gin/utils/msgcode"
)

// 投票内容
type Vote struct {
	Base
	Title  string    `db:"title" json:"title"`
	Type   int       `db:"type" json:"type"`
	Status int       `db:"status" json:"status"`
	Time   time.Time `db:"time" json:"time"`
	UserID int       `db:"user_id" json:"user_id"`
}

// CreateVote 创建投票 TODO 还要创建投票选项
func CreateVote(v Vote) int {
	sqlStr := "insert into vote(title, type, status, time, UserID, created_time, updated_time) values(?, ?, ?, ?, ?, ?, ?)"

	ret, err := db.NamedExec(sqlStr, v)

	if err != nil {
		sugar.Errorf("insert vote error: %s", err.Error())
		return msgcode.ERROR
	}

	id, err := ret.LastInsertId() // 获取新插入数据id

	if err != nil {
		sugar.Errorf("get insert id failed, %s", err.Error())
		return msgcode.ERROR
	}

	sugar.Infof("insert vote success, the id is %d", id)

	return msgcode.SUCCESS
}

// GetVote 获取投票内容
func GetVote(id int) (Vote, []VoteOpt, int) {
	var err error
	var vote Vote
	var voteOpts []VoteOpt

	// 获取投票
	sqlStr := "select id, title, type, status, time, user_id from vote where id = ?"
	err = db.Get(&vote, sqlStr, id)

	if err != nil {
		return vote, voteOpts, msgcode.ERROR_VOTE_NOT_EXIST
	}

	// 获取投票选项
	sqlStr = "select id, vote_id, name, count from vote_opt where vote_id = ?"
	err = db.Select(&voteOpts, sqlStr, vote.ID)

	return vote, voteOpts, msgcode.SUCCESS
}

// GetVotes 获取投票列表
func GetUsers(status int, pageSize int, pageNum int) ([]Vote, int, int) {
	var err error
	var votes []Vote

	sqlStr := "select id, title, type, status, time, user_id from vote where status = ? Limit ? OFFSET ?"

	err = db.Select(&votes, sqlStr, status, pageSize, ((pageNum - 1) * pageSize))
	if err != nil {
		return votes, msgcode.ERROR_STATUS_NOT_EXISIT, 0
	}

	return votes, msgcode.SUCCESS, len(votes)
}
