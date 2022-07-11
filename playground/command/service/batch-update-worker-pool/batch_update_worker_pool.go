package batch_update_worker_pool

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

var (
	testchan chan amqp.Delivery
)

func Consumer() {
	testchan = make(chan amqp.Delivery, 3)
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != err {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.Qos(5, 0, false)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume("testqueue", "testkey",
		false,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			log.Print(4)
			select {
			case t := <-testchan:
				log.Printf("In %d last consume a message: %s\n", 0, t.Body)
				err = t.Ack(false)

				if err != nil {
					log.Print(err)
				}
				log.Printf("Done")
			}
		}
	}()

	go func() {
		for {
			log.Print(1)
			select {
			case msg := <-msgs:
				log.Printf("In %d consume a message: %s\n", 0, msg.Body)
				testchan <- msg
				log.Print(len(testchan))
				log.Printf("Done")
			}
		}
	}()

	time.Sleep(100000000000)
}
