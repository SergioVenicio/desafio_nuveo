package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Connection {
	user := os.Getenv("RABBITMQ_USER")
	pwd := os.Getenv("RABBITMQ_PWD")
	host := os.Getenv("RABBITMQ_HOST")
	mqDns := fmt.Sprintf("amqp://%s:%s@%s:5672/", user, pwd, host)
	conn, err := amqp.Dial(mqDns)
	if err != nil {
		panic(err)
	}

	return conn
}
func GetChannel() *amqp.Channel {
	conn := Connect()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func QueueDeclare(name string) *amqp.Queue {
	ch := GetChannel()
	queue, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return &queue
}

func Publish(queue *amqp.Queue, body []byte) error {
	ch := GetChannel()
	log.Printf("[Publish] message on %s", queue.Name)
	err := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func Init() {
	QueueDeclare(os.Getenv("RABBITMQ_CUSTOMER_CREATE_QUEUE"))
}
