package model

import "time"

// 投票内容
type Vote struct {
	Base
	Title 	string 		`db:"title" json:"title"`
	Type 	int 		`db:"type" json:"type"`
	Status	int			`db:"status" json:"status"`
	Time	time.Time	`db:"time" json:"time"`	
	UserID	int			`db:"user_id" json:"user_id"`
}