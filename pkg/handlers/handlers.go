package handlers

import (
	service "github.com/MailService/pkg/service"
	servicemodel "github.com/MailService/pkg/service/models"
)

type MailHandler struct {
	service service.IMailService
}

func New() (*MailHandler, error) {
	ser, err := service.New()
	if err != nil {
		return nil, err
	}

	return &MailHandler{service: ser}, nil
}

func (handler *MailHandler) SendHTMLMessage(request *servicemodel.MailRequest) error {
	err := handler.service.SendMessage(request)
	return err
}
