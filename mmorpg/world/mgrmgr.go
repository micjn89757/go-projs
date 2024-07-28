package world

import (
	"mmorpg/manager"
	"mmorpg/network"
	"mmorpg/network/protocol/gen/messageId"
)

// 统一管理服务
type MgrMgr struct {
	Pm 				*manager.PlayerMgr	// 玩家管理
	Server          *network.Server		// 统一管理服务器
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: &manager.PlayerMgr{}}
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	return m
}


var MM *MgrMgr

func (mm *MgrMgr) Run() {
	go mm.Server.Run()
	go mm.Pm.Run()
}

func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
	}

	if p := mm.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}