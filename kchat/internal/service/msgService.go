package service

import (
	"time"

	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/spf13/cast"
)

var MsgSrv messageService = &MessageService{}

type MessageService struct{}

type messageService interface {
	NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message
	NewWelcomeMessage(chatter *model.Chatter) *model.Message
	NewUserEnterMessage(chatter *model.Chatter) *model.Message
	NewUserLeaveMessage(chatter *model.Chatter) *model.Message
	NewErrorMessage(content string) *model.Message
}

// 之后改成房间信息
func (ms *MessageService) NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message {
	message := &model.Message{
		Chatter: chatter,
		Type:    model.MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

func (ms *MessageService) NewWelcomeMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		Chatter: chatter,
		Type:    model.MsgTypeWelcome,
		Content: chatter.Name + " Welcome to Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *MessageService) NewUserEnterMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		Chatter: chatter,
		Type:    model.MsgTypeUserEnter,
		Content: chatter.Name + " enters Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *MessageService) NewUserLeaveMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		Chatter: chatter,
		Type:    model.MsgTypeUserLeave,
		Content: chatter.Name + " leaves Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *MessageService) NewErrorMessage(content string) *model.Message {
	return &model.Message{
		Chatter: System,
		Type:    model.MsgTypeUserLeave,
		Content: content,
		MsgTime: time.Now(),
	}
}
