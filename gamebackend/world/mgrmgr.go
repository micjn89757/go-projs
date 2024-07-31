package world

import (
	"gamebackend/common"
	"gamebackend/manager"
	"gamebackend/network"
	"gamebackend/network/protocol/gen/messageId"
	"os"
	"syscall"

	"go.uber.org/zap"
)

// 世界管理服务
type MgrMgr struct {
	Pm              *manager.PlayerMgr // 玩家管理服务
	Server          *network.Server    
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: &manager.PlayerMgr{}}
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	m.Handlers = make(map[messageId.MessageId]func(message *network.SessionPacket))

	return m
}

var MM *MgrMgr

func (mm *MgrMgr) Run() {
	mm.HandlerRegister()	// 注册handler
	go mm.Server.Run() // 世界管理服务
	go mm.Pm.Run()     // 玩家管理服务
}

func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
		return
	}

	if p := mm.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (mm *MgrMgr) OnSystemSignal(signal os.Signal) bool {
	common.Logger.Debug("[MgrMgr] 收到信号", zap.Any("信号:", signal))
	tag := true
	switch signal {
	case syscall.SIGHUP: // 挂起
		//todo
	case syscall.SIGPIPE:
	default:
		common.Logger.Debug("[MgrMgr] 收到信号准备退出...")
		tag = false
	}
	return tag
}
