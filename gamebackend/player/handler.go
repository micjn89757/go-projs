package player

import (
	"fmt"
	"gamebackend/common"
	"gamebackend/network"
	"gamebackend/network/protocol/gen/player"
	"gamebackend/network/protocol/gen/messageId"

	"google.golang.org/protobuf/proto"
)

type Handler func(packet *network.Message)

// AddFriend 添加好友
func (p *Player) AddFriend(msg *network.Message) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(msg.Data, req) // 解析proto消息
	if err != nil {
		return
	}

	// 查看是否已经添加过
	if !common.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}

	bytes, err := proto.Marshal(&player.SCAddFriend{})

	if err != nil {
		return 
	}

	rsp := &network.Message{
		Id: uint64(messageId.MessageId_SCAddFriend),
		Data: bytes,
	}

	p.session.SendMsg(rsp)
}

// DelFriend 删除好友
func (p *Player) DelFriend(msg *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = common.DelEleInSlice(req.UId, p.FriendList)

	bytes, err := proto.Marshal(&player.SCDelFriend{})

	if err != nil {
		return
	}

	rsp := &network.Message{
		Id: uint64(messageId.MessageId_SCDelFriend),
		Data: bytes,
	}

	p.session.SendMsg(rsp)
}

// ResolveChatMsg 处理消息
func (p *Player) ResolveChatMsg(msg *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)

	bytes, err := proto.Marshal(&player.SCSendChatMsg{})

	if err != nil {
		return
	}


	rsp := &network.Message{
		Id: uint64(messageId.MessageId_SCSendChatMsg),
		Data: bytes,
	}

	p.session.SendMsg(rsp)
}
