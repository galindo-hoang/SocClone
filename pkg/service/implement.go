package service

import (
	"bytes"
	"github.com/MailService/pkg/repositories"
	"github.com/MailService/pkg/repositories/models"
	service "github.com/MailService/pkg/service/models"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
)

func New() (*MailService, error) {
	repo, err := repositories.New()
	if err != nil {
		return nil, err
	}

	return &MailService{repository: repo}, nil
}

type MailService struct {
	repository repositories.IMailRepository
}

func (s *MailService) SendMessage(request *service.MailRequest) error {
	var parsedMsg string
	request.DataMap = map[string]any{
		"message": request.Data,
	}
	if request.ContentType == mail.TextHTML {
		parsedMessage, err := s.buildHTMLMessage(request)
		if err != nil {
			return err
		}
		parsedMsg = parsedMessage
	} else {
		parsedMessage, err := s.buildPlainTextMessage(request)
		if err != nil {
			return err
		}
		parsedMsg = parsedMessage
	}

	var msg = &models.Message{
		From:        request.From,
		FromName:    request.FromName,
		To:          request.To,
		Subject:     request.Subject,
		Attachments: request.Attachments,
		Message:     parsedMsg,
		ContentType: mail.TextHTML,
	}

	err := s.repository.SendMail(msg)

	return err
}

func (s *MailService) buildHTMLMessage(msg *service.MailRequest) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"
	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}
	formattedMsg := tpl.String()
	formattedMsg, err = s.inlineCSS(formattedMsg)
	if err != nil {
		return "", err
	}
	return formattedMsg, nil
}

func (s *MailService) inlineCSS(str string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(str, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}
	return html, nil
}

func (s *MailService) buildPlainTextMessage(msg *service.MailRequest) (string, error) {
	templateToRender := "./templates/mail.plain.gohtml"
	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()
	return plainMessage, nil
}
