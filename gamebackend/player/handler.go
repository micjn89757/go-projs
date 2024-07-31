package player

import (
	"fmt"
	"gamebackend/common"
	"gamebackend/network"
	"gamebackend/network/protocol/gen/player"

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

	if !common.CheckInNumberSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

// DelFriend 删除好友
func (p *Player) DelFriend(msg *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = common.DelEleInSlice(req.UId, p.FriendList)
}

// ResolveChatMsg 处理消息
func (p *Player) ResolveChatMsg(msg *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
	// TODO: 收到好友消息，然后转发给客户端
}
