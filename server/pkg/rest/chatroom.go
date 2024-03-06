package rest

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateChatRoom(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	ErrorHandler(ctx, err)

	log.Printf("request receive! : %s\n", string(body))

}

func UpdateChatRoom(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	ErrorHandler(ctx, err)

}

func GetChatRoom(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	ErrorHandler(ctx, err)

}
