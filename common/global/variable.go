package global

import (
	"gin/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB // 数据库
)

func init() {
	DB = config.InitMysql()
}
