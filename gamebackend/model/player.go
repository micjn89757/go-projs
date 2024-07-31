package model


type Player struct {
	Uid			uint64
	NickName 	string 
	Sex 		int 
	FriendList	[]uint64
}