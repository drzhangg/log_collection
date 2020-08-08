package server

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	conn, err := amqp.Dial("amqp://admin:admin@47.103.9.218:5672/")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	q, err := ch.QueueDeclare(
		"testqueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         []byte("this is golang to rabbitmq"),
	})
	if err != nil {
		log.Panic(err)
	}

}

func TestConsumer(t *testing.T) {
	conn, err := amqp.Dial("amqp://admin:admin@47.103.9.218:5672/")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch := make(chan bool)
	for i := 0; i < 4; i++ {
		go func(rt int) {
			ch, err := conn.Channel()
			if err != nil {
				log.Panic(err)
			}
			defer ch.Close()

			q, err := ch.QueueDeclare("testqueue", true, false, false, false, nil)
			if err != nil {
				log.Panic(err)
			}
			defer ch.Close()

			msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
			if err != nil {
				log.Panic(err)
			}

			for msg := range msgs {
				log.Printf("协程 %d    Received a message: %s", rt, msg.Body)
				time.Sleep(5 * time.Second)
				msg.Ack(true)
			}
		}(i)
	}
	<-ch
}
