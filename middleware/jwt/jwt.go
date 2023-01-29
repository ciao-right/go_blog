package middleware

import (
	"github.com/gin-gonic/gin"
	"go_blog/utils"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := "success"
		token := c.Request.Header.Get("token")
		if token == "" {
			message = "token为空"
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    201,
				"message": message,
				"data":    0,
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			message = "未知错误"
		} else if utils.GetNow().Unix() > claims.ExpiresAt {
			message = "token已经过期"
		}
		if message != "success" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    201,
				"message": message,
				"data":    0,
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
