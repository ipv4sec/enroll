package middleware

import (
	"enroll/redis"
	"github.com/gin-gonic/gin"
)

func TokenChecker(exceptions []string) gin.HandlerFunc {
	return func (c *gin.Context) {
		for _, exception := range exceptions {
			if c.Request.URL.Path == exception {
				c.Next()
				return
			}
		}

		token := c.Request.Header.Get("X-Access-Token")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Header字段的X-Access-Token为空",
				"message": "未发现认证信息",
			})
			return
		}

		uid, err := redis.FindUidByToken(token)
		if err != nil || uid == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Redis不存在此Token值",
				"message": "认证信息过期或不存在",
			})
			return
		}
		c.Set("adminId", uid)
		c.Next()
	}
}
