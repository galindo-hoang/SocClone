package repositories

import "github.com/MailService/pkg/repositories/models"

type IMailRepository interface {
	SendMail(entity *models.Message) error
}
