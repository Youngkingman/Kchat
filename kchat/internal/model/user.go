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
	ImageURL string `json:"img_url"`
	Website  string `json:"website"`
}

func GetUser(ctx context.Context, uid int) (u *User, err error) {
	u = &User{}
	err = global.MySQL.Get(u, dbutil.Prefix("SELECT * FROM #__user WHERE uid=?"), uid)
	global.Logger.Debugf(ctx, dbutil.Prefix("SELECT * FROM #__user WHERE uid=?"))
	return
}
