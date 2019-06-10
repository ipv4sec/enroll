package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AdminChecker(include []string) gin.HandlerFunc {
	return func (c *gin.Context) {
		for _, exception := range include {
			if c.Request.URL.Path == exception {
				adminId, ok := c.Get("adminId")
				if ok && adminId.(int64) == 1 {
					c.Next()
					return
				}
				c.AbortWithStatusJSON(403, gin.H{
					"error": fmt.Sprintf("没有此权限的用户Id为:%v", adminId),
					"message": "用户没有此权限",
				})
				return
			}
			c.Next()
			return
		}
		c.Next()
		return
	}
}

