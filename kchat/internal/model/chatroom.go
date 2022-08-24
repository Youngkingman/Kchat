package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

type ChatRoom struct {
	RoomID int          `json:"room_id"`
	Name   string       `json:"name"`
	Avatar string       `json:"img_url"`
	Users  map[int]bool `json:"users"`
}

type tranChatRoom struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Avatar string `json:"img_url"`
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
	err = tx.QueryRow(dbutil.Prefix("SELECT users FROM #__chatroom WHERE room_id=?"), rid).Scan(&usersStr)
	if err != nil {
		return err
	}
	usersArr := make([]int, 0)
	usersMap := make(map[int]bool)
	err = json.Unmarshal([]byte(usersStr), &usersArr)
	if err != nil {
		return err
	}
	for _, v := range usersArr {
		usersMap[v] = true
	}
	for _, v := range uids {
		if ok, _ := usersMap[v]; !ok {
			usersArr = append(usersArr, v)
		}
	}
	// 更新用户列表
	userBytes, err := json.Marshal(usersArr)
	if err != nil {
		return err
	}
	_, err = tx.Exec(dbutil.Prefix("UPDATE #__chatroom SET users=? WHERE room_id=?"), string(userBytes), rid)
	return
}

func GetChatRoomByRoomId(ctx context.Context, rid int) (*ChatRoom, error) {
	tmp := tranChatRoom{}
	s := "SELECT * FROM #__chatroom WHERE room_id=?"
	err := global.MySQL.Get(&tmp, dbutil.Prefix(s), rid)
	if err != nil {
		fmt.Println(dbutil.Prefix(s))
		return nil, err
	}
	// 解析用户map
	usersMap := make(map[int]bool)
	userArr := make([]int, 0)
	err = json.Unmarshal([]byte(tmp.Users), &userArr)
	if err != nil {
		return nil, err
	}
	for _, v := range userArr {
		usersMap[v] = true
	}

	chatRoom := &ChatRoom{
		RoomID: rid,
		Name:   tmp.Name,
		Users:  usersMap,
	}
	return chatRoom, nil
}

func GetAllChatRoom(ctx context.Context) ([]*ChatRoom, error) {
	tmp := make([]tranChatRoom, 0)
	s := "SELECT * FROM #__chatroom"
	err := global.MySQL.Select(&tmp, dbutil.Prefix(s))
	if err != nil {
		return nil, err
	}
	ret := make([]*ChatRoom, 0)
	for _, v := range tmp {
		usersMap := make(map[int]bool)
		userArr := make([]int, 0)
		err = json.Unmarshal([]byte(v.Users), &userArr)
		if err != nil {
			return nil, err
		}
		for _, v := range userArr {
			usersMap[v] = true
		}
		chatRoom := &ChatRoom{
			RoomID: v.RoomID,
			Name:   v.Name,
			Users:  usersMap,
		}
		ret = append(ret, chatRoom)
	}
	return ret, nil
}

func GetAllUserFromChatRoom(ctx context.Context, rid int) (users []*User, err error) {
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
	err = tx.QueryRow(dbutil.Prefix("SELECT users FROM #__chatroom WHERE room_id=?"), rid).Scan(&usersStr)
	if err != nil {
		return nil, err
	}
	usersArr := make([]int, 0)
	err = json.Unmarshal([]byte(usersStr), &usersArr)
	if err != nil {
		return nil, err
	}
	for _, v := range usersArr {
		user := &User{}
		err = tx.QueryRow(dbutil.Prefix("SELECT * FROM #__user WHERE uid=?"), v).Scan(&user.UID, &user.Name, &user.Email, &user.Password, &user.ImageURL, &user.Website)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}
	return
}
