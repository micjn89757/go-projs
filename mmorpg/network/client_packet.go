package network

import "net"

// 客户端数据包格式
type ClientPacket struct {
	Msg 	*Message 
	Conn 	net.Conn
}