package mq

import (
	"fmt"
	"github.com/MailService/pkg/handlers"
	service "github.com/MailService/pkg/service/models"
	"github.com/MailService/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

type MqClient struct {
	mqhost string
	mqport string
	mquser string
	mqpwd  string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func New() {
	handler, err := handlers.New()
	if err != nil {
		panic(err)
	}

	var client = MqClient{
		mqhost: os.Getenv("MQ_HOST"),
		mqport: os.Getenv("MQ_PORT"),
		mquser: os.Getenv("MQ_USERNAME"),
		mqpwd:  os.Getenv("MQ_PASSWORD"),
	}

	fmt.Println(client)

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", client.mquser, client.mqpwd, client.mqhost, client.mqport))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			failOnError(err, "Failed to close connection")
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Failed to close channel")
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"mail_queue",
		false,
		true,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// setup Fair dispatch
	//err = ch.Qos(
	//	1,     // prefetch count
	//	0,     // prefetch size
	//	false, // global
	//)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			req, err := utils.Byte2Json[service.MailRequest](d.Body)
			if err != nil {
				failOnError(err, "Failed to convert request body")
			}
			if err := handler.SendHTMLMessage(&req); err != nil {
				failOnError(err, "Failed to convert request body")

				if err := d.Reject(true); err != nil {
					failOnError(err, "Failed to reject request")
					return
				}
			} else {
				if err := d.Ack(false); err != nil {
					failOnError(err, "Failed to send ack request")
					return
				}

			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
