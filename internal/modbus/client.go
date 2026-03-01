package modbus

import (
	"encoding/binary"
	"log"
	"sync"
	"time"

	"RS-backend/internal/database"

	"github.com/goburrow/modbus"
)

// motorSpeed: any,             // 电机转速
//
//	strokesNumber: any,          // 冲程数
//		distance: any,               // 冲程长度
//		rodDensity: any,             // 抽油杆密度
//		transmissionRatio: any,      // 传动比
//		area: any,                   // 截面积
//		inclination: any,            // 安装倾角
//		pumpInsertionDepth: any,     // 泵下入深度
//		oilDensity: any              // 原油密度
type Data struct {
	Time               string  `json:"Time"`
	Position           float64 `json:"Position"`
	Load               float64 `json:"Load"`
	MotorSpeed         float64 `json:"MotorSpeed"`
	StrokesNumber      float64 `json:"StrokesNumber"`
	Distance           float64 `json:"Distance"`
	RodDensity         float64 `json:"RodDensity"`
	TransmissionRatio  float64 `json:"TransmissionRatio"`
	Area               float64 `json:"Area"`
	Inclination        float64 `json:"Inclination"`
	PumpInsertionDepth float64 `json:"PumpInsertionDepth"`
	OilDensity         float64 `json:"OilDensity"`
}

type Client struct {
	handler *modbus.TCPClientHandler
	client  modbus.Client
	mu      sync.Mutex
	Data    Data
	addr    string
}

func NewClient(addr string) *Client {
	handler := modbus.NewTCPClientHandler(addr)
	handler.Timeout = 3 * time.Second
	handler.SlaveId = 1

	return &Client{
		handler: handler,
		client:  modbus.NewClient(handler),
		addr:    addr,
	}
}

func (c *Client) connect() error {
	if err := c.handler.Connect(); err != nil {
		log.Printf("无法连接PLC: %v", err)
		return err
	}
	log.Println("成功连接到PLC")
	return nil
}

func (c *Client) Poll(interval time.Duration) {
	// 初始连接
	if err := c.connect(); err != nil {
		log.Printf("初始连接失败: %v", err)
	}

	go func() {
		for {
			// 从PLC读取11个寄存器
			results, err := c.client.ReadHoldingRegisters(0, 11)
			if err != nil {
				log.Printf("Modbus读取失败: %v", err)
				// 尝试重新连接
				err := c.handler.Close()
				if err != nil {
					log.Printf("断开连接失败: %v", err)
				}
				time.Sleep(2 * time.Second)
				if err := c.connect(); err != nil {
					log.Printf("重新连接失败: %v", err)
				} else {
					log.Println("重新连接成功")
				}
				time.Sleep(interval)
				continue
			}

			// 检查结果长度
			if len(results) < 4 {
				log.Printf("Modbus返回数据不足: %v", results)
				time.Sleep(interval)
				continue
			}

			// 从寄存器读取uint16值并转换为float64
			position := float64(binary.BigEndian.Uint16(results[0:2]))
			load := float64(binary.BigEndian.Uint16(results[2:4]))
			motorSpeed := float64(binary.BigEndian.Uint16(results[4:6]))
			strokesNumber := float64(binary.BigEndian.Uint16(results[6:8]))
			distance := float64(binary.BigEndian.Uint16(results[8:10]))
			rodDensity := float64(binary.BigEndian.Uint16(results[10:12]))
			transmissionRatio := float64(binary.BigEndian.Uint16(results[12:14]))
			area := float64(binary.BigEndian.Uint16(results[14:16]))
			inclination := float64(binary.BigEndian.Uint16(results[16:18]))
			pumpInsertionDepth := float64(binary.BigEndian.Uint16(results[18:20]))
			oilDensity := float64(binary.BigEndian.Uint16(results[20:22]))
			c.mu.Lock()
			c.Data = Data{
				Time:               time.Now().Format("2006-01-02 15:04:05"),
				Position:           position,
				Load:               load,
				MotorSpeed:         motorSpeed,
				StrokesNumber:      strokesNumber,
				Distance:           distance,
				RodDensity:         rodDensity,
				TransmissionRatio:  transmissionRatio,
				Area:               area,
				Inclination:        inclination,
				PumpInsertionDepth: pumpInsertionDepth,
				OilDensity:         oilDensity,
			}
			c.mu.Unlock()

			// 将数据保存到数据库
			point := database.Point{
				Time:     time.Now(),
				Position: position,
				Load:     load,
				//暂时不保存电机速度等数据
			}
			database.SavePoint(point)

			log.Printf("读取到数据: Position=%f, Load=%f,MotorSpeed=%f,StrokesNumber=%f,Distance=%f", position, load, motorSpeed, strokesNumber, distance)
			time.Sleep(interval)
		}
	}()
}

func (c *Client) GetData() Data {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Data
}
