package middware

import (
	"github.com/gin-gonic/gin"
)

func JsonWebToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的token
		Authorization := c.Request.Header.Get("authorization")
		// 判断token是否为空
		if Authorization == "" {
			c.JSON(401, gin.H{
				"message": "Authorization is empty",
			})
			c.Abort()
			return
		} else {
			// 判断token是否正确
			if Authorization == "123456" {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Authorization is wrong",
				})
				c.Abort()
				return
			}
		}
	}
}
