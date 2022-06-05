package main

import (
	"IMfourm-go/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("hello world")

	router := gin.New()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":3000")
	if err!=nil{
		//错误处理
		fmt.Println(err.Error())
	}
}
