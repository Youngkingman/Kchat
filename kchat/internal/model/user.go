package model

import (
	"context"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

type User struct {
	UID      int    `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageURL string `json:"img_url"`
	Website  string `json:"website"`
}

func GetUserByUid(ctx context.Context, uid int) (u *User, err error) {
	u = &User{}
	err = global.MySQL.Get(u, dbutil.Prefix("SELECT * FROM #__user WHERE uid=?"), uid)
	return
}

func GetUserPasswordByEmail(ctx context.Context, email string) (psw string, err error) {
	err = global.MySQL.Get(psw, dbutil.Prefix("SELECT password FROM #__user WHERE email=?"), email)
	return
}

func GetUserByEmail(ctx context.Context, email string) (u *User, err error) {
	u = &User{}
	err = global.MySQL.Get(u, dbutil.Prefix("SELECT * FROM #__user WHERE email=?"), email)
	return
}

func AddUser(ctx context.Context, u *User) (err error) {
	s := `INSERT INTO #__user (name, email, password, img_url, website)  VALUES (?,?,?,?,?);`
	_, err = global.MySQL.Exec(dbutil.Prefix(s), u.Name, u.Email, u.Password, u.ImageURL, u.Website)
	return
}
