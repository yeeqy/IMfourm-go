package make

import (
	"IMfourm-go/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use: "apicontroller",
	Short: "Create api controller, example: make apicontroller v1/user",
	Run: runMakeAPIController,
	Args: cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string){
	//处理参数，要求附带API版本（v1、v2）
	array := strings.Split(args[0],"/")
	if len(array) != 2{
		console.Exit("api controller name format: v1/user")
	}

	//apiVersion 用来拼接目标路径
	//name 用来生成cmd.Model实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	//组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go",apiVersion,model.TableName)

	//基于模板创建文件
	createFileFromStub(filePath,"apicontroller",model)

}