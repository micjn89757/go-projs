package world

import "mmorpg/manager"

// 统一管理服务
type MgrMgr struct {
	Pm manager.PlayerMgr
}

func NewMgrMgr() *MgrMgr {
	return &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
}


var MM *MgrMgr

func (mm *MgrMgr) name() {
	
}