package main

import (
	"messenger_server/pkg/rest"
)

func main() {

	rest.Router.Run(":8080")
}
