package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	minSize, digit, special, lower, upper := ValidatePassword("hasicoghwif*4YY")
	if !minSize || !digit || !special || !lower || !upper {
		t.Log("无效密码")
		msg := "密码："
		var errs []string
		if !minSize {
			errs = append(errs, "最少 8 个字符")
		}
		if !digit {
			errs = append(errs, "至少要有 1 个数字")
		}
		if !special {
			errs = append(errs, "至少要有 1 个特殊字符")
		}
		if !lower {
			errs = append(errs, "至少要有 1 个小写字母")
		}
		if !upper {
			errs = append(errs, "至少要有 1 个大写字母")
		}
		
		t.Log(msg + strings.Join(errs, "，"))
	}
}



func TestEncrypt(t *testing.T) {
	plainText := "vbwivbvo;079r23"
	
	hashedPassword, err := HashPassword(plainText)
	if err != nil {
		t.Fatal(err)
	}

	err = CheckPassword(plainText, hashedPassword)
	require.NoError(t, err)
}
