package gomail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	var mailConf MailConfForm
	mailConf.EmailHost = "xx"
	mailConf.EmailUser = "xx" // 其他全填写完整,此处省略
	var conf Config
	conf.Config = mailConf
	conf.MailTo = append(conf.MailTo, mailConf.EmailTest)
	conf.Subject = mailConf.EmailTestTitle
	SendMail(conf)
}
