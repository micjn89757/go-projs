package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// 会话管理
// 多个客户端可以建立多个会话
type Session struct {
	conn net.Conn
	packer *NormalPacker
	writeCh	chan *Message
}

func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn, packer: NewNormalPack(binary.BigEndian), writeCh: make(chan *Message, 1)}
}


func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 读取客户端消息
func (s *Session) Read() {
	// 设置读超时时间
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}

	for {
		msg, err := s.packer.Unpack(s.conn)
		if err != nil {
			fmt.Println(err)
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
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case msg := <- s.writeCh:
			s.send(msg)
		}
	}
}

func (s *Session) send(message *Message) {
	bytes, err := s.packer.Pack(message)

	if err != nil {
		return
	}

	_, err = s.conn.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}