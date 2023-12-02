package main

import (
	"fmt"
	"messenger_server/pkg/queue"
)

func main() {

	var testUserId *string

	fmt.Scanln(testUserId)
	queue.CreateUserQueue(*testUserId)

}
