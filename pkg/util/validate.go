package util

import (
	"regexp"
)

var (
	nameRegexp  = regexp.MustCompile(`^[a-z][a-z0-9-]{0,39}$`)
)

// 小写字母开头，后面可以跟小写字母、数字或破折号-，最多允许 40 个字符的长度。
func ValidateName(name string) bool {
	return nameRegexp.MatchString(name)
}