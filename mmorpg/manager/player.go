package manager

import "mmorpg/player"

// 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player
	pCh	chan player.Player	
}

// Add 添加在线玩家
func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.UId] = p
	go p.Run()
}

// Run 运行玩家管理服务
func (pm *PlayerMgr) Run() {
	for {
		select {
		case player := <- pm.pCh:
			pm.Add(player)
		}
	}
}