package model

import "time"

type Chatter struct {
	RoomID   int       `json:"room_id"`
	UID      int       `json:"uid"`
	Addr     string    `json:"addr"` //ip地址，用来识别罕见和五十万
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	ImageURL string    `json:"img_url"`
	EnterAt  time.Time `json:"enter_at"`
}

func NewChatterFromUser(u *User, rid int) *Chatter {
	return &Chatter{
		RoomID:   rid,
		UID:      u.UID,
		Email:    u.Email,
		Name:     u.Name,
		ImageURL: u.ImageURL,
	}
}

// 保存用户在房间发表的数据记录，配合message表进行，出于个人喜好不打算留记录
