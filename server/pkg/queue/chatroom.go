package queue

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AmqpBroker struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"amqp_broker"`
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("system error : %s\n", err)
	}
}

var amqpConnection *amqp091.Connection

func init() {
	// Read configuration properties on yaml file
	filePath, err := os.Getwd()
	handleError(err)

	configData, err := os.ReadFile(strings.Join([]string{filePath, SETTING_FILE_NAME}, ""))

	handleError(err)

	var config Config

	err = yaml.Unmarshal(configData, &config)

	handleError(err)

	brokerUrl := strings.Join(
		[]string{
			"amqp://",
			config.AmqpBroker.UserName, ":",
			config.AmqpBroker.Password, "@",
			config.AmqpBroker.Host, ":",
			strconv.Itoa((config.AmqpBroker.Port))}, "")

	fmt.Println(brokerUrl)
	amqpConnection, err = amqp091.Dial(brokerUrl)
	handleError(err)

	err = initExchange()
	handleError(err)
}

func initExchange() error {
	channel, err := amqpConnection.Channel()

	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(
		MAIN_EXCHANGE,
		amqp091.ExchangeHeaders,
		false,
		false,
		false,
		false,
		nil,
	)

	return err
}

func generateID() string {
	uuidStr := uuid.New()
	return uuidStr.String()
}

func CreateChatRoom() {

}

func createChatRoomQeueue() *ChatRoomConsumer {
	roomId := generateID()

	channel, err := amqpConnection.Channel()

	handleError(err)

	queue, err := channel.QueueDeclare(
		roomId,
		false,
		false,
		false,
		false,
		nil,
	)

	handleError(err)

	err = channel.QueueBind(
		queue.Name,
		"",
		"MAIN_EXCHANGE",
		false,
		amqp091.Table{
			"x-match":   "any",
			"roomId":    queue.Name,
			"queueType": QUEUE_TPYE_CHAT,
		},
	)

	handleError(err)

	queueConsumer, err := channel.Consume(queue.Name, queue.Name, false, false, false, false, nil)

	handleError(err)

	return &ChatRoomConsumer{
		RoomId:    roomId,
		RoomQueue: queueConsumer,
	}

}

type ChatRoomConsumer struct {
	RoomId    string
	UserList  []User
	RoomQueue <-chan amqp091.Delivery
}

type User struct {
	UserId    string
	UserQueue <-chan amqp091.Delivery
}
