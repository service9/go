package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

//发送邮件
//go get github.com/jordan-wright/email
func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "你是猪 <ljj001210@163.com>"
	e.To = []string{"6988314@qq.com"}
	e.Subject = "验证码发送测试" //主题
	e.HTML = []byte("您的验证码是:<b>123456</b>")
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "ljj001210@163.com", "AGWTEXDWASMZDRBI", "smtp.163.com"))
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "ljj001210@163.com", "AGWTEXDWASMZDRBI", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	//返回EOF,关闭SSL重试
	if err != nil {
		t.Fatal(err)
	}
}
