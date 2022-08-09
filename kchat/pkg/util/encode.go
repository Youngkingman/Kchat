package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/Youngkingman/Kchat/kchat/global"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// 不是什么敏感信息简单混杂一下
func EncodeUID(uid int) (ret string) {
	uidStr := strconv.Itoa(uid)
	ToEncode := fmt.Sprintf("%s.%s", uidStr, global.AppSetting.UidSalt)
	ret = base64.StdEncoding.EncodeToString([]byte(ToEncode))
	return
}

func DecodeUID(uidStr string) (ret int, err error) {
	decoded, err := base64.StdEncoding.DecodeString(uidStr)
	if err != nil {
		return
	}
	mixedStr := strings.Split(string(decoded), ".")
	ret, err = strconv.Atoi(mixedStr[0])
	return
}
