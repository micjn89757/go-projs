package v1

import (
	"net/http"
	"vote-gin/model"
	"vote-gin/utils/msgcode"

	"github.com/gin-gonic/gin"
)

// AddVote 添加投票
func AddVote(ctx *gin.Context) {
	var err error
	var vote model.Vote
	err = ctx.ShouldBindJSON(&vote)

	if err != nil {
		sugar.Infof("get params failed: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": msgcode.ERROR,
			"msg":    err.Error(),
		})
		ctx.Abort()
		return
	}

	code := model.CreateVote(vote)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msgcode.GetErrMsg(code),
		"data":   vote,
	})
}
