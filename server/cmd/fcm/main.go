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

	opt := option.WithCredentialsFile("my-messenger-project-388816-firebase-adminsdk-lqvzt-57f4ef07a9.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initialize app : %v\n", err)
	}

	message := &messaging.Message{
		Data:         map[string]string{},
		Notification: &messaging.Notification{},
		Token:        "dzChXvGCRW-i563F2-MUx6:APA91bFM3pMcW4Q1-_1HVaQaMtJv_XWv2jY1UOGaqlwHiwdxkAwlExDBrbzAJDNg-cxomDlWvAHtZs3xyl6u_JlZ_LOrRV1cSVbfsjH05Hl9Oo3qfigq20nqMCMJxXMNub7Iwv_pLrEn",
		// Token: "dkOrpxihRGiUUK3tOVvapa:APA91bFoiJEYDiY4KsHfPfSRUhvlIs8RsIBh27X-uturp7vuKytRMDs6V4VPKFQGI7RKeucbLGLQVibcqKIYicp0yQuTohZ4-19VNJf66qbxwixRl1331N976IWkKVLM-sxb05gZiqKT",
		// Token: "e-YVSrnfRFCOhEY2Obwr-b:APA91bGYWEYX7FEwCNuvRu53n367_--IuVVoDlWgq9GFLVIsJaSK82uyEaZ2bhIKZbwXAOsUvEpoCgfVwxMGVhNdL8Cdgeh84C80_z31JBtLx-NVUnD9r0yEb3Q_pHvmYYtM5Lj_an0D",
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
