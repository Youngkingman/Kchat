package service

import (
	"errors"
	"io"
	"regexp"

	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var System = &model.Chatter{
	RoomID: 0,
	UID:    0,
}

type Chatter struct {
	Chatter        *model.Chatter
	MessageChannel chan *model.Message
	conn           *websocket.Conn
}

func NewChatter(user *model.User, conn *websocket.Conn, rid int) *Chatter {
	return &Chatter{
		Chatter:        model.NewChatterFromUser(user, rid),
		MessageChannel: make(chan *model.Message, 32),

		conn: conn,
	}
}

func (c *Chatter) SendMessage(ctx *gin.Context) {
	for msg := range c.MessageChannel {
		wsjson.Write(ctx.Request.Context(), c.conn, msg)
	}
}

func (c *Chatter) CloseMessageChannel() {
	close(c.MessageChannel)
}

func (c *Chatter) ReceiveMessage(ctx *gin.Context, bc *broadcaster) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		err = wsjson.Read(ctx.Request.Context(), c.conn, &receiveMsg)
		if err != nil {
			// 判定连接是否关闭了，正常关闭，不认为是错误
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			} else if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		// 内容发送到聊天室
		sendMsg := MsgSrv.NewMessage(c.Chatter, receiveMsg["content"], receiveMsg["send_time"])
		//sendMsg.Content = MsgSrv.FilterSensitive(sendMsg.Content)

		// 解析 content，看看 @ 谁了
		reg := regexp.MustCompile(`@[^\s@]{2,20}`)
		sendMsg.Ats = reg.FindAllString(sendMsg.Content, -1)

		bc.Broadcast(ctx, sendMsg)
	}
}
