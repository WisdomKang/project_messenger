package queue

import (
	"testing"

	"github.com/rabbitmq/amqp091-go"
)

func TestCreateRoomId(t *testing.T) {
	t.Log("왜 안나옴?")
	roomId := createRoomId()

	if len(roomId) < 1 {
		t.Fail()
	}

	t.Log(roomId)
}

func TestConnection(t *testing.T) {
	conn, err := amqp091.Dial("amqp://user:password@localhost:5672")

	if err != nil {
		t.Log(err)
	}

	t.Log(conn.Properties)
}

func TestCreateChatRoom(t *testing.T) {

}

func TestCreateChatRoomQueue(t *testing.T) {

}
