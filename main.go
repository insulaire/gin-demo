package main

import (
	_ "gin-demo/global"
	"gin-demo/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	routers.InitRouter(g)
	g.Run(":9999")
}
