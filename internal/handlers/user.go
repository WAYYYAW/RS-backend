package handlers

import (
	"RS-backend/internal/modbus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConnectionInfo 上位机连接信息结构体
type ConnectionInfo struct {
	ID   interface{} `json:"id"`
	IP   interface{} `json:"ip"`
	Port interface{} `json:"port"`
}

// DeviceParams 设备参数结构体
type DeviceParams struct {
	MotorSpeed         interface{} `json:"motorSpeed"`         // 电机转速
	StrokesNumber      interface{} `json:"strokesNumber"`      // 冲程数
	Distance           interface{} `json:"distance"`           // 冲程长度
	RodDensity         interface{} `json:"rodDensity"`         // 抽油杆密度
	TransmissionRatio  interface{} `json:"transmissionRatio"`  // 传动比
	Area               interface{} `json:"area"`               // 截面积
	Inclination        interface{} `json:"inclination"`        // 安装倾角
	PumpInsertionDepth interface{} `json:"pumpInsertionDepth"` // 泵下入深度
	OilDensity         interface{} `json:"oilDensity"`         // 原油密度
}

// DeviceData 设备数据响应结构体
type DeviceData struct {
	Status int          `json:"status"`
	Data   DeviceParams `json:"data"`
}

// ConnectDevice 连接设备处理函数
func ConnectDevice(client *modbus.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var connInfo ConnectionInfo
		if err := c.ShouldBindJSON(&connInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "Invalid connection info",
			})
			return
		}

		// 这里可以使用client参数进行实际的设备连接
		// 模拟连接成功并返回设备参数
		response := DeviceData{
			Status: 200, // 正常发200~299
			Data: DeviceParams{
				MotorSpeed:         "1200",  // 电机转速
				StrokesNumber:      "5",     // 冲程数
				Distance:           "3.5",   // 冲程长度
				RodDensity:         "7850",  // 抽油杆密度
				TransmissionRatio:  "50",    // 传动比
				Area:               "0.001", // 截面积
				Inclination:        "30",    // 安装倾角
				PumpInsertionDepth: "1000",  // 泵下入深度
				OilDensity:         "850",   // 原油密度
			},
		}

		c.JSON(http.StatusOK, response)
	}
}
