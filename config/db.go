package config

import (
	"gin/utils"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

// 初始化数据库(统一入口)
func init() {
	var (
		db  *gorm.DB
		err error
	)

	switch Conf.Databases.DbConnection {
	case "mysql":
		db, err = openMysql()
	case "pgsql":
		db, err = openPgsql()
	case "sqlite":
		db, err = openSqlite()
	case "sqlsrv":
		db, err = openSqlsrv()
	default:
		color.Red(utils.Error+"  不支持的数据库类型: %s", Conf.Databases.DbConnection)
		os.Exit(1)
	}

	if err != nil {
		color.Red(utils.Error+"  %s数据库连接失败: %v", Conf.Databases.DbConnection, err)
		os.Exit(1)
	}

	// 注册gorm sql回调
	SqlCallback(db)

	DB = db
}
