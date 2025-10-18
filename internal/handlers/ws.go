package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"RS-backend/internal/modbus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(c *gin.Context, client *modbus.Client) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		data := client.GetData()
		msg, _ := json.Marshal(data)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}
