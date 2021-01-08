package routers

import (
	"gin-demo/internal/handlers"
	"gin-demo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.Use(middleware.Recovery())
	g.Use(middleware.AccessLog())
	g.Use(middleware.Tracing())
	g.GET("/ping", handlers.Ping)
}
