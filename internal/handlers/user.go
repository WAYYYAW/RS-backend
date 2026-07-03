package handlers

import (
	"RS-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deviceConnectService interface {
	ConnectDevice(info models.DeviceConnectionRequest) models.DeviceConnectionResponse
}

type UserHandler struct {
	deviceService deviceConnectService
}

func NewUserHandler(deviceService deviceConnectService) *UserHandler {
	return &UserHandler{deviceService: deviceService}
}

// ConnectDevice 负责处理设备连接请求。
func (h *UserHandler) ConnectDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var connInfo models.DeviceConnectionRequest
		if err := c.ShouldBindJSON(&connInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "Invalid connection info",
			})
			return
		}

		c.JSON(http.StatusOK, h.deviceService.ConnectDevice(connInfo))
	}
}
