package world

import "gamebackend/network/protocol/gen/messageId"

// 注册handler
func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_CSLogin] = mm.UserLogin
	mm.Handlers[messageId.MessageId_CSCreatePlayer] = mm.CreatePlayer
}