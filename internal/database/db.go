package database

import (
	"database/sql"
	"log"
	"time"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
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
	// 使用 modernc.org/sqlite (纯Go实现,无需cgo)
	sqlDB, err := sql.Open("sqlite", "data.db?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		log.Fatalf("❌ 数据库连接失败: %v", err)
	}

	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ 数据库初始化失败: %v", err)
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
