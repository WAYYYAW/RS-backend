package modbus

import (
	"encoding/binary"
	"log"
	"sync"
	"time"

	"RS-backend/internal/database"

	"github.com/goburrow/modbus"
)

type Data struct {
	Time     string  `json:"time"`
	Position float64 `json:"position"`
	Load     float64 `json:"load"`
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
			// 从PLC读取2个寄存器 (Position和Load，每个值占用一个寄存器)
			results, err := c.client.ReadHoldingRegisters(0, 2)
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

			c.mu.Lock()
			c.Data = Data{
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Position: position,
				Load:     load,
			}
			c.mu.Unlock()

			// 将数据保存到数据库
			point := database.Point{
				Time:     time.Now(),
				Position: position,
				Load:     load,
			}
			database.SavePoint(point)

			log.Printf("读取到数据: Position=%f, Load=%f", position, load)
			time.Sleep(interval)
		}
	}()
}

func (c *Client) GetData() Data {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Data
}
