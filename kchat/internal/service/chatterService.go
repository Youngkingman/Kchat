package service

import (
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"golang.org/x/net/websocket"
)

var System = &model.Chatter{}

type ChatterService struct {
	model.Chatter
	MessageChannel chan *model.Message
	conn           *websocket.Conn
}
