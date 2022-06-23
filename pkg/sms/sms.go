package sms

import (
	"IMfourm-go/pkg/config"
	"sync"
)

//短信结构体
type Message struct {
	Template string
	Data map[string]string

	Content string
}

//SMS是我们发送短信的操作类
type SMS struct {
	Driver Driver
}
//单例模式
var once sync.Once

//内部使用的SMS对象
var internalSMS *SMS

//单例模式获取
func NewSMS() *SMS{
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone,message,config.GetStringMapString("sms.aliyun"))
}