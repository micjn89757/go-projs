/*
分发到各个不同的管理模块处理并返回结果
*/
package world

import (
	"fmt"
	"gamebackend/network"
	"gamebackend/network/protocol/gen/player"
	logicPlayer "gamebackend/player"
	"time"

	"google.golang.org/protobuf/proto"
)

// CreatePlayer 创建玩家
func (mm *MgrMgr) CreatePlayer(message *network.SessionPacket) {
	msg := &player.CSCreateUser{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("[MgrMgr.CreatePlayer]", msg)
	mm.SendMsg(message.Msg.Id, &player.SCCreateUser{}, message.Sess)

}

// UserLogin 用户登录
func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	newPlayer.UId = uint64(time.Now().Unix())
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	mm.Pm.Add(newPlayer)
	newPlayer.Run()
}

// SendMsg 发送消息
// 参数： message必须是一个protobuf定义的消息类型
func (mm *MgrMgr) SendMsg(id uint64, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &network.Message{
		Id:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}
