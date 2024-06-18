package model


type VoteOptUser struct {
	Base
	UserID	int `db:"user_id" json:"user_id"`
	VoteID 	int `db:"vote_id" json:"vote_id"`
	VoteOptID int `db:"vote_opt_id" json"vote_opt_id"`
}


