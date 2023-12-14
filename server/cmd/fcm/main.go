package main

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {

	opt := option.WithCredentialsFile("project_privite_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initialize app : %v\n", err)
	}

	message := &messaging.Message{
		Data:         map[string]string{},
		Notification: &messaging.Notification{},
		Token:        "Firebase Token ",
	}

	msgClient, err := app.Messaging(context.Background())

	if err != nil {
		log.Fatalf("error msgClient create : %s\n", err)
	}

	for i := 1; i <= 3; i++ {
		message.Data["title"] = "제목입니다"
		message.Data["body"] = fmt.Sprintf("%d번째 메세지", i)
		response, err := msgClient.Send(context.Background(), message)

		if err != nil {
			log.Fatalf("error send message : %s\n", err)
		}

		log.Printf("Successfully sent message : %s", response)

		time.Sleep(time.Second * 1)
	}
}
