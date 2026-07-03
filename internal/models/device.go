package models

// DeviceConnectionRequest 表示设备连接请求参数。
type DeviceConnectionRequest struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// DeviceParams 表示设备基础参数。
type DeviceParams struct {
	MotorSpeed         string `json:"motorSpeed"`
	StrokesNumber      string `json:"strokesNumber"`
	Distance           string `json:"distance"`
	RodDensity         string `json:"rodDensity"`
	TransmissionRatio  string `json:"transmissionRatio"`
	Area               string `json:"area"`
	Inclination        string `json:"inclination"`
	PumpInsertionDepth string `json:"pumpInsertionDepth"`
	OilDensity         string `json:"oilDensity"`
}

// DeviceConnectionResponse 表示连接设备接口返回值。
type DeviceConnectionResponse struct {
	Status int          `json:"status"`
	Data   DeviceParams `json:"data"`
}
