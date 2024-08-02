package model

import "time"

type FriendShip struct {
	PlayerId 	uint64 		`gorm:"column:player_id;uniqueIndex"`
	FriendId	uint64		`gorm:"column:friend_id;uniqueIndex"`
	CreatedAt	time.Time	`gorm:"column:create_time;autoCreateTime:milli"`
	UpdatedAt	time.Time	`gorm:"column:update_time"`
}

func (FriendShip) TableName() string {
	
}