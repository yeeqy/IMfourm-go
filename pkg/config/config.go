package config
// 负责配置信息

import (
	"IMfourm-go/pkg/helpers"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"os"
)
// viper库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init(){
	//1. 初始化Viper库
	viper = viperlib.New()

	//2. 配置类型
	viper.SetConfigType("env")

	//3. 环境变量配置文件查找的路径，相对于main.go
	viper.AddConfigPath(".")

	//4. 设置环境变量前缀，用以区分Go的系统环境变量
	viper.SetEnvPrefix("appenv")

	//5. 读取环境变量（支持flag）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及config信息的加载
func InitConfig(env string){
	//1. 加载环境变量
	loadEnv(env)

	//2. 注册配置信息
	loadConfig()
}

func loadConfig(){
	for name, fn := range ConfigFuncs{
		viper.Set(name,fn())
	}
}
func loadEnv(envSuffix string){
	// 默认加载.env文件
	envPath := ".env"
	if len(envSuffix)>0{
		filepath := ".env." + envSuffix
		if _,err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}

	//加载env
	viper.SetConfigName(envPath)
	if err:= viper.ReadInConfig();err!= nil{
		panic(err)
	}
	//监控 .env文件，变更时重新加载
	viper.WatchConfig()
}
// Env读取环境变量，支持默认值
func Env(envName string, defalutValue ...interface{}) interface{}{
	if len(defalutValue) > 0 {
		return internalGet(envName,defalutValue[0])
	}
	return internalGet(envName)
}
// 新增配置项
func Add(name string,configFn ConfigFunc){
	ConfigFuncs[name] = configFn
}

// 获取配置项
func Get(path string,defaultValue ...interface{}) string{
	return GetString(path,defaultValue...)
}

func internalGet(path string,defaultValue ...interface{}) interface{}{
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)){
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func GetString(path string,defaultValue ...interface{})string{
	return cast.ToString(internalGet(path,defaultValue...))
}
func GetInt(path string,defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path,defaultValue...))
}
func GetFloat64(path string,defaultValue ...interface{})float64{
	return cast.ToFloat64(internalGet(path,defaultValue...))
}
func GetInt64(path string,defaultValue ...interface{})int64{
	return cast.ToInt64(internalGet(path,defaultValue...))
}
func GetUint(path string,defaultValue ...interface{})uint{
	return cast.ToUint(internalGet(path,defaultValue...))
}
func GetBool(path string,defaultValue ...interface{})bool{
	return cast.ToBool(internalGet(path,defaultValue...))
}
//获取数据结构
func GetStringMapString(path string) map[string]string{
	return viper.GetStringMapString(path)
}