package database

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 全局变量
var DB *gorm.DB

// 数据结构（表）
type Point struct {
	ID                 uint `gorm:"primaryKey"`
	Time               time.Time
	Position           float64
	Load               float64
	Inclination        float64 // 安装倾角
	MotorSpeed         float64 // 电机转速
	OilDensity         float64 // 原油密度
	PumpInsertionDepth float64 // 泵下入深度
	RodDensity         float64 // 抽油杆密度
	StrokesNumber      float64 // 冲程数
	Distance           float64 // 冲程长度
	TransmissionRatio  float64 // 传动比
	Area               float64 // 截面积
}

// 初始化数据库
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ 数据库连接失败: %v", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(&Point{})
	if err != nil {
		log.Fatalf("❌ 自动建表失败: %v", err)
	}

	log.Println("✅ 数据库已初始化！")
}

// 保存点数据
func SavePoint(p Point) {
	DB.Create(&p)
}
