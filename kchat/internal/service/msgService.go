package service

import (
	"time"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
)

var MsgSrv messageService = &Message{}

type Message struct {
	Msg *model.Message
}

type messageService interface {
	NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message
	NewWelcomeMessage(chatter *model.Chatter) *model.Message
	NewChatterEnterMessage(chatter *model.Chatter) *model.Message
	NewChatterLeaveMessage(chatter *model.Chatter) *model.Message
	NewErrorMessage(content string) *model.Message
	NewImageMessage(chatter *model.Chatter, content, clientTime string) *model.Message
}

func (ms *Message) NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message {
	message := &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Name:    chatter.Name,
		Avatar:  chatter.ImageURL,
		Type:    model.MsgTypeNormal,
		Content: model.ContetnObj{Text: content},
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		// 和前端一样整个解析器format下
		message.ClientSendTime = time.Now().Format("2006-01-02 15:04:05")
	}
	return message
}

func (ms *Message) NewImageMessage(chatter *model.Chatter, imgUrl, clientTime string) *model.Message {
	message := &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeImage,
		Content: model.ContetnObj{Text: imgUrl},
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		// 需要format成字符串
		message.ClientSendTime = time.Now().Format("2006-01-02 15:04:05")
	}
	return message
}

func (ms *Message) NewWelcomeMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Name:    "System Information",
		Avatar:  global.AppSetting.SystemAvatar,
		Type:    model.MsgTypeWelcome,
		Content: model.ContetnObj{Text: chatter.Name + " Welcome to Kchat"},
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewChatterEnterMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeUserEnter,
		Content: model.ContetnObj{Text: chatter.Name + " enters Kchat"},
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewChatterLeaveMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Name:    "System Information",
		Avatar:  global.AppSetting.SystemAvatar,
		Type:    model.MsgTypeUserLeave,
		Content: model.ContetnObj{Text: chatter.Name + " leaves Kchat"},
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewErrorMessage(content string) *model.Message {
	return &model.Message{
		UID:     System.UID,
		RoomID:  System.RoomID,
		Name:    "System Information",
		Avatar:  global.AppSetting.SystemAvatar,
		Type:    model.MsgTypeUserLeave,
		Content: model.ContetnObj{Text: content},
		MsgTime: time.Now(),
	}
}
