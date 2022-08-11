package model

import "time"

const (
	MsgTypeNormal    = iota // 普通 用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeUserLeave        // 用户退出
	MsgTypeError            // 错误消息
)

type Message struct {
	Chatter        *Chatter  `json:"chatter"`
	Type           int       `json:"type"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msg_time"`
	ClientSendTime time.Time `json:"client_send_time"`
	// 消息 @ 了谁
	Ats []string `json:"ats"`
}
