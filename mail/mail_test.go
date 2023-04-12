package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_emailer_SendSmtpEmail(t *testing.T) {
	emailer := NewEmailer(&Config{
		Subject: "星火链网统一认证平台-注册验证码",
		Owner:   "黑云科技",

		Email:    "1213383851@qq.com",
		Password: "jduodyunajwkfhci",
		Host:     "smtp.qq.com",
		Port:     587,
	})

	assert.NoError(t, emailer.SendSmtpEmail("13545826685@163.com", "hello world"))
}
