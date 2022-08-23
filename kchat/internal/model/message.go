package model

import "time"

const (
	MsgTypeNormal    = iota // 普通 用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeUserLeave        // 用户退出
	MsgTypeError            // 错误消息
	MsgTypeImage            // 图片数据, 此时content为文件服务器url
)

type Message struct {
	RoomID         int       `json:"room_id"`
	UID            int       `json:"chatter"`
	Name           string    `json:"name"`
	Avatar         string    `json:"avatar"`
	Type           int       `json:"type"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msg_time"`
	ClientSendTime time.Time `json:"client_send_time"`
	Ats            []string  `json:"ats"` // 消息 @ 了谁
}

// 暂时不提供聊天消息保存的功能，但是这么设计可以应付业务的扩展，虽然看着有点多余
