package config

import (
	"fmt"
	"gin/utils"
	"gin/utils/debugger"
	"gin/utils/message"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

var (
	DB *gorm.DB
)

const startTimeKey = "gorm_start_time"

func init() {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局关闭表名复数化
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
			logger.Config{
				SlowThreshold: Conf.Mysql.SlowQueryDuration * time.Second, // 慢sql阈值转Duration
				LogLevel:      logger.Info,
				Colorful:      true, // 彩色日志
				// IgnoreRecordNotFoundError: true, // 如果需要忽略 record not found
			},
		),
	})
	if err != nil {
		color.Red(utils.Error+"  数据库连接失败: %v", err)
		os.Exit(1)
	}

	// sql回调
	SqlCallback(db)

	DB = db
}

// getDsn 获取数据库dns
func getDsn() string {
	return utils.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		Conf.Mysql.Username, Conf.Mysql.Password, Conf.Mysql.Host, Conf.Mysql.Port, Conf.Mysql.Database,
	)
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
		ZapLogger.Warn(
			"Slow Sql",
			zap.Float64("costMs", costMs),
			zap.String("sql", sql),
		)
	}

	message.MsgEventBus.Publish(debugger.TopicSql, debugger.SqlEvent{
		Sql:  sql,
		Rows: db.Statement.RowsAffected,
		Ms:   costMs,
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
