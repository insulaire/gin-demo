package middleware

import (
	"gin-demo/global"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
			c.Abort()
		}
		c.Next()
	}
}
