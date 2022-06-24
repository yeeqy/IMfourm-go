package verifycode

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/redis"
	"time"
)

//redis作为驱动来实现interface

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix string
}

//实现verifycode.Store interface的set方法

func (s *RedisStore) Set(key string,value string) bool{
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	//本地环境方便调试
	if app.IsLocal(){
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key,value,ExpireTime)
}

func(s *RedisStore) Get(key string,clear bool)(value string){
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}
func (s *RedisStore) Verify(key,answer string,clear bool) bool {
	v := s.Get(key,clear)
	return v == answer
}