package v1

import (
	"net/http"
	"strconv"
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
			"msg":    msgcode.GetErrMsg(msgcode.ERROR),
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

// AddVoteOpt 添加投票选项
func AddVoteOpt(ctx *gin.Context) {
	var err error
	var code int
	var voteOpt model.VoteOpt

	err = ctx.ShouldBindJSON(voteOpt)
	if err != nil {
		sugar.Infof("add vote opt failed: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": msgcode.ERROR,
			"msg":    msgcode.GetErrMsg(msgcode.ERROR),
		})
	}

	_, code = model.CreateVoteOpt(voteOpt)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msgcode.GetErrMsg(code),
		"data":   voteOpt,
	})

}

// GetVoteInfo 获取投票
func GetVoteInfo(ctx *gin.Context) {
	var code int
	var vote model.Vote
	var voteOpts []model.VoteOpt
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		sugar.Errorf("get vote id failed:%s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": msgcode.ERROR,
			"msg":    msgcode.GetErrMsg(msgcode.ERROR),
		})
		ctx.Abort()
	}

	vote, voteOpts, code = model.GetVoteInfo(uint(id))

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msgcode.GetErrMsg(code),
		"data": gin.H{
			"Vote":     vote,
			"VoteOpts": voteOpts,
		},
	})
}
