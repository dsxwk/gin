package config

import (
	"gin/pkg"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func openSqlsrv() (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(getSqlsrvDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不复数
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: Conf.Sqlsrv.SlowQueryDuration,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
}

func getSqlsrvDsn() string {
	/*
	   官方推荐格式(最稳定)：
	   sqlserver://username:password@host:port?database=dbname
	   常见坑：
	   - password 有特殊字符需要 url.QueryEscape
	   - SQLServer 默认端口 1433
	*/

	return pkg.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		Conf.Sqlsrv.Username,
		Conf.Sqlsrv.Password,
		Conf.Sqlsrv.Host,
		Conf.Sqlsrv.Port,
		Conf.Sqlsrv.Database,
	)
}
