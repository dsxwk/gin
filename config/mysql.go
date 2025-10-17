package config

import (
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

// InitMysql 初始化mysql
func InitMysql() *gorm.DB {
	var b strings.Builder
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

	dsn := b.String()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		color.Red("❌ 数据库连接失败: %v", err)
		log.Fatal()
	}

	return db
}
