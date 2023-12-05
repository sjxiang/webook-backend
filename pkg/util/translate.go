package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// 翻译
func Translate(err error) string {
	// errs 是一个错误切片，保存验证失败的字段和错误信息
	errs, ok := err.(validator.ValidationErrors)
	
	// 其他错误直接返回
	if !ok {	
		return err.Error()
	}

	// 字段校验错误
	var invalidArgs []string
	for _, err := range errs {
		invalidArgs = append(invalidArgs, 
			fmt.Sprintf("字段 %s 值 %s，不满足条件 %s=%s", 
				err.Field(), 
				err.Value().(string), 
				err.Tag(), 
				err.Param(),
			))
	}

	return strings.Join(invalidArgs, "\t")
}

// 参数绑定失败，类型不匹配 json.Unmarshal
// 参数校验失败，