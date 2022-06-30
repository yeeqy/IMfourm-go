// Package str 字符串辅助方法
package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 转为复数
func Plural(word string) string{
	return pluralize.NewClient().Plural(word)
}

// Singular 转为单数
func Singular(word string) string{
	return pluralize.NewClient().Singular(word)
}

//Snake -> snake_case, eg: TopicComment -> topic_comment
func Snake(s string) string  {
	return strcase.ToSnake(s)
}

//Camel -> CamelCase, eg: topic_comment -> TopicComment
func Camel(s string) string{
	return strcase.ToCamel(s)
}
//LowerCamel -> lowerCamelCase eg: TopicComment -> topicComment
func LowerCamel(s string) string{
	return strcase.ToLowerCamel(s)
}
