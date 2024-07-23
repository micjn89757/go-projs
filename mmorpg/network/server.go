package network

import (
	"net"
)

type Server struct {
	Listener 	net.Listener
}


func NewServer(address string) *Server {
	// tcp6指的是ipv6网络中的TCP连接，核心功能保持不变
	// tcp4指的是ipv4网络中的TCP连接
	// 这里返回一个tcp6的地址对象，包含地址和端口等信息
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", address)
	if err != nil {
		panic(err)
	}

	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)

	if err != nil {
		panic(err)
	}

	s := &Server{}
	s.Listener = tcpListener
	return s
}


func (s *Server) Run() {

	for {	// 持续等待客户端连接
		conn, err := s.Listener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				continue
			}
		}

		go func() {	// 处理会话
			session := NewSession(conn)
			SessionMgrInstance.AddSession(session)
			session.Run()
			SessionMgrInstance.DelSession(session.UId)
		}()
	}
		
}

