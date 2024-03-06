package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Any("/auth", func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		accessTokenHeader := ctx.Request.Header.Get("AccessToken")

		// Test : All request will pass. if it has Headers value.
		if len(authHeader) <= 0 || len(accessTokenHeader) <= 0 {
			log.Printf("해당 요청이 조건을 만족하지 못합니다")
			ctx.Status(http.StatusUnauthorized)
			return
		} else {
			ctx.Status(http.StatusOK)
			return
		}

	})

	server.Run(":8080")

}
