package model

import (
	"context"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

type User struct {
	UID      int    `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	Website  string `json:"website"`
}

func GetUser(ctx context.Context, uid int) (u *User, err error) {
	err = global.MySQL.Select(u, dbutil.Prefix("SELECT * FROM #__user WHERE uid=?"), uid)
	return
}
