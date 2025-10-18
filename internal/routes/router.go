package routes

import (
	"RS-backend/internal/handlers"
	"RS-backend/internal/modbus"

	"github.com/gin-gonic/gin"
)

func SetupRouter(client *modbus.Client) *gin.Engine {
	r := gin.Default()

	//托管静态文件
	r.Static("/static", "./static")
	//路由重定向，访问根目录时跳转到index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	// REST API
	api := r.Group("/api")
	{
		api.GET("/realtime", handlers.GetRealtime(client))
	}

	// WebSocket
	r.GET("/ws", func(c *gin.Context) {
		handlers.WSHandler(c, client)
	})

	return r
}
