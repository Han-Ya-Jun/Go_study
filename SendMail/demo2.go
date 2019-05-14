package main

import "gopkg.in/gomail.v2"

/*
* @Author:15815
* @Date:2019/5/8 17:20
* @Name:main
* @Function:
 */
func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1581532052@qq.com")
	m.SetHeader("To", "1581532052@qq.com", "hanyajun5876@163.com")
	//m.SetAddressHeader("Cc", "1581532052@qq.com", "韩亚军")
	m.SetHeader("Cc", "1581532052@qq.com", "hanyajun5876@163.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "123")
	d := gomail.NewDialer("smtp.qq.com", 25, "1581532052@qq.com", "vwhfadvqlqwlgcdh")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
