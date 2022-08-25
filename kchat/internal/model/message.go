package model

import "time"

const (
	MsgTypeNormal    = iota // 普通 用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeUserLeave        // 用户退出
	MsgTypeError            // 错误消息
	MsgTypeImage            // 图片数据, 此时content为文件服务器url的html标签
)

// 用别人模板有些字段不太一样，以后再来优化
type Message struct {
	RoomID         int        `json:"room_id"`
	UID            int        `json:"uid"`
	Name           string     `json:"name"`
	Avatar         string     `json:"img"`
	Type           int        `json:"type"`
	Content        ContetnObj `json:"text"`
	MsgTime        time.Time  `json:"msg_time"`
	ClientSendTime string     `json:"date"`
	// Ats            []string  `json:"ats"` // 消息 @ 了谁
}

// 作为text字段的内嵌功能，被前端拿捏了，我恨啊
type ContetnObj struct {
	Text string `json:"text"`
}

// 暂时不提供聊天消息保存的功能，但是这么设计可以应付业务的扩展，虽然看着有点多余
