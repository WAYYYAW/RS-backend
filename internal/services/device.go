package services

import (
	"RS-backend/internal/modbus"
	"RS-backend/internal/models"
)

// DeviceService 负责封装设备连接相关业务。
type DeviceService struct {
	client *modbus.Client
}

func NewDeviceService(client *modbus.Client) *DeviceService {
	return &DeviceService{client: client}
}

func (s *DeviceService) ConnectDevice(_ models.DeviceConnectionRequest) models.DeviceConnectionResponse {
	return models.DeviceConnectionResponse{
		Status: 200,
		Data: models.DeviceParams{
			MotorSpeed:         "1200",
			StrokesNumber:      "5",
			Distance:           "3.5",
			RodDensity:         "7850",
			TransmissionRatio:  "50",
			Area:               "0.001",
			Inclination:        "30",
			PumpInsertionDepth: "1000",
			OilDensity:         "850",
		},
	}
}
