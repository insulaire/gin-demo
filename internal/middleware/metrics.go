package middleware

import (
	"gin-demo/pkg/metrics"

	"github.com/gin-gonic/gin"
)

func RequestCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.Inc(c)
	}
}
