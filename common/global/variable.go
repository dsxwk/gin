package global

import (
	"gin/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config // 配置
	DB     *gorm.DB       // 数据库
	Log    *config.Logger // 日志
)

func init() {
	Config = config.InitConfig()
	DB = config.InitMysql()
	Log = config.InitLog()
}
