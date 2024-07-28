package player

import "mmorpg/network/protocol/gen/messageId"

// HandlerRegister 注册玩家的各种handler
func (p *Player) HandlerRegister() {
	p.handlers[messageId.MessageId_CSAddFriend] = p.AddFriend
	p.handlers[messageId.MessageId_CSDelFriend] = p.DelFriend
	p.handlers[messageId.MessageId_CSSendChatMsg] = p.ResolveChatMsg
}