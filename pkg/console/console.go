// Package console 命令行辅助方法
package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

//内部使用 设置高亮颜色
func colorOut(message,color string){
	fmt.Fprintln(os.Stdout,ansi.Color(message,color))
}

//打印一条成功消息，绿色输出
func Success(msg string){
	colorOut(msg,"green")
}
//打印一条报错消息，红色输出
func Error(msg string){
	colorOut(msg,"red")
}
//打印一条提示消息，黄色输出
func Warning(msg string){
	colorOut(msg,"yellow")
}
//打印一条报错消息，并退出os.Exit(1)
func Exit(msg string){
	Error(msg)
	os.Exit(1)
}
//语法糖，自带err != nil 判断
func ExitIf(err error){
	if err!=nil{
		Exit(err.Error())
	}
}
