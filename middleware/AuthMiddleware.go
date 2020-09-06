package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    40004,
				"message": "权限不足！",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    40005,
				"message": "权限不足！",
			})
			ctx.Abort()
			return
		}

		// 验证通过后获取claims 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    5001,
				"message": "用户不存在！",
			})
			ctx.Abort()
			return
		}

		// 用户存在 将 user 信息下入上下文
		ctx.Set("user", user)

		ctx.Next()

	}
}
