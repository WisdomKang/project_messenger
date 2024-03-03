package main

import (
	"fmt"
	"log"
	"messenger_server/pkg/queue"
)

func handleError(logMessage string, err error) {
	if err != nil {
		log.Fatalf("%s : %s", logMessage, err)
	}
}

func main() {
	test := queue.Config{}

	fmt.Print(test)

}
