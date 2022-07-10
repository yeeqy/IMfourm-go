package validators

import (
	"IMfourm-go/pkg/database"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strconv"
	"strings"
	"unicode/utf8"
)

//存放自定义规则和验证器

func init(){
	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户
	govalidator.AddCustomRule("not_exists",func(field string,rule string,message string, value interface{})error{
		rng := strings.Split(strings.TrimPrefix(rule,"not_exists:"),",")
		//第一个参数 表名称
		tableName := rng[0]
		//第二个参数 字段名称
		dbFiled := rng[1]
		//第三个参数 排除ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}
		//用户请求过来的数据
		reqValue := value.(string)

		//拼接sql
		query := database.DB.Table(tableName).Where(dbFiled+" = ?",reqValue)
		//如果有第三个参数，加上sql where过滤
		if len(exceptID) > 0 {
			query.Where("id != ?",exceptID)
		}

		//查询数据库
		var count int64
		query.Count(&count)

		//验证不通过，数据库能找到对应数据
		if count!=0{
			//如果有自定义错误消息
			if message!=""{
				return errors.New(message)
			}
			//默认错误消息
			return fmt.Errorf("%v 已被占用",reqValue)
		}
		//验证通过
		return nil
	})

	//max_cn:8 中文长度设定不超过8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l,_ := strconv.Atoi(strings.TrimPrefix(rule,"max_cn:"))
		if valLength > l {
			if message != ""{
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字",l)
		}
		return nil
	})
	//min_cn:2 中文长度设定不小于2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l,_ := strconv.Atoi(strings.TrimPrefix(rule,"min_cn:"))
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字",l)
		}
		return nil
	})
}
