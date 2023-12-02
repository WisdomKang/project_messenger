package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// singleton?
var MController *MessageController

func init() {
	log.Println("Initialize Message Controllor")
	log.Println("Connecting AMQP Broker Server...")
	conn, err := amqp.Dial("amqp://user:password@localhost:5672")

	MController = &MessageController{
		conn: conn,
	}

	if err != nil {
		log.Fatalf("Fail to connected AMQP BrokerServer : %s", err.Error())
	}

	log.Println("Connected AMQP Broker Server.")
}

type MessageController struct {
	conn *amqp.Connection
}

func (MController *MessageController) InitController() {

	err := MController.initMainHubExchange()

	if err != nil {
		log.Fatal(err)
	}

}

func (MController *MessageController) initMainHubExchange() (err error) {
	channel, err := MController.conn.Channel()

	if err != nil {
		return
	}

	err = channel.ExchangeDeclare(
		MAIN_HUB_EXCHANGE_NAME,
		amqp.ExchangeHeaders,
		true,
		false,
		false,
		false,
		nil,
	)

	return
}
