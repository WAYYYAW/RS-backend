package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳api
func GetRealtime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // 返回数据格式
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"timestamp": time.Now().Unix(),
			//直观的时间格式
			"realtime": time.Now().Format("2006-01-02 15:04:05"), //大坑！时间的固定参数
			//"rpm":         120.5,
			//"pressure":    1.03,
			//"temperature": 32.7,
		},
	})
}
