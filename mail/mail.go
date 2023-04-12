package email

import (
	"fmt"
	"net/smtp"

	eemail "github.com/jordan-wright/email"
)

type Config struct {
	Subject string //主题
	Owner   string //发送人

	Email    string //邮箱 eg: 1213383851@qq.com
	Password string //邮箱密码
	Host     string //邮箱域名 eg: smtp.qq.com
	Port     int    //邮箱端口 smtp 465/587
}

type Emailer struct {
	conf *Config
}

func NewEmailer(conf *Config) *Emailer {
	return &Emailer{
		conf: conf,
	}
}

func (er *Emailer) SendSmtpEmail(email, content string) error {
	e := eemail.NewEmail()
	e.Subject = er.conf.Subject
	e.From = fmt.Sprintf("%s <%s>", er.conf.Owner, er.conf.Email)

	e.To = append(e.To, email)
	e.Text = []byte(content)
	return e.Send(fmt.Sprintf("%s:%v", er.conf.Host, er.conf.Port), smtp.PlainAuth("", er.conf.Email, er.conf.Password, er.conf.Host))
}
