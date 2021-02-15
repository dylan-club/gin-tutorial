package middleware

import (
	"com.nicklaus/ginpractice/common"
	"com.nicklaus/ginpractice/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取header
		tokenString := context.GetHeader("Authorization")

		//判断token是否有效
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		//判断token是否解析成功
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		db := common.GetDB()
		user := dao.FindUserById(db, claims.UserId)
		//如果用户不存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		context.Set("user", user)
		context.Next()
	}
}
