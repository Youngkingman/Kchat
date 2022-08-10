package psw

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// 用户存储的密码应该包括了 {s1}.{s2}两部分, {s1}是加密文, {s2}是盐,每个人的盐都是随机生成的
func hashPassword(password string) (string, error) {
	// example for making salt - https://play.golang.org/p/_Aw6WeWC42I
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// using recommended cost parameters from - https://godoc.org/golang.org/x/crypto/scrypt
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// return hex-encoded string with salt appended to password
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")

	// check supplied password salted with hash
	salt, err := hex.DecodeString(pwsalt[1])

	if err != nil {
		return false, fmt.Errorf("Unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
