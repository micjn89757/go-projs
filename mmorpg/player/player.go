package player

import (
	"mmorpg/network"
	"mmorpg/network/protocol/gen/messageId"
)

type Player struct {
	UId        		uint64
	FriendList 		[]uint64 	// 好友列表
	HandlerParamCh	chan *network.Message
	handlers		map[messageId.MessageId]Handler
	session 		*network.Session
}

// 创建玩家
func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 0),
		handlers: make(map[messageId.MessageId]Handler),
	}

	p.HandlerRegister()
	return p 
}

// Run 角色运行
func (p * Player) Run() {
	for {
		select {
		// 从HandlerParamCh获取玩家需要做的处理
		case handlerParam := <- p.HandlerParamCh:
			// key对应handler的key，data对应handler要传入的参数
			if fn, ok := p.handlers[messageId.MessageId(handlerParam.Id)]; ok {
				fn(handlerParam)
			}
		}
	}
}

