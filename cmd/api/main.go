package main

import (
	"fmt"
	"github.com/MailService/pkg/handlers"
	service "github.com/MailService/pkg/service/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func main() {
	h, err := handlers.New()
	if err != nil {
		panic(err)
	}
	err = h.SendHTMLMessage(&service.MailRequest{
		To:          "hmhuy191101@gmail.com",
		From:        "hmhuy19110111@gmail.com",
		Data:        "requestPayload.Message",
		Subject:     "requestPayload.Subject",
		ContentType: mail.TextPlain,
	})
	fmt.Println("mail sent successfully")
}
