package handlers

import (
	"RS-backend/internal/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const maxHistoryRangeLimit = 1000

func historyPointResponse(point database.Point) gin.H {
	return gin.H{
		"id":                 point.ID,
		"timestamp":          point.Time.Unix(),
		"time":               point.Time.Format("2006-01-02 15:04:05"),
		"position":           point.Position,
		"load":               point.Load,
		"inclination":        point.Inclination,
		"motorSpeed":         point.MotorSpeed,
		"oilDensity":         point.OilDensity,
		"pumpInsertionDepth": point.PumpInsertionDepth,
		"rodDensity":         point.RodDensity,
		"strokesNumber":      point.StrokesNumber,
		"transmissionRatio":  point.TransmissionRatio,
		"area":               point.Area,
		"distance":           point.Distance,
	}
}

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
	result := database.DB.Where("time >= ? AND time < ?", t, t.Add(time.Second)).Order("time ASC").First(&point) //再次感谢Lingma老师的帮助，时间戳令人头大
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
		"data": historyPointResponse(point),
	})
}

// GetHistoryRange 根据起止时间戳获取时间段历史数据。
func GetHistoryRange(c *gin.Context) {
	startParam := c.Query("start")
	endParam := c.Query("end")
	if startParam == "" || endParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "缺少开始或结束时间戳参数",
		})
		return
	}

	startUnix, err := strconv.ParseInt(startParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "开始时间戳格式错误，应为Unix时间戳格式",
		})
		return
	}

	endUnix, err := strconv.ParseInt(endParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "结束时间戳格式错误，应为Unix时间戳格式",
		})
		return
	}

	if endUnix < startUnix {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "结束时间不能早于开始时间",
		})
		return
	}

	limit := maxHistoryRangeLimit
	if limitParam := c.Query("limit"); limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err != nil || parsedLimit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  "limit必须是正整数",
			})
			return
		}
		if parsedLimit < limit {
			limit = parsedLimit
		}
	}

	startTime := time.Unix(startUnix, 0)
	endTime := time.Unix(endUnix, 0)

	var points []database.Point
	result := database.DB.
		Where("time >= ? AND time <= ?", startTime, endTime).
		Order("time ASC").
		Limit(limit).
		Find(&points)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "查询历史数据失败",
		})
		return
	}

	data := make([]gin.H, 0, len(points))
	for _, point := range points {
		data = append(data, historyPointResponse(point))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
		"meta": gin.H{
			"count": len(data),
			"limit": limit,
			"start": startUnix,
			"end":   endUnix,
		},
	})
}
