package service

import servicemodel "github.com/MailService/pkg/service/models"

type IMailService interface {
	SendMessage(request *servicemodel.MailRequest) error
}
