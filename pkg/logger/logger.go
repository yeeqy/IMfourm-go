package logger

import (
	"IMfourm-go/pkg/app"
	"encoding/json"
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

//添加一些日志辅助方法，既方便调用，又对zap进行封装
//对 zap 的封装，除了有方便调用的好处外，如果因为特殊情况，我们不得不弃用 zap 而使用其他日志库，封装会极大减少我们重构的工作量。

// Dump 调试专用，会在终端打印出warning信息
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
//         logger.Dump(user.User{Name:"test"})
//         logger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value interface{}, msg ...string){
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump",zap.String(msg[0],valueString))
	} else {
		Logger.Warn("Dump",zap.String("data",valueString))
	}
}
//当 err != nil 时 记录error等级的日志
func LogIf(err error){
	if err!=nil{
		Logger.Error("Error Occurred:",zap.Error(err))
	}
}

func LogWarnIf(err error){
	if err != nil {
		Logger.Error("Error Occurred:",zap.Error(err))
	}
}
func LogInfoIf(err error){
	if err != nil {
		Logger.Error("Error Occurred:",zap.Error(err))
	}
}

//调试日志，详尽的程序日志
func Debug (moduleName string,fields ...zap.Field){
	Logger.Debug(moduleName,fields...)
}
func Info (moduleName string,fields ...zap.Field){
	Logger.Info(moduleName,fields...)
}
func Warn (moduleName string,fields ...zap.Field){
	Logger.Warn(moduleName,fields...)
}
func Error (moduleName string,fields ...zap.Field){
	Logger.Error(moduleName,fields...)
}
func Fatal (moduleName string,fields ...zap.Field){
	Logger.Fatal(moduleName,fields...)
}

//记录一条字符串类型的debug日志
func DebugString(modulename,name,msg string){
	Logger.Debug(modulename,zap.String(name,msg))
}
func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//         logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}