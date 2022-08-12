package model

import (
	"context"
	"encoding/json"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

type ChatRoom struct {
	RoomID int          `json:"room_id"`
	Name   string       `json:"name"`
	Users  map[int]bool `json:"users"`
}

type tranChatRoom struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Users  string `json:"users"`
}

func AddChatRoom(ctx context.Context, name string, users []int) error {
	storeUsers, err := json.Marshal(users)
	if err != nil {
		return err
	}
	s := `INSERT INTO #__chatroom (name,users) VALUES (?,?)`
	_, err = global.MySQL.Exec(dbutil.Prefix(s), name, string(storeUsers))
	return err
}

func AddUserSToChatRoom(ctx context.Context, rid int, uids []int) (err error) {
	tx, err := global.MySQL.Begin()
	defer func() {
		if err != nil {
			global.Logger.Errorf(ctx, "something failed with error: %s\n", err)
			tx.Rollback() //事务回滚
			return
		}
		err = tx.Commit()
	}()
	// 获取用户列表
	usersStr := ""
	err = tx.QueryRow(dbutil.Prefix("SELECT users FROM #__chatroom WHERE rid=?"), rid).Scan(&usersStr)
	usersMap := make(map[int]bool)
	err = json.Unmarshal([]byte(usersStr), &usersMap)
	if err != nil {
		return err
	}
	for _, v := range uids {
		usersMap[v] = true
	}
	// 更新用户列表
	userBytes, err := json.Marshal(usersMap)
	if err != nil {
		return err
	}
	_, err = tx.Exec(dbutil.Prefix("UPDATE #__chatroom SET users=?"), string(userBytes))
	return
}

func GetChatRoomByRoomId(ctx context.Context, rid int) (*ChatRoom, error) {
	tmp := &tranChatRoom{}
	s := "SELECT * FROM #__chatroom WHERE rid=?"
	err := global.MySQL.Get(tmp, dbutil.Prefix(s), rid)
	if err != nil {
		return nil, err
	}
	// 解析用户map
	usersMap := make(map[int]bool)
	err = json.Unmarshal([]byte(tmp.Users), &usersMap)
	if err != nil {
		return nil, err
	}

	chatRoom := &ChatRoom{
		RoomID: rid,
		Name:   tmp.Name,
		Users:  usersMap,
	}
	return chatRoom, nil
}
