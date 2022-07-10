package redis

import (
	"IMfourm-go/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Context context.Context
}
//确保全局的redis对象只实例一次
var once sync.Once

var Redis *RedisClient

// NewClient 创建一个新的redis连接
func NewClient(address string,username string,password string,db int) *RedisClient{
	//1. 初始化自定的redisClient实例
	rds := &RedisClient{}
	//2. 使用默认的context
	rds.Context = context.Background()
	//3. 使用redis库里的NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr: address,
		Username: username,
		Password: password,
		DB: db,
	})
	//测试一下连接
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}
func (rds RedisClient) Ping() error{
	_,err := rds.Client.Ping(rds.Context).Result()
	return err
}

// ConnectRedis 连接redis数据库，设置全局的redis对象
func ConnectRedis(address string,username string,password string,db int){
	once.Do(func() {
		Redis = NewClient(address,username,password,db)
	})
}

// Set 存储key对应的value，且设置过期时间
func (rds RedisClient) Set(key string,value interface{},expiration time.Duration)bool{
	if err := rds.Client.Set(rds.Context,key,value,expiration).Err();err!=nil{
		logger.ErrorString("Redis","Set",err.Error())
		return false
	}
	return true
}

func (rds RedisClient) Get(key string) string{
	result,err := rds.Client.Get(rds.Context,key).Result()
	if err != nil {
		if err != redis.Nil{
			logger.ErrorString("Redis","Get",err.Error())
		}
		return ""
	}
	return result
}
func (rds RedisClient) Has(key string) bool  {
	_,err := rds.Client.Get(rds.Context,key).Result()
	if err!=nil {
		if err != redis.Nil{
			logger.ErrorString("Redis","Has",err.Error())
		}
		return false
	}
	return true
}
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context,keys...).Err();err!=nil{
		logger.ErrorString("Redis","del",err.Error())
		return false
	}
	return true
}
func (rds RedisClient) FlushDB() bool{
	if err := rds.Client.FlushDB(rds.Context).Err();err!=nil{
		logger.ErrorString("Redis","FlushDB",err.Error())
		return false
	}
	return true
}
func(rds RedisClient) Increment (param ...interface{}) bool {
	switch len(param) {
	case 1:
		key := param[0].(string)
		if err := rds.Client.Incr(rds.Context,key).Err();err != nil {
			logger.ErrorString("Redis","Increment",err.Error())
			return false
		}
	case 2:
		key := param[0].(string)
		value := param[1].(int64)
		if err := rds.Client.IncrBy(rds.Context,key,value).Err();err!=nil{
			logger.ErrorString("Redis","Increment",err.Error())
			return false
		}
	default:
		logger.ErrorString("Reids","Increment","参数过多")
		return false
	}
	return true
}

func(rds RedisClient) Decrement (param ...interface{}) bool {
	switch len(param) {
	case 1:
		key := param[0].(string)
		if err := rds.Client.Decr(rds.Context,key).Err();err != nil {
			logger.ErrorString("Redis","Decrement",err.Error())
			return false
		}
	case 2:
		key := param[0].(string)
		value := param[1].(int64)
		if err := rds.Client.DecrBy(rds.Context,key,value).Err();err!=nil{
			logger.ErrorString("Redis","Decrement",err.Error())
			return false
		}
	default:
		logger.ErrorString("Reids","Decrement","参数过多")
		return false
	}
	return true
}

