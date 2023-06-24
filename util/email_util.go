package util

import (
	"fmt"
	"net/smtp"
)

func SendCode(toEmail, code string) error {

	to := []string{toEmail}
	auth := smtp.PlainAuth("", "943155756@qq.com", "rinxamuyqopkbdib", "smtp.qq.com")
	msg := []byte("From: " + "943155756@qq.com" + "\r\n" +
		"To: " + toEmail + "\r\n" +
		"Subject: 注册验证码\r\n" +
		"\r\n" +
		"亲爱的用户你好，你的智慧医疗平台验证码为：" + code + "，三分钟内有效，若不是您本人的操作请忽略\r\n")

	addr := fmt.Sprintf("%s:%s", "smtp.qq.com", "587")

	return smtp.SendMail(addr, auth, "943155756@qq.com", to, msg)
}
