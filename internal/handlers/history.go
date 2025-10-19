package handlers

import (
	"RS-backend/internal/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 根据时间戳获取历史数据
func GetHistory(c *gin.Context) {
	// 从查询参数获取时间戳
	timestamp := c.Query("timestamp")
	if timestamp == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "缺少时间戳参数",
		})
		return
	}

	// 解析Unix时间戳
	unixTime, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "时间戳格式错误，应为Unix时间戳格式（例如：1760869745）",
		})
		return
	}

	// 将Unix时间戳转换为时间对象
	t := time.Unix(unixTime, 0)

	var point database.Point
	// 修复查询语句，改为在一秒范围内查找记录
	result := database.DB.Where("time >= ? AND time < ?", t, t.Add(time.Second)).First(&point) //再次感谢Lingma老师的帮助，时间戳令人头大
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 1,
			"msg":  "未找到指定时间的数据",
		})
		return
	}

	// 返回查询到的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"id":        point.ID,
			"timestamp": point.Time.Unix(),
			"time":      point.Time.Format("2006-01-02 15:04:05"),
			"position":  point.Position,
			"load":      point.Load,
		},
	})
}
