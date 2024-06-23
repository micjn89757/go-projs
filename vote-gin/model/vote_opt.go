package model

type VoteOpt struct {
	Base
	VoteID int    `db:"vote_id" json:"vote_id"`
	Name   string `db:"name" json:"name"`
	Count  int    `db:"count" json:"count"`
}

// TODO CreateVoteOpt 创建选项
func CreateVoteOpt(vo VoteOpt) int {
	return 1
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
