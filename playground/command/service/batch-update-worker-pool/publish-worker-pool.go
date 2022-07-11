package batch_update_worker_pool

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func Publish() {

	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != err {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	_, err = ch.QueueDeclare(
		"testqueue", //Queue name
		true,        //durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(
		"testexchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.QueueBind("testqueue", "testkey", "testexchange", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= 20; i++ {
		if err = ch.Publish("testexchange", "testkey", true, false, amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "text/plain",
			Body:        []byte("Hello Golang and AMQP(Rabbitmq)!"),
		}); err != nil {
			panic(err)
		}
	}

}
