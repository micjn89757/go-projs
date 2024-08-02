package model

import "time"

type Player struct {
	Id         uint64 		`gorm:"column:id;primarykey"`
	Uid        uint64		`gorm:"column:uid"`
	NickName   string		`gorm:"column:nickname"`
	Sex        int			`gorm:"column:sex"`
	CreatedAt  time.Time	`gorm:"column:create_time;autoCreateTime:milli"`
	UpdatedAt  time.Time	`gorm:"column:update_time"`
}


func (Player) TableName() string { // 显示指定表名
	return "player" 		
}