package models

import mail "github.com/xhit/go-simple-mail/v2"

type MailEntity struct {
	Domain      string `json:"domain"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Encryption  string `json:"encryption"`
	FromAddress string `json:"from"`
	FromName    string `json:"fromName"`
}

type Message struct {
	From               string           `json:"from"`
	FromName           string           `json:"fromName"`
	To                 string           `json:"to"`
	Subject            string           `json:"subject"`
	Attachments        []string         `json:"attachments"`
	Message            string           `json:"message"`
	AlternativeMessage string           `json:"alternative_message"`
	ContentType        mail.ContentType `json:"content_type"`
}
