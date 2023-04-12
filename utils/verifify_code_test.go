package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenVerifyCode(t *testing.T) {
	var (
		code1 = GenVerifyCode()
		code2 = GenVerifyCode()
	)
	t.Logf("第一次生成的验证码为:%s, 第二次生成的验证码为:%s", code1, code2)

	assert.NotEqual(t, code1, code2)
}
