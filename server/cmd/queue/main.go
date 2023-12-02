package main

import (
	"log"
	"messenger_server/pkg/queue"
)

func main() {

	messageController := queue.MController

	messageController.InitController()

	userList := []string{"test0001", "test0002"}
	testChatRoom := queue.CreateChatRoom(userList)

	log.Println(testChatRoom.ID)

	wating := make(chan bool)

	<-wating

}
