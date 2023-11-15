package middleware

import (
	"blog/common"
	"blog/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//獲取Authorization header
		tokenString := c.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 401,
				"msg":  "TOKEN無效1",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 401,
				"msg":  "token無效2",
			})
			c.Abort()
			return
		}
		//驗證通過後獲取claim中的userId
		userId := claims.UserID
		var user model.User
		common.DB.First(&user, userId)
		//用戶不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 401,
				"msg":  "用戶不存在",
			})
			c.Abort()
			return
		}
		//用戶存在，將user信息寫入上下文
		c.Set("user", user)
		c.Next()
	}
}
