package routers

import (
	"gin-demo/internal/handlers"
	"gin-demo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter(g *gin.Engine) {
	g.Use(middleware.Recovery())
	g.Use(middleware.AccessLog())
	g.Use(middleware.Tracing())
	g.Use(middleware.RequestCount())
	g.GET("/ping", handlers.Ping)
	g.GET("/metrics", Warp(promhttp.Handler()))
}

func Warp(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
