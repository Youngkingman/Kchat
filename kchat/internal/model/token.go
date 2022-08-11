package model

import (
	"fmt"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
)

func GetToken(key string) (token string, err error) {
	token, err = global.Redis.Get(dbutil.Prefix("#__" + key + "_token"))
	return
}

func SetToken(email string, token string) (err error) {
	trueKey := dbutil.Prefix(fmt.Sprintf("#__%s_token", email))
	err = global.Redis.Set(trueKey, token, global.JWTSetting.Expire)
	return
}

func DeleteToken(email string) (err error) {
	trueKey := dbutil.Prefix(fmt.Sprintf("#__%s_token", email))
	err = global.Redis.Del(trueKey)
	return
}
