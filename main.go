package main

// https://www.loginradius.com/blog/engineering/sending-emails-with-golang/

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "info.zhn.2023@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "tien.huynh@newhorizon-tech.vn")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Dear New Horizon's member! \r\nThis is internal mail to test new feature \r\n Please don't reply! \r\n Have a nice day! \r\n Thanks You!")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "info.zhn.2023@gmail.com", "kbmuvmkowccrmecg")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Send mail failed", err)
		panic(err)
	}

	fmt.Printf("Send mail success")
}
