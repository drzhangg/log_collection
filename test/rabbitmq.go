package test

import "github.com/streadway/amqp"

const (
	addr      = "amqp://admin:admin@47.103.9.218:5672/"
	queuename = "test_queue"
)

func main() {

}

func publicErr(err error) {
	if err !=nil{
		panic(err)
	}
}

func Producer() {
	conn, err := amqp.Dial(addr)
	publicErr(err)
	defer conn.Close()


}
