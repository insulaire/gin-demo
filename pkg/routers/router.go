package routers

import (
	"gin-demo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.Use(middleware.Tracing())

	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}
