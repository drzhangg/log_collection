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
	p, err := nsq.NewProducer("0:4150", config)
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

}
