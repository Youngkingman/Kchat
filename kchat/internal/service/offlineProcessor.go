package service

// type offlineProcessor struct {
// 	n int

// 	// 所有用户最近n条信息
// 	recentRing *ring.Ring

// 	//某个用户最近n条信息
// 	userRing map[string]*ring.Ring
// }

// var OfflineProcessor = newOfflinering()

// func newOfflinering() *offlineProcessor {
// 	n := global.ChatRoomSetting.OfflineNum

// 	return &offlineProcessor{
// 		n:          n,
// 		recentRing: ring.New(n),
// 		userRing:   make(map[string]*ring.Ring),
// 	}
// }
