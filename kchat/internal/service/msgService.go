package service

import (
	"time"

	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/spf13/cast"
)

var MsgSrv messageService = &Message{}

type Message struct {
	Msg *model.Message
}

type messageService interface {
	NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message
	NewWelcomeMessage(chatter *model.Chatter) *model.Message
	NewUserEnterMessage(chatter *model.Chatter) *model.Message
	NewUserLeaveMessage(chatter *model.Chatter) *model.Message
	NewErrorMessage(content string) *model.Message
	NewImageMessage(chatter *model.Chatter, content, clientTime string) *model.Message
}

func (ms *Message) NewMessage(chatter *model.Chatter, content, clientTime string) *model.Message {
	message := &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

func (ms *Message) NewImageMessage(chatter *model.Chatter, imgUrl, clientTime string) *model.Message {
	message := &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeImage,
		Content: imgUrl,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

func (ms *Message) NewWelcomeMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeWelcome,
		Content: chatter.Name + " Welcome to Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewUserEnterMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeUserEnter,
		Content: chatter.Name + " enters Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewUserLeaveMessage(chatter *model.Chatter) *model.Message {
	return &model.Message{
		UID:     chatter.UID,
		RoomID:  chatter.RoomID,
		Type:    model.MsgTypeUserLeave,
		Content: chatter.Name + " leaves Kchat",
		MsgTime: time.Now(),
	}
}

func (ms *Message) NewErrorMessage(content string) *model.Message {
	return &model.Message{
		UID:     System.UID,
		RoomID:  System.RoomID,
		Type:    model.MsgTypeUserLeave,
		Content: content,
		MsgTime: time.Now(),
	}
}
