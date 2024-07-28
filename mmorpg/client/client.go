package main

import (
	"fmt"
	"mmorpg/network"
	"mmorpg/network/protocol/gen/messageId"
)

// 模拟客户端
type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewClientConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.MsgCh = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:	// 接收控制台输入
				fmt.Printf("cmd:%s,param:%#v <<<\n", input.Command, input.Param)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}

			}
		}
	}()
	go c.console.Run() 	// 接收控制台输入
	go c.cli.Run()		// 客户端运行
}

func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
	}
}
