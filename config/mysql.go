package config

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	SqlRes []string
	mu     sync.Mutex
	DB     *gorm.DB
)

func init() {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局关闭表名复数化
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
			logger.Config{
				SlowThreshold: Conf.Mysql.SlowQuerySeconds * time.Second, // 慢sql阈值转Duration
				LogLevel:      logger.Info,
				Colorful:      true, // 彩色日志
				// IgnoreRecordNotFoundError: true, // 如果需要忽略 record not found
			},
		),
	})
	if err != nil {
		color.Red("❌  数据库连接失败: %v", err)
		os.Exit(1)
	}

	// sql回调
	SqlCallback(db)

	DB = db
}

// getDsn 获取数据库dns
func getDsn() string {
	var (
		b strings.Builder
	)

	// 预分配容量
	b.Grow(128)

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

// SqlCallback sql回调
func SqlCallback(db *gorm.DB) {
	// 注册查询前回调
	_ = db.Callback().Query().Before("gorm:query").Register("slowquery:begin", func(db *gorm.DB) {
		// 记录查询开始时间
		db.InstanceSet("gorm:query_slowquery_start_time", time.Now())
	})

	// 注册查询后回调
	_ = db.Callback().Query().After("gorm:query").Register("slowquery:end", func(db *gorm.DB) {
		var (
			sql  = db.Statement.SQL.String()
			vars = db.Statement.Vars
		)

		// 追加sql数据
		AddSql(sql, vars)

		// 获取查询开始时间
		startTime, _ := db.InstanceGet("gorm:query_slowquery_start_time")
		if startTime != nil {
			elapsed := time.Since(startTime.(time.Time))
			// 超过阈值则记录慢查询
			if elapsed > Conf.Mysql.SlowQuerySeconds {
				// 记录慢查询的执行时间和查询语句
				ZapLogger.Warn(fmt.Sprintf("执行慢查询: %s, sql: %s", elapsed.String(), getSQL(sql, vars)))
			}
		}
	})

	// 注册创建
	_ = db.Callback().Create().After("gorm:create").Register("slowquery:create", func(db *gorm.DB) {
		var (
			sql  = db.Statement.SQL.String()
			vars = db.Statement.Vars
		)

		// 追加sql数据
		AddSql(sql, vars)
	})

	// 注册更新
	_ = db.Callback().Update().After("gorm:update").Register("slowquery:update", func(db *gorm.DB) {
		var (
			sql  = db.Statement.SQL.String()
			vars = db.Statement.Vars
		)

		// 追加sql数据
		AddSql(sql, vars)
	})

	// 注册删除
	_ = db.Callback().Delete().After("gorm:delete").Register("slowquery:delete", func(db *gorm.DB) {
		var (
			sql  = db.Statement.SQL.String()
			vars = db.Statement.Vars
		)

		// 追加sql数据
		AddSql(sql, vars)
	})
}

// getSQL 替换 SQL 中的占位符 "?" 为实际值
func getSQL(sql string, vars []interface{}) string {
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

		// 替换第一个 "?" 为实际值
		sql = strings.Replace(sql, "?", formattedValue, 1)
	}

	return sql
}

// AddSql 线程安全地追加sql数据
func AddSql(sql string, vars []interface{}) {
	mu.Lock()
	defer mu.Unlock()

	SqlRes = append(SqlRes, getSQL(sql, vars))
}

// GetAllSql 获取所有sql记录
func GetAllSql() []string {
	mu.Lock()
	defer mu.Unlock()

	// 复制一份SqlRes返回,避免外部修改
	sqlRes := make([]string, len(SqlRes))
	copy(sqlRes, SqlRes)

	return sqlRes
}
