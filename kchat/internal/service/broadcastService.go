package service

import "github.com/Youngkingman/Kchat/kchat/internal/model"

type broadcast struct {
	// 所有 channel 统一管理，可以避免外部乱用

	enteringChannel chan *Chatter
	leavingChannel  chan *Chatter
	messageChannel  chan *model.Message

	// 判断该昵称用户是否可进入聊天室（重复与否）：true 能，false 不能
	checkUserChannel      chan string
	checkUserCanInChannel chan bool

	// 获取用户列表
	requestUsersChannel chan struct{}
	usersChannel        chan []*Chatter
}
