package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener 	net.Listener
	addr  		string 
	network 	string
}


func NewServer(address, network string) *Server {
	return &Server{
		listener: nil,
		addr: address,
		network: network,
	}
}


func (s *Server) Run() {
	// tcp6指的是ipv6网络中的TCP连接，核心功能保持不变
	// tcp4指的是ipv4网络中的TCP连接
	// 这里返回一个tcp6的地址对象，包含地址和端口等信息
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.addr)
	if err != nil {
		fmt.Println(err)
		return 
	}

	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	s.listener = tcpListener
	for {	// 持续等待客户端连接
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		go func() {	// 处理会话
			session := NewSession(conn)
			session.Run()
		}()
	}
		
}

