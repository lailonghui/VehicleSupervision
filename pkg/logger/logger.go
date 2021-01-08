package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type LogMode int

const (
	// 日志模式-控制台
	MODE_CONSOLE LogMode = iota
	// 文件
	MODE_FILE
	// 控制台+文件
	MODE_ALL
)

// 配置选项
type ConfigOption struct {
	// 日志保存路径
	Path string
	// 日志等级
	Level string
	// 日志模式
	Mode string
	// 应用名称
	AppName string
}

// 应用日志实例
var AppLogger *zap.Logger

// gorm日志实例
var GLogger *zap.Logger

// gin日志实例
var GinLogger *zap.Logger

func Setup(option *ConfigOption) {
	logPath := option.Path
	logLevel := getLevel(option.Level)
	logMode := getLogMode(option.Mode)
	AppLogger = newLogger(option.AppName, logPath, logLevel, logMode)
	// gorm日志实例-关掉默认的caller打印设置
	GLogger = AppLogger.WithOptions(zap.WithCaller(false))
	// gin日志实例-关掉默认的caller打印设置
	GinLogger = AppLogger.WithOptions(zap.WithCaller(false))

}

// 新建日志
func newLogger(logName, logPath string, logLevel zapcore.Level, logMode LogMode) *zap.Logger {
	// 日志格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志输出配置
	var ws zapcore.WriteSyncer
	if logMode == MODE_CONSOLE {
		ws = zapcore.AddSync(os.Stdout)
	} else if logMode == MODE_FILE {
		logName = logName + ".log"
		filePath := path.Join(logPath, logName)
		file := getFileWriter(filePath)
		ws = zapcore.AddSync(file)
	} else {
		logName = logName + ".log"
		filePath := path.Join(logPath, logName)
		file := getFileWriter(filePath)
		ws = zapcore.AddSync(io.MultiWriter(file, os.Stdout))
	}

	core := zapcore.NewCore(encoder, ws, logLevel)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// 获取文件writer
func getFileWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Fatal(err)
	}
	return hook
}

// 获取日志级别
func getLevel(lv string) zapcore.Level {
	switch lv {
	case "":
		return zapcore.DebugLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

// 获取日志模式
func getLogMode(mode string) LogMode {
	switch mode {
	case "console":
		return MODE_CONSOLE
	case "file":
		return MODE_FILE
	case "all":
		return MODE_ALL
	default:
		return MODE_CONSOLE

	}
}
