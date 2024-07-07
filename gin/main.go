package main

import (
	"github.com/BeLEEU/blackretire/black/interna/web"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	web := web.NewUserHandler()
	web.RegisterRoutes(server)
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
