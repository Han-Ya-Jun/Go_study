package main

import (
	"github.com/veqryn/go-email/email"
	"net/smtp"
)

/*
* @Author:hanyajun
* @Date:2019/5/8 16:58
* @Name:SendMail
* @Function:
 */
func main() {
	text := "123"
	// gopherReader = io.Reader with the content of an image (as an example)
	// docBytes = []byte with the content of a pdf (as an example)

	h := email.Header{}
	h.SetFrom("charlie@gmail.com")
	h.SetTo("john@outlook.com", "laura@yahoo.com")
	h.SetCc("sarah@icloud.com", "chris@gmail.com")
	h.SetSubject("Hello from go-email")

	//imagePart, err := email.NewPartAttachment(gopherReader, "gopher.png")
	//pdfPart := email.NewPartAttachmentFromBytes(docBytes, "documentation.pdf"
	msg := email.NewMessage(h, text, "")
	msg.Send("smtp.gmail.com:587", smtp.PlainAuth("", "username@gmail.com", "", "smtp.gmail.com"))
}
