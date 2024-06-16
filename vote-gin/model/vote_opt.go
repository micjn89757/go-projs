package model


type VoteOpt struct{
	Base
	VoteID	int `db:"vote_id" json:"vote_id"`
	Count	int `db:"count" json:"count"`
}