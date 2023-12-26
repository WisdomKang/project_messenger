package rest

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// cors 설정
	Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT"},
		AllowHeaders: []string{"Origin"},
	}))

	Router.Use(checkMyURL)

	setUserRouter()
	setChatroomRouter()
	setMessageRouter()

}

// 
func setUserRouter() {
	userAPI := Router.Group("/user")
	userAPI.POST("" , func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	userAPI.PUT("")
	userAPI.DELETE("")
	userAPI.GET("/contract")
	userAPI.POST("/sync")

}

func setChatroomRouter() {
	chatroomAPI := Router.Group("/chatroom")
	chatroomAPI.POST("")
	chatroomAPI.GET("")
	chatroomAPI.GET("/messages")
}

func setMessageRouter() {
	messageAPI := Router.Group("/message")
	messageAPI.GET("")
}

func checkMyURL(r *gin.Context) {
	fmt.Printf("이거 호출이 되야하는데? URL : %s\n", r.Request.URL)

}
