package config

import (
	"context"
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

var (
	ZapLogger *Logger
)

// Logger 包装器
type Logger struct {
	*zap.Logger
}

// NewLogger 创建Logger包装器
func NewLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{zapLogger}
}

func init() {
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

	ZapLogger = NewLogger(zapLogger)
}

func (l *Logger) WithContext(c context.Context) *zap.Logger {
	if c == nil {
		// 没有请求上下文, 返回普通logger
		return l.Logger
	}

	traceId := ctx.GetValue(c, ctx.KeyTraceId).(string)
	// 动态计算 ms
	var ms any
	if startTime, ok := ctx.GetValue(c, ctx.KeyStartTime).(time.Time); ok {
		ms = time.Since(startTime).Milliseconds()
	} else {
		ms = ctx.GetValue(c, ctx.KeyMs)
	}

	return l.Logger.With(
		zap.Any("traceId", traceId),
		zap.Any("ip", ctx.GetValue(c, ctx.KeyIp)),
		zap.Any("path", ctx.GetValue(c, ctx.KeyPath)),
		zap.Any("method", ctx.GetValue(c, ctx.KeyMethod)),
		zap.Any("params", ctx.GetValue(c, ctx.KeyParams)),
		zap.Any("ms", ms),
		zap.Any("debugger", ctx.GetDebugger()),
	)
}

func (l *Logger) Debug(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Debug(msg, fields...)
}

func (l *Logger) Info(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Info(msg, fields...)
}

func (l *Logger) Warn(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Warn(msg, fields...)
}

func (l *Logger) Error(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Error(msg, fields...)
}

func (l *Logger) Panic(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Panic(msg, fields...)
}

func (l *Logger) Fatal(c context.Context, msg string, fields ...zap.Field) {
	l.WithContext(c).Fatal(msg, fields...)
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
