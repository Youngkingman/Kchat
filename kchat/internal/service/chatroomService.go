package service

import (
	"context"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
)

type ChatRoom struct {
	ChatRoom  *model.ChatRoom
	BroadCast *broadcaster
}

// 根据房间号索引对应的房间
// 1.先校验用户是否注册过房间以及房间是否存在
// 2.启动时直接加载所有房间，项目小问题不大，项目大emmmm分布式分配聊天房间
// 缺少用户注册接口，测试先自己强开一个
// 跨域问题需要解决,不然整不动vue-cli
var ChatRoomMap map[int]*ChatRoom

func LoadChatRoom(ctx context.Context) error {
	rooms, err := model.GetAllChatRoom(ctx)
	if err != nil {
		global.Logger.Panicf(ctx, "load rooms failed with error %v", err)
		return err
	}
	for _, v := range rooms {
		cr := &ChatRoom{
			ChatRoom:  v,
			BroadCast: NewBroadCast(v.RoomID),
		}
		ChatRoomMap[v.RoomID] = cr
	}
	return nil
}
