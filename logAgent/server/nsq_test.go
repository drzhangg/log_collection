package server

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
	"time"
)

func TestNsqProducer(t *testing.T) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("47.103.9.218:4150", config)
	if err != nil {
		log.Panic(err)
	}
	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("num-%d:", i)
		log.Println("pub:", msg)
		err = p.Publish("testTopic", []byte(msg))
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(time.Second * 1)
	}

	p.Stop()
}

func TestNsqCustomer(t *testing.T) {

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("testTopic", "ch", config)
	if err != nil {
		log.Panic(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body))
		time.Sleep(2 * time.Second)
		return nil
	}))

	err = consumer.ConnectToNSQD("47.103.9.218:4150")
	if err !=nil {
		log.Panic("Could not connect")
	}
	time.Sleep(3600*time.Second)
}
