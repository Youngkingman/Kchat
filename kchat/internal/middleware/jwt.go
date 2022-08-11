package middleware

import (
	"fmt"

	"github.com/Youngkingman/Kchat/kchat/internal/model"
	"github.com/Youngkingman/Kchat/kchat/pkg/app"
	"github.com/Youngkingman/Kchat/kchat/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claim, err := app.ParseToken(token)
			tokenInRedis, errRedis := model.GetToken(claim.User.Email)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
			fmt.Println(tokenInRedis)
			// 系统时钟的细微差别使得它们可能有那么一丝可能Redis过期了token没过期，用户多了就是必然
			if errRedis != nil {
				ecode = errcode.UnauthorizedTokenError
			} else if ecode == errcode.UnauthorizedTokenTimeout {
				model.DeleteToken(claim.User.Email)
			}
			// 应用这个中间件的路由都可以从contex得到用户的信息,email一定是有的,不保证是最新信息
			c.Set("user", claim.User)
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
