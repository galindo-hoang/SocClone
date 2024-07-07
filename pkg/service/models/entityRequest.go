package service

import mail "github.com/xhit/go-simple-mail/v2"

type MailRequest struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	ContentType mail.ContentType
}
