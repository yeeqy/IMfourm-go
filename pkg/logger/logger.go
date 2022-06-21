package logger

import (
	"IMfourm-go/pkg/app"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

//全局Logger对象
var Logger *zap.Logger

//日志初始化
func InitLogger(filename string,maxSize, maxBackup, maxAge int, compress bool, logType string, level string){

	//获取日志写入介质
	writeSyncer := getLogWriter(filename,maxSize,maxBackup,maxAge,compress,logType)

	//设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level));err!=nil{
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}
	//初始化 core
	core := zapcore.NewCore(getEncoder(),writeSyncer,logLevel)

	//初始化 logger
	Logger = zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1),zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(Logger)
}

//设置日志存储格式
func getEncoder() zapcore.Encoder{
	//日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "logger",
		CallerKey:           "caller",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "stacktrace",
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalLevelEncoder,
		EncodeTime:          customTimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
	}
	if app.IsLocal(){
		//终端输入的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		//本地设置内置的console解码器（支持stacktrace换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	//线上环境使用JSON编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

//自定义友好的时间格式
func customTimeEncoder(t time.Time,enc zapcore.PrimitiveArrayEncoder){
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
//日志记录介质 使用了两种介质，os.Stdout 和文件
func getLogWriter(filename string,maxSize,maxBackup,maxAge int,compress bool,logType string)zapcore.WriteSyncer{
	//如果配置了按照日期记录日志文件
	if logType == "daily"{
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename,"logs.log",logname)
	}
	//滚动日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}
	//配置输出介质
	if app.IsLocal(){
		//本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(lumberjackLogger))
	}else {
		//生产环境只记录文件
		return zapcore.AddSync(lumberjackLogger)
	}
}