package manager

import "gamebackend/player"

// 维护在线玩家
type PlayerMgr struct {
	players map[uint64]*player.Player
	addCh   chan *player.Player
}

// Add 添加在线玩家
func (pm *PlayerMgr) Add(p *player.Player) {
	pm.players[p.UId] = p
	go p.Run() // 角色启动
}

// Del 删除在线玩家
func (pm *PlayerMgr) Del(p *player.Player) {
	delete(pm.players, p.UId)
}

// Run 运行玩家管理服务
func (pm *PlayerMgr) Run() {
	for {
		select {
		case player := <-pm.addCh:
			pm.Add(player)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
