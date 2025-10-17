package handlers

import (
	"RS-backend/internal/data"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	//"math/rand"
	"net/http"
	//"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

//func WSHandler(c *gin.Context) {
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		return
//	}
//	defer conn.Close()
//
//	rand.Seed(time.Now().UnixNano())
//
//	for {
//		// 每秒推送一次数据
//		msg := map[string]interface{}{
//			"type": "update",
//			"data": map[string]interface{}{
//				"timestamp": time.Now().Unix(),
//				"rpm":       100 + rand.Float64()*100, //模拟数据：转速
//				"pressure":  80 + rand.Float64()*80,   //模拟数据：压力
//			},
//		}
//		conn.WriteJSON(msg)
//		time.Sleep(time.Second) //秒
//	}
//}

func WSHandler(c *gin.Context, dataCh <-chan data.Point) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for p := range dataCh {
		conn.WriteJSON(p)
	}
}
