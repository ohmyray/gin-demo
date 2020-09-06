package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/model"
	"github.com/ohmyray/gin-demo01/response"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Fail(ctx, nil,"权限不足！")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Fail(ctx, nil, "权限不足！")
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
			//ctx.JSON(http.StatusUnauthorized, gin.H{
			//	"code":    5001,
			//	"message": "用户不存在！",
			//})
			response.Fail(ctx, nil, "用户不存在!")
			ctx.Abort()
			return
		}

		// 用户存在 将 user 信息下入上下文
		ctx.Set("user", user)

		ctx.Next()

	}
}
