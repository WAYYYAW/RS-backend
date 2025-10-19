package main

import (
	"RS-backend/internal/database"
	"RS-backend/internal/modbus"
	"RS-backend/internal/routes"
	"time"
)

func main() {
	// 初始化数据库
	database.InitDB()
	// 创建Modbus客户端
	client := modbus.NewClient("127.0.0.1:5020")
	client.Poll(500 * time.Millisecond)

	// 启动Gin服务器
	r := routes.SetupRouter(client)
	r.Run(":8080")
}
