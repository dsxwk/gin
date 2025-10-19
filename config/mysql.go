package config

import (
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

// InitMysql 初始化mysql
func InitMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局关闭表名复数化
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
			logger.Config{
				SlowThreshold: time.Millisecond * Conf.Mysql.SlowQuerySeconds, // 慢SQL阈值
				LogLevel:      logger.Info,
				Colorful:      true, // 彩色日志
				// IgnoreRecordNotFoundError: true, // 如果需要忽略 record not found
			},
		),
	})
	if err != nil {
		color.Red("❌ 数据库连接失败: %v", err)
		log.Fatal()
	}

	return db
}

// getDsn 获取数据库dns
func getDsn() string {
	var (
		b strings.Builder
	)

	// 预分配容量
	b.Grow(128)

	Init()
	b.WriteString(Conf.Mysql.Username)
	b.WriteString(":")
	b.WriteString(Conf.Mysql.Password)
	b.WriteString("@tcp(")
	b.WriteString(Conf.Mysql.Host)
	b.WriteString(":")
	b.WriteString(Conf.Mysql.Port)
	b.WriteString(")/")
	b.WriteString(Conf.Mysql.Database)
	b.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	return b.String()
}
