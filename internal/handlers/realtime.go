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
				"timestamp":          time.Now().Unix(),
				"realtime":           time.Now().Format("2006-01-02 15:04:05"), //大坑！时间的固定参数
				"position":           data.Position,
				"load":               data.Load,
				"motorSpeed":         data.MotorSpeed,
				"strokesNumber":      data.StrokesNumber,
				"distance":           data.Distance,
				"time":               data.Time,
				"rodDensity":         data.RodDensity,
				"transmissionRatio":  data.TransmissionRatio,
				"area":               data.Area,
				"inclination":        data.Inclination,
				"pumpInsertionDepth": data.PumpInsertionDepth,
				"oilDensity":         data.OilDensity,
			},
		})
	}
}

//motorSpeed: any,             // 电机转速
//	strokesNumber: any,          // 冲程数
//		distance: any,               // 冲程长度
//		rodDensity: any,             // 抽油杆密度
//		transmissionRatio: any,      // 传动比
//		area: any,                   // 截面积
//		inclination: any,            // 安装倾角
//		pumpInsertionDepth: any,     // 泵下入深度
//		oilDensity: any              // 原油密度
