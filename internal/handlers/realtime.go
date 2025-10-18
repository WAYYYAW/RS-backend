package handlers

import (
	"RS-backend/internal/modbus"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳api和最近一次数据api
func GetRealtime(client *modbus.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := client.GetData()

		c.JSON(http.StatusOK, gin.H{ // 返回数据格式
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"timestamp": time.Now().Unix(),
				"realtime":  time.Now().Format("2006-01-02 15:04:05"), //大坑！时间的固定参数
				"position":  data.Position,
				"load":      data.Load,
			},
		})
	}
}
