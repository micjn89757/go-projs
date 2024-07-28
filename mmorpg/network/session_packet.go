package network

// Session数据包格式
type SessionPacket struct {
	Msg  *Message
	Sess *Session
}
