package model

import (
	"time"
	"vote-gin/utils/msgcode"
)

// 投票内容
type Vote struct {
	Base
	Title 	string 		`db:"title" json:"title"`
	Type 	int 		`db:"type" json:"type"`
	Status	int			`db:"status" json:"status"`
	Time	time.Time	`db:"time" json:"time"`	
	UserID	int			`db:"user_id" json:"user_id"`
}


// GetVote 获取投票内容
func GetVote(id int) (Vote, int) {
	var err error 
	var vote Vote 

	sqlStr := "select id, title, type, status, time, user_id from vote where id = ?"
	err = db.Get(&vote, sqlStr, id)

	if err != nil {
		return vote, msgcode.ERROR_VOTE_NOT_EXIST
	}

	return vote, msgcode.SUCCESS
}


// GetVotes 获取投票列表
func GetUsers(status int, pageSize int, pageNum int) ([]*Vote,int, int) {
	var err error 
	var votes []*Vote 

	sqlStr := "select id, title, type, status, time, user_id from vote where status = ? Limit ? OFFSET ?"

	err = db.Select(&votes, sqlStr, status, pageSize, ((pageNum-1) * pageSize))
	if err != nil {
		return votes, msgcode.ERROR_STATUS_NOT_EXISIT, 0
	}

	return votes, msgcode.SUCCESS, len(votes)
}