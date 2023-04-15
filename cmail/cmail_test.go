package cmail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_emailer_SendSmtpEmail(t *testing.T) {
	emailer := NewMailer(&Config{
		Subject: "星火链网统一认证平台-注册验证码",
		Owner:   "黑云科技",

		Email:    "1213383851@qq.com",
		Password: "ezfzdvivqenkhjbj",
		Host:     "smtp.qq.com",
		Port:     465,
	})

	assert.NoError(t, emailer.SendSmtpMail("1213383851@qq.com", "hello world"))
}
