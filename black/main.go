package main

import (
	"strings"

	"github.com/BeLEEU/blackretire/black/interna/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		//允许cookie session
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return origin == "未来可能的容器服务器"
		},
		MaxAge: 12 * time.Hour,
	}))

	web := web.NewUserHandler()
	web.RegisterRoutes(server)

	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
