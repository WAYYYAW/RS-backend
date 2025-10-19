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
	ID       uint `gorm:"primaryKey"`
	Time     time.Time
	Position float64
	Load     float64
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
