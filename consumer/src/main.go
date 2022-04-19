package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type Customer struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"nome"`
	Address   string    `json:"endereco"`
	CreatedAt string    `json:"cadastrado_em"`
	UpdatedAt string    `json:"atualizado_em"`
}

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
	log.Printf("[QueueDeclare] %s", name)
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

func Consume(queue *amqp.Queue, callback func(amqp.Delivery)) error {
	ch := GetChannel()
	_ = ch.Qos(
		100,
		0,
		false,
	)
	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		log.Printf("[Consume] message from %s", queue.Name)
		go callback(msg)
		log.Printf("[Consume] done!")
	}
	return nil
}

func ConsumerCreateCustomer(msg amqp.Delivery) {
	var customer Customer
	err := json.Unmarshal(msg.Body, &customer)
	if err != nil {
		panic(err)
	}

	folder := os.Getenv("NOVOS_CLIENTES")
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
	}
	path := fmt.Sprintf("%s/%s.json", folder, customer.Uuid.String())
	log.Printf("[ConsumerCreateCustomer] creating file %s", path)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	log.Printf("[ConsumerCreateCustomer] done!")
	defer file.Close()

	fileData, _ := json.Marshal(customer)
	file.Write(fileData)
	if err != nil {
		panic(err)
	}
}

func main() {
	queueName := os.Getenv("RABBITMQ_CUSTOMER_CREATE_QUEUE")
	queue := QueueDeclare(queueName)

	forever := make(chan bool)
	go Consume(queue, ConsumerCreateCustomer)
	<-forever
}
