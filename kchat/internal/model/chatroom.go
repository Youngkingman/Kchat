package model

import (
	"context"
	"encoding/json"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

type ChatRoom struct {
	RoomID int            `json:"room_id"`
	Name   string         `json:"name"`
	Users  map[string]int `json:"users"`
}

type tranChatRoom struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Users  string `json:"users"`
}

func AddChatRoom(ctx context.Context, name string, users map[string]int) error {
	storeUsers, err := json.Marshal(users)
	if err != nil {
		return err
	}
	s := `INSERT INTO #__chatroom (name,users) VALUES (?,?)`
	_, err = global.MySQL.Exec(dbutil.Prefix(s), name, storeUsers)
	return err
}

// 需要用到事务
// func AddUserToChatRoom(ctx context.Context, rid, uid int) error {

// }

func GetChatRoomByRoomId(ctx context.Context, rid int) (*ChatRoom, error) {
	tmp := &tranChatRoom{}
	s := "SELECT * FROM #__chatroom WHERE rid=?"
	err := global.MySQL.Get(tmp, dbutil.Prefix(s), rid)
	if err != nil {
		return nil, err
	}
	users := make(map[string]int)
	err = json.Unmarshal([]byte(tmp.Users), &users)
	if err != nil {
		return nil, err
	}
	chatRoom := &ChatRoom{
		RoomID: rid,
		Name:   tmp.Name,
		Users:  users,
	}
	return chatRoom, nil
}
