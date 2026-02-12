package config

import (
	"gin/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

const startTimeKey = "gorm_start_time"

func openMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getMysqlDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局关闭表名复数化
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
			logger.Config{
				SlowThreshold: Conf.Mysql.SlowQueryDuration, // 慢sql阈值转Duration
				LogLevel:      logger.Info,
				Colorful:      true, // 彩色日志
				// IgnoreRecordNotFoundError: true, // 如果需要忽略 record not found
			},
		),
	})

	return db, err
}

// getMysqlDsn 获取数据库dns
func getMysqlDsn() string {
	return pkg.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		Conf.Mysql.Username, Conf.Mysql.Password, Conf.Mysql.Host, Conf.Mysql.Port, Conf.Mysql.Database,
	)
}
