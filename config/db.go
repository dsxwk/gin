package config

import (
	"fmt"
	"gin/common/ctxkey"
	"gin/pkg"
	"gin/pkg/debugger"
	"gin/pkg/message"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// GetDB 初始化数据库(统一入口)
func GetDB() *gorm.DB {
	var (
		err error
	)

	dbOnce.Do(func() {
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
			color.Red(pkg.Error+"  不支持的数据库类型: %s", Conf.Databases.DbConnection)
			os.Exit(1)
		}

		if err != nil {
			color.Red(pkg.Error+"  %s数据库连接失败: %v", Conf.Databases.DbConnection, err)
			os.Exit(1)
		}
	})

	// 注册gorm sql回调
	SqlCallback(db)

	return db
}

// SqlCallback sql回调
func SqlCallback(db *gorm.DB) {
	// 查询
	_ = db.Callback().Query().Before("gorm:query").Register("log:before_query", before)
	_ = db.Callback().Query().After("gorm:query").Register("log:after_query", after)

	// 创建
	_ = db.Callback().Create().Before("gorm:create").Register("log:before_create", before)
	_ = db.Callback().Create().After("gorm:create").Register("log:after_create", after)

	// 更新
	_ = db.Callback().Update().Before("gorm:update").Register("log:before_update", before)
	_ = db.Callback().Update().After("gorm:update").Register("log:after_update", after)

	// 删除
	_ = db.Callback().Delete().Before("gorm:delete").Register("log:before_delete", before)
	_ = db.Callback().Delete().After("gorm:delete").Register("log:after_delete", after)
}

func before(db *gorm.DB) {
	db.InstanceSet(startTimeKey, time.Now())
}

func after(db *gorm.DB) {
	ctx := db.Statement.Context
	start, ok := db.InstanceGet(startTimeKey)
	if !ok {
		return
	}

	// 耗时
	cost := time.Since(start.(time.Time))
	costMs := float64(cost.Nanoseconds()) / 1e6 // 精确到小数
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	// 慢查询警告
	if cost > Conf.Mysql.SlowQueryDuration {
		GetLogger().Warn(
			"Slow Sql",
			zap.Float64("costMs", costMs),
			zap.String("sql", sql),
		)
	}

	message.GetEventBus().Publish(debugger.TopicSql, debugger.SqlEvent{
		TraceId: ctx.Value(ctxkey.TraceIdKey).(string),
		Sql:     sql,
		Rows:    db.Statement.RowsAffected,
		Ms:      costMs,
	})
}

// getSql 替换Sql中的占位符"?"为实际值
func getSql(sql string, vars []interface{}) string {
	for _, v := range vars {
		// 将参数值格式化为字符串
		var (
			formattedValue string
		)
		switch value := v.(type) {
		case string:
			formattedValue = fmt.Sprintf("'%s'", value)
		case []byte:
			formattedValue = fmt.Sprintf("'%s'", string(value))
		case time.Time:
			formattedValue = fmt.Sprintf("'%s'", value.Format("2006-01-02 15:04:05"))
		case *time.Time:
			if value != nil {
				formattedValue = fmt.Sprintf("'%s'", value.Format("2006-01-02 15:04:05"))
			} else {
				formattedValue = "NULL"
			}
		case *gorm.DeletedAt:
			if value == nil {
				formattedValue = "NULL"
			}
		default:
			formattedValue = fmt.Sprintf("%v", value)
		}

		// 替换第一个"?"为实际值
		sql = strings.Replace(sql, "?", formattedValue, 1)
	}

	return sql
}
