package logger

import (
	"IMfourm-go/pkg/helpers"
	"context"
	"errors"
	"go.uber.org/zap"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

//操作对象，实现gormlogger.Interface
type GormLogger struct {
	ZapLogger *zap.Logger
	SlowThreshold time.Duration
}


// NewGormLogger 实例化一个 GormLogger 对象
//示例：
//     DB, err := gorm.Open(dbConfig, &gorm.Config{
//         Logger: logger.NewGormLogger(),
//     })
func NewGormLogger() GormLogger{
	return GormLogger{
		ZapLogger: Logger,
		SlowThreshold: 200 * time.Millisecond,
	}
}

//实现gormlogger.Interface的LogMode方法
func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface{
	return GormLogger{
		ZapLogger: l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

func (l GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Debugf(s,i...)
}

func (l GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Warnf(s,i...)
}

func (l GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Errorf(s,i...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//获取运行时间
	elapsed := time.Since(begin)
	//获取sql请求和返回条数
	sql,rows := fc()
	//通用字段
	logFields := []zap.Field{
		zap.String("sql",sql),
		zap.String("time",helpers.MicrosecondsStr(elapsed)),
		zap.Int64("rows",rows),
	}
	//gorm错误
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			l.logger().Warn("database errRecordNotFound",logFields...)
		}else {
			logFields = append(logFields,zap.Error(err))
			l.logger().Error("database err",logFields...)
		}
	}
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold{
		l.logger().Warn("database slow log",logFields...)
	}
	//记录所有sql请求
	l.logger().Debug("database query",logFields...)
}

//logger内用的辅助方法
//确保zap内置信息caller的准确性
func (l GormLogger) logger() *zap.Logger {

	// 跳过 gorm 内置的调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
