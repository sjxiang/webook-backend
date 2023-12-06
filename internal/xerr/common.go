package xerr

import (
	"errors"
)

var (
	UserDuplicateEmail    = errors.New("邮箱冲突")
	InvalidUserOrPassword = errors.New("账号/邮箱或密码不对")
	UserNotFound          = errors.New("无用户记录")
)