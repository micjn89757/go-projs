package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// 多个客户端可以建立多个会话
type Session struct {
	UId		int64 	// 对session进行标识
	Conn 	net.Conn
	packer 	IPacker
	writeCh	chan *Message
	IsClose	bool
}

func NewSession(conn net.Conn) *Session {
	return &Session{Conn: conn, packer: &NormalPacker{binary.BigEndian}, writeCh: make(chan *Message, 1)}
}


func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 读取客户端消息
func (s *Session) Read() {
	for {	
		// 设置读超时时间
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg, err := s.packer.Unpack(s.Conn)
		if _, ok := err.(net.Error); ok {
			continue
		}

		// TODO: 序列化成需要的格式
		fmt.Println("server receive message: ", string(msg.Data))

		// TODO:
		s.writeCh <- &Message{ // 回复
			Id: 999,
			Data: []byte("hi calvin"),
		}
	}
	
}


// Write 向客户端返回消息
func (s *Session) Write() {
	for {
		select {
		case resp := <- s.writeCh:
			s.send(resp)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return 
	}
	bytes, err := s.packer.Pack(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = s.Conn.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}