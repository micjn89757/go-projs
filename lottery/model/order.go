package model

import (
	"lottery/utils"

	"go.uber.org/zap"
)

type Order struct {
	ID          uint
	InventoryId uint
	UserId      uint
}

var (
	orderCh = make(chan Order, 10000) // 最高瞬时可以下10000单
	stopCh  = make(chan struct{}, 1)
)

func InitChannel() {
	go func() { // 等待接收关闭订单通道信号
		<-stopCh
		close(orderCh)
	}()
}

// PutOrder 将订单放入channel
func PutOrder(userId, inventoryID uint) {
	order := Order{UserId: userId, InventoryId: inventoryID}
	orderCh <- order
}

// TakeOrder 从channel中取出订单，写入Mysql
func TakeOrder() {
	for {
		order, ok := <- orderCh

		if !ok {
			utils.Logger.Info("order channel is closed")
			break
		}

		CreateOrder(order.UserId, order.InventoryId) // 写入mysql
	}
}


// CreateOrder 创建订单
func CreateOrder(userId, inventoryId uint) int {
	order := Order{InventoryId: inventoryId, UserId: userId}
	if err := lotteryDB.Create(&order).Error; err != nil {
		utils.Logger.Error("create order failed", zap.String("errmsg", err.Error()))
		return 0
	} else {
		utils.Logger.Error("create order", zap.Uint("id", order.ID))
		return int(order.ID)
	}
}


// ClearOrders 清除全部订单
func ClearOrders() error {
	err := lotteryDB.Where("id > 0").Delete(Order{}).Error

	if err != nil {
		utils.Logger.Error("delete order failed", zap.String("errmsg", err.Error()))
	}

	return err
}


// CloseChannel 关闭orderCh
func CloseChannel() {
	select {
	case stopCh <- struct{}{}: // 不让函数阻塞再本行代码，外套一个select
	default:
	}
}