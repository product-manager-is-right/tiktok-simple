package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dsn = "root1:root1234@tcp(120.25.2.146:3306)/tiktok_ums?charset=utf8&parseTime=True&loc=Local"
var dsv = "root1:root1234@tcp(120.25.2.146:3306)/tiktok_vms?charset=utf8&parseTime=True&loc=Local"
var DB *gorm.DB
var DBV *gorm.DB

func Init() {
	// gorm日志打印
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			Colorful:      true,          // 彩色打印
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	DBV, err = gorm.Open(mysql.Open(dsv), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
