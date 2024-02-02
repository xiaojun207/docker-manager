package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"strings"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string // 发件人
	MailPass string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func sendMail(o *Options) {

	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", o.MailUser)
	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)
	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)
	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	log.Println(mailArrTo, "邮件发送完成")
}
