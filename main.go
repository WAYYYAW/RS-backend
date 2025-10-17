package main

import (
	"RS-backend/internal/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // 监听 8080 端口
}
