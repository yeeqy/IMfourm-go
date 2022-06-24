package verifycode
//用于将验证码模块与特定的存储服务器区分开来
//方便于未来使用其他服务器时进行切换

type Store interface {
	//保存验证码
	Set(id string,value string) bool
	//获取验证码
	Get(id string,clear bool) string
	//检查验证码
	Verify(id,answer string,clear bool) bool
}
