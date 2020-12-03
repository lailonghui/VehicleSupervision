package logger

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// gorm源码目录
var gormSourceDir string

// gorm包地址
var gormPackageName string = "gorm.io"

// 应用源码目录
var appSourceDir string

// 应用源码目录长度
var appSourceDirLength int

// 当前文件最上层包名
var currentFileTopPackage string = "pkg"

func init() {
	_, file, _, _ := runtime.Caller(0)
	appSourceDir = file[0:strings.Index(file, currentFileTopPackage)]
	appSourceDirLength = len(appSourceDir)
}

type GormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.LogLevel = level
	return g
}

func (g GormLogger) Info(ctx context.Context, s string, i ...interface{}) {

	if g.LogLevel >= logger.Info {
		GLogger.Info(g.warpMsg(s), zap.Any("infos", i))
	}
}

func (g GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		GLogger.Info(g.warpMsg(s), zap.Any("infos", i))
	}
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		GLogger.Info(g.warpMsg(s), zap.Any("infos", i))
	}
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if g.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && g.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.String("rows", "-"), zap.Error(err))
			} else {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.Int64("rows", rows), zap.Error(err))
			}
		case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= logger.Warn:
			sql, rows := fc()
			if rows == -1 {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.String("rows", "-"), zap.Error(err))
			} else {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.Int64("rows", rows), zap.Error(err))
			}
		case g.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.String("rows", "-"), zap.Error(err))
			} else {
				GLogger.Debug(g.warpMsg(sql), zap.String("times", elapsed.String()), zap.Int64("rows", rows), zap.Error(err))
			}
		}
	}
}

// 获取打印日志当前应用的文件名和行号
func FileWithLineNum() string {
	for i := 3; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if gormSourceDir == "" {
			if strings.Contains(file, gormPackageName) {
				gormSourceDir = file[0 : strings.Index(file, gormPackageName)+len(gormPackageName)]
			} else {
				continue
			}
		}
		if ok && (!strings.HasPrefix(file, gormSourceDir) || strings.HasSuffix(file, "_test.go")) {
			return file[appSourceDirLength:len(file)] + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}

func (g GormLogger) warpMsg(msg string) string {
	return FileWithLineNum() + "  " + msg
}
