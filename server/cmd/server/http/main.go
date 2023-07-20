package main

import (
	"messenger_server/pkg/rest"
)

func main() {
	r := rest.Router
	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
