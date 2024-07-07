package repositories

import (
	"github.com/MailService/pkg/repositories/models"
	"github.com/MailService/pkg/utils"
	mail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"time"
)

type MailRepository struct {
	mail *models.MailEntity
}

func New() (IMailRepository, error) {
	var repo MailRepository
	port, err := strconv.Atoi(utils.GetValue("MAIL_PORT"))
	if err != nil {
		return &repo, err
	}

	repo.mail = &models.MailEntity{
		Domain:      utils.GetValue("MAIL_DOMAIN"),
		Host:        utils.GetValue("MAIL_HOST"),
		Port:        port,
		Username:    utils.GetValue("MAIL_USERNAME"),
		Password:    utils.GetValue("MAIL_PASSWORD"),
		Encryption:  utils.GetValue("MAIL_ENCRYPTION"),
		FromName:    utils.GetValue("FROM_NAME"),
		FromAddress: utils.GetValue("FROM_ADDRESS"),
	}

	return &repo, nil
}

func (r *MailRepository) SendMail(msg *models.Message) error {
	server := mail.NewSMTPClient()
	server.Host = r.mail.Host
	server.Port = r.mail.Port
	server.Username = r.mail.Username
	server.Password = r.mail.Password
	server.KeepAlive = false
	server.Encryption = r.getEncryption(r.mail.Encryption)
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject).
		SetBody(msg.ContentType, msg.Message).
		AddAlternative(msg.ContentType, msg.AlternativeMessage)

	if len(msg.Attachments) > 0 {
		for _, attachment := range msg.Attachments {
			email.AddAttachment(attachment)
		}
	}

	if err := email.Send(smtpClient); err != nil {
		return err
	}

	return nil
}

func (r *MailRepository) getEncryption(msg string) mail.Encryption {
	switch msg {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
