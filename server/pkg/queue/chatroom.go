package queue

import (
	"log"
	"strings"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChatRoom struct {
	ID       string
	hubQueue *amqp.Queue
	users    []string
}

func CreateChatRoom(users []string) (chatroom *ChatRoom) {
	chatroom = &ChatRoom{
		ID:    createRoomId(),
		users: users,
	}

	err := chatroom.CreateChatRoomQueue(chatroom.ID)
	if err != nil {
		log.Fatalf("Error when create chatroom queue : %s", err)
	}

	go chatroom.ConsumeAndPushMessage()

	return
}

// func isExistQueue(roomId string) bool {
// 	channel , err := MController.conn.Channel()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

func CreateUserQueue(userId string) (err error) {
	channel, err := MController.conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(userId, true, false, false, false, nil)
	return

}

func (chatroom *ChatRoom) CreateChatRoomQueue(roomId string) (err error) {
	channel, err := MController.conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	roomQueue, err := channel.QueueDeclare(combineName(roomId), false, false, false, false, nil)

	if err != nil {
		return
	}

	err = channel.QueueBind(
		roomQueue.Name,
		"",
		MAIN_HUB_EXCHANGE_NAME,
		false,
		amqp.Table{
			"x-match": "all",
			"roomId":  chatroom.ID,
		},
	)

	chatroom.hubQueue = &roomQueue

	return
}

func (chatroom *ChatRoom) ConsumeAndPushMessage() {
	channel, err := MController.conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	msgs, err := channel.Consume(chatroom.hubQueue.Name, chatroom.ID, true, false, false, false, nil)

	if err != nil {
		return
	}

	for msg := range msgs {
		log.Println("receive message")
		log.Printf("msg headers : %s", msg.Headers)
		log.Printf("msg body %s", msg.Body)
		// save message
		// publish message to chatroom client
	}

}

func combineName(roomId string) string {
	return strings.Join([]string{PREFIX_CHATROOM_HUB_NAME, roomId}, ".")
}

func createRoomId() string {
	return uuid.New().String()
}
