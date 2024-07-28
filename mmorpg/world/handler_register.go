package world


// 注册handler
func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[1] = mm.UserLogin
}