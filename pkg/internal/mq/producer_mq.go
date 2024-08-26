package mq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AuthService/pkg/models/rbmq"
	"github.com/AuthService/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessageMail(mail rbmq.MailRequest) error {
	var (
		UserName = utils.GetValue("MQ_USERNAME")
		password = utils.GetValue("MQ_PASSWORD")
		host     = utils.GetValue("MQ_HOST")
		port     = utils.GetValue("MQ_PORT")
	)
	connect, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", UserName, password, host, port))
	if err != nil {
		return err
	}

	if connect == nil {
		panic("please connect to broker first!!")
	}
	ch, err := connect.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"mail_queue", //name
		false,        // durable
		true,         // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parsedBody, err := utils.JSON2Byte(mail)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        parsedBody,
	})
	if err != nil {
		return err
	}

	log.Printf("Sent Message Successfully!!")
	return nil
}
