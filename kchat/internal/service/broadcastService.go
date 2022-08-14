package service

import (
	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/gin-gonic/gin"
)

type broadcaster struct {
	// 所有 channel 统一管理，可以避免外部乱用
	RoomID          int
	chatters        map[string]*Chatter //在线用户信息
	enteringChannel chan *Chatter
	leavingChannel  chan *Chatter
	messageChannel  chan *model.Message

	// 判断该昵称用户是否可进入聊天室（重复与否）：true 能，false 不能
	checkChatterChannel      chan string
	checkChatterCanInChannel chan bool

	// 获取用户列表
	requestChattersChannel chan struct{}
	chattersChannel        chan []*Chatter
}

func NewBroadCast(rid int) *broadcaster {
	return &broadcaster{
		chatters: make(map[string]*Chatter),

		enteringChannel: make(chan *Chatter),
		leavingChannel:  make(chan *Chatter),
		messageChannel:  make(chan *model.Message, global.ChatRoomSetting.MessageQueueLen),

		checkChatterChannel:      make(chan string),
		checkChatterCanInChannel: make(chan bool),

		requestChattersChannel: make(chan struct{}),
		chattersChannel:        make(chan []*Chatter),
	}
}

// Start 启动广播器
// 需要在一个新 goroutine 中运行，因为它不会返回
func (b *broadcaster) Start() {
	for {
		select {
		case chatter := <-b.enteringChannel:
			// 新用户进入
			b.chatters[chatter.Chatter.Name] = chatter

			//OfflineProcessor.Send(user)
		case chatter := <-b.leavingChannel:
			// 用户离开
			delete(b.chatters, chatter.Chatter.Name)
			// 避免 goroutine 泄露
			chatter.CloseMessageChannel()
		case msg := <-b.messageChannel:
			// 给所有在线用户发送消息
			for _, chatter := range b.chatters {
				if chatter.Chatter.UID == msg.UID {
					continue
				}
				chatter.MessageChannel <- msg
			}
			//OfflineProcessor.Save(msg)
		case name := <-b.checkChatterChannel:
			if _, ok := b.chatters[name]; ok {
				b.checkChatterCanInChannel <- false
			} else {
				b.checkChatterCanInChannel <- true
			}
		case <-b.requestChattersChannel:
			chatterList := make([]*Chatter, 0, len(b.chatters))
			for _, chatter := range b.chatters {
				chatterList = append(chatterList, chatter)
			}

			b.chattersChannel <- chatterList
		}
	}
}

func (b *broadcaster) ChatterEntering(c *Chatter) {
	b.enteringChannel <- c
}

func (b *broadcaster) ChatterLeaving(c *Chatter) {
	b.leavingChannel <- c
}

func (b *broadcaster) Broadcast(ctx *gin.Context, msg *model.Message) {
	if len(b.messageChannel) >= global.ChatRoomSetting.MessageQueueLen {
		global.Logger.Debug(ctx, "broadcast queue overfull")
	}
	b.messageChannel <- msg
}

func (b *broadcaster) CanEnterRoom(nickname string) bool {
	b.checkChatterChannel <- nickname

	return <-b.checkChatterCanInChannel
}

func (b *broadcaster) GetChatterList() []*Chatter {
	b.requestChattersChannel <- struct{}{}
	return <-b.chattersChannel
}
