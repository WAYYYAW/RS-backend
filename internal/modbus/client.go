package modbus

import (
	"log"
	"sync"
	"time"

	"RS-backend/internal/database"

	"github.com/simonvetter/modbus"
)

// motorSpeed: any,             // 电机转速
//
//	strokesNumber: any,          // 冲程数
//	distance: any,               // 冲程长度
//	rodDensity: any,             // 抽油杆密度
//	transmissionRatio: any,      // 传动比
//	area: any,                   // 截面积
//	inclination: any,            // 安装倾角
//	pumpInsertionDepth: any,     // 泵下入深度
//	oilDensity: any              // 原油密度
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
	client *modbus.ModbusClient // 替换为 simonvetter 的客户端指针
	mu     sync.Mutex
	Data   Data
	addr   string
}

func NewClient(addr string) *Client {
	// simonvetter 库的 URL 需要明确协议头 tcp://
	url := "tcp://" + addr
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     url,
		Timeout: 3 * time.Second,
	})
	if err != nil {
		// 只有 URL 格式严重错误时才会在此报错
		log.Fatalf("初始化 Modbus 客户端失败: %v", err)
	}

	// 等同于旧版 handler.SlaveId = 1
	client.SetUnitId(1)

	return &Client{
		client: client,
		addr:   addr,
	}
}

func (c *Client) connect() error {
	// Open() 替代了 Connect()
	if err := c.client.Open(); err != nil {
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
			// 重点优化：ReadRegisters 直接返回 []uint16，省去了手动处理 binary.BigEndian 的痛苦
			regs, err := c.client.ReadRegisters(0, 11, modbus.HOLDING_REGISTER)
			if err != nil {
				log.Printf("Modbus读取失败: %v", err)

				// 尝试重新连接
				if err := c.client.Close(); err != nil {
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

			// 检查结果长度 (因为一次读了 11 个寄存器，所以长度应该就是 11)
			if len(regs) < 11 {
				log.Printf("Modbus返回数据不足，期望11个寄存器，实际: %d", len(regs))
				time.Sleep(interval)
				continue
			}

			// 直接从 uint16 转换为 float64,雅
			position := float64(regs[0])
			load := float64(regs[1])
			motorSpeed := float64(regs[2])
			strokesNumber := float64(regs[3])
			distance := float64(regs[4])
			rodDensity := float64(regs[5])
			transmissionRatio := float64(regs[6])
			area := float64(regs[7])
			inclination := float64(regs[8])
			pumpInsertionDepth := float64(regs[9])
			oilDensity := float64(regs[10])

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

			log.Printf("读取到数据: Position=%f, Load=%f,MotorSpeed=%f,StrokesNumber=%f,Distance=%f",
				position, load, motorSpeed, strokesNumber, distance)
			time.Sleep(interval)
		}
	}()
}

func (c *Client) GetData() Data {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Data
}
