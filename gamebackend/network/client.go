package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Addr   		string
	packer 		IPacker
	MsgCh		chan *Message
	OnMessage	func(message *ClientPacket)
}

func NewClient(address string) *Client {
	return &Client{
		Addr: address,
		packer: &NormalPacker{
			ByteOrder: binary.BigEndian,
		},
		MsgCh: make(chan *Message, 1),
	}
}


func (c *Client) Run() {
	conn, err := net.Dial("tcp6", c.Addr)
	if err != nil {
		fmt.Println(err)
		return 
	}

	go c.Write(conn)
	go c.Read(conn)
}


func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)

	for {
		select {
		case <- tick.C:
			c.MsgCh <- &Message{
				Id: 111,
				Data: []byte("hello world"),
			}
		case msg := <- c.MsgCh:
			c.send(conn, msg)
		}

	}
}


func (c *Client) send(conn net.Conn, msg *Message) {
	bytes, err := c.packer.Pack(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}


func (c *Client) Read(conn net.Conn) {
	for {
		msg, err := c.packer.Unpack(conn)
		if _, ok := err.(net.Error); ok { // 如果是因为网络原因
			fmt.Println(err)
			continue
		}

		c.OnMessage(&ClientPacket{
			Msg: msg,
			Conn: conn,
		})

		fmt.Println("resp message: ", string(msg.Data))
	}
}