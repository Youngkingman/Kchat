package app

import (
	"time"

	"github.com/Youngkingman/Kchat/kchat/global"
	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/dgrijalva/jwt-go"
)

// 需要修改带上用户信息
type Claims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(user *model.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	//如果此处采用的是RSA256的话，应该在Generate的SingedString使用rsa.PrivateKey
	//在parse(Validation)处采用rsa.PublicKey //TODO
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//如果是非对称加密这里的GetJWTSecret要换成公钥
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
