package model

type VoteOpt struct {
	Base
	VoteID int    `db:"vote_id" json:"vote_id"`
	Name   string `db:"name" json:"name"`
	Count  int    `db:"count" json:"count"`
}

// CreateVoteOpt 创建选项
func CreateVoteOpt(vo VoteOpt) int {
	return 1
}
