package main

import (
	"fmt"
	"gin-demo/global"
	"gin-demo/internal/routers"
	"gin-demo/pkg/consul"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	routers.InitRouter(g)
	consul.Registra()
	g.Run(fmt.Sprintf("%s:%d", global.ServiceSetting.Host, global.ServiceSetting.Port))
}
