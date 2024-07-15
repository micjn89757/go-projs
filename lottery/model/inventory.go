package model

import (
	"fmt"
	"lottery/utils/errmsg"

	"gorm.io/gorm"
)

const (
	prefix = "gift_count_" // 设置redis的key统一前缀，方便按前缀遍历key
)

type Inventory struct {
	*gorm.Model
	Name 		string 	`gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Picture 	string 	`gorm:"column:picture"`
	Price 		int 	`gorm:"column:price"`
	Count 		int		`gorm:"column:count"`
}

func (Inventory) TableName() string {
	return "inventory"
}

// TODO: InitGiftInventory 从mysql读出所有奖品的初始库存，存入redis。如果同时有很多用户参与抽奖，不能发去Mysql里减库存，Mysql扛不住这么高的并发，Redis可以
func InitInventory() {

}

// GetAllInventoryCount 获取所有奖品的剩余库存量，返回的结果只包含id和count
func GetAllInventoryCount() []*Inventory {
	// redis key是prefix+id， value是count
	return nil
}

// GetAllInventoryV1 读取全部数据
func GetAllInventoryV1() ([]*Inventory, int, int64) {
	var inventoryList []*Inventory
	var total int64
	err := lotteryDB.Select([]string{"id", "name", "description", "picture", "price", "count"}).Find(&inventoryList).Error
	if err != nil {
		fmt.Println("read table failed")
		return inventoryList, errmsg.ERROR_GIFTS_NOT_EXIST, 0
	}

	lotteryDB.Model(&inventoryList).Count(&total)

	return inventoryList, errmsg.SUCCESS, total
}


// 千万级以上大表遍历方案，简单说每次读取一小份，将这一小份放到channel，另一个Goroutine从channel里读
func GetAllInventoryV2(ch chan <- Inventory) {}

// AddInventory
func AddInventory(InvId int) int {
	return 0
}


// DeleteInventory 
func DeleteInventory(InvId int) int {
	return 0
}