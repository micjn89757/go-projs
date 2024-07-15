package v1

import (
	"lottery/model"
	"lottery/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllGifts 初始化所有奖品信息，用于初始化轮盘
func GetAllInvent(ctx *gin.Context) {
	var code int
	inventories, code, total := model.GetAllInventoryV1()
	
	// 去除敏感信息
	for _, inv := range inventories {
		inv.Count = 0
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": inventories,
		"total": total,
		"msg": errmsg.GetErrMsg(code),
	})
	
}

// Lottery 抽奖，并发执行
func Lottery(ctx *gin.Context) {
	for try := 0; try < 10; try++ { // 最多重试10次
		// 1. 首先获取所有奖品剩余库存
		// 2. 确保从redis中查到的奖品库存量大于0，因为抽奖算法Lottery不支持抽中概率为0的奖品（过滤）

		// 3. 如果所有奖品都被抽空了，则返回"0"，表示奖品抽完
		// 4. 抽中第Index个奖品则返回id
		// 5. 减redis的库存(注意并发问题，不过redis的decr是并发安全的)，有可能返回负数（并发减库存），但是不会返回抽奖失败，而是尝试10次，有一次成功就返回，如果10次都抽不到，就返回失败，1

		// 抽奖算法是个概率问题，但是10次都不成功是小概率时间，但是不能为了特别小的概率时间做太多处理
	}
}