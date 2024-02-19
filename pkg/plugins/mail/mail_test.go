package mail

import (
	"log"
	"testing"
)

func TestSendMail(t *testing.T) {
	o := &Options{
		MailHost: "smtp.qq.com",
		MailPort: 465,
		MailUser: "aaaa@qq.com",
		MailPass: "aaaaa",
	}
	err := sendMail(o)
	if err != nil {
		log.Println(err)
	}
	log.Println("邮件发送完成")
}
