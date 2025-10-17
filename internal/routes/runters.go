package routes

import (
	"RS-backend/internal/data"
	"RS-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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
		api.GET("/realtime", handlers.GetRealtime)
	}

	// WebSocket
	//r.GET("/ws", handlers.WSHandler)
	dataCh := make(chan data.Point, 100)
	go data.SimulateData(dataCh, "data.csv")
	r.GET("/ws", func(c *gin.Context) {
		handlers.WSHandler(c, dataCh)
	})
	return r
}
