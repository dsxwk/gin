package config

import (
	"gin/utils/ctx"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

// Logger 包装器
type Logger struct {
	*zap.Logger
}

// NewLogger 创建Logger包装器
func NewLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{zapLogger}
}

// InitLog 初始化日志系统
func InitLog() *Logger {
	// 确保日志目录存在
	logDir := "storage/logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		color.Red("❌  创建日志目录失败:", err)
		os.Exit(1)
	}

	// 动态日志路径
	logPath := filepath.Join(logDir, time.Now().Format("2006-01")+".log")

	// 滚动日志配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    Conf.Log.MaxSize,
		MaxBackups: Conf.Log.MaxBackups,
		MaxAge:     Conf.Log.MaxDay,
		Compress:   true,
	}

	// 编码配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 格式化堆栈输出(多行缩进)
	encoderConfig.StacktraceKey = "stackTrace"

	// 创建encoder,同时输出到文件 + 控制台
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 动态设置日志级别
	level := zap.NewAtomicLevel()
	switch strings.ToLower(Conf.Log.Level) {
	case "debug":
		level.SetLevel(zap.DebugLevel)
	case "warn":
		level.SetLevel(zap.WarnLevel)
	case "error":
		level.SetLevel(zap.ErrorLevel)
	default:
		level.SetLevel(zap.InfoLevel)
	}

	// 创建核心
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberJackLogger), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	// 初始化 Logger
	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel), // 自动为error级别以上日志添加堆栈
	)
	zap.ReplaceGlobals(zapLogger) // 替换全局 zap.L()

	return NewLogger(zapLogger)
}

// GetReqFields 从context获取附加日志字段
func GetReqFields() []zap.Field {
	logger := ctx.GetContext(ctx.KeyLogger)
	if logger == nil {
		return nil
	}

	traceID := logger.GetString(ctx.KeyTraceID)
	clientIP := logger.ClientIP()
	method := logger.Request.Method
	path := logger.Request.URL.Path
	params := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(logger.GetString(ctx.KeyParams), "\r", ""), "\n", ""))

	fields := []zap.Field{
		zap.String("traceId", traceID),
		zap.String("clientIp", clientIP),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("params", params),
	}

	return fields
}

// WithFields 添加日志字段
func (l *Logger) WithFields(fields ...zap.Field) []zap.Field {
	baseFields := GetReqFields()
	if len(baseFields) == 0 {
		return fields // 无上下文时只返回传入字段
	}

	return append(baseFields, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Debug(msg, mergedFields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Info(msg, mergedFields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Warn(msg, mergedFields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Error(msg, mergedFields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Panic(msg, mergedFields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	mergedFields := l.WithFields(fields...)
	if mergedFields == nil {
		mergedFields = fields
	}
	l.Logger.Fatal(msg, mergedFields...)
}

type StackTrace struct{}

func (s StackTrace) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	stack := strings.Split(string(debug.Stack()), "\n")
	return enc.AddArray("stack", zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
		for _, line := range stack {
			line = strings.TrimSpace(line)
			if line != "" {
				arr.AppendString(line)
			}
		}
		return nil
	}))
}
