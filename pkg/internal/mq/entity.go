package mq

// content-type 0 TextPlain
// content-type 1 TextHTML

type MailRequest struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	ContentType int
}
