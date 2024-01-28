package rest

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	setUserRouter()
	setChatroomRouter()
	setMessageRouter()

}

var secretKey = []byte("test-secret-key")

func Authorization(r *gin.Context) {
	authHeader := r.Request.Header.Get("Authorization")
	accessTokenHeader := r.Request.Header.Get("AccessToken")

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	accessTokenString := strings.TrimPrefix(accessTokenHeader, "Bearer")

	log.Printf("idToken : %s", tokenString)
	log.Printf("accessToken : %s", accessTokenString)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return "t", nil
	})

	if err != nil {
		log.Fatalf("err : %s", err)
	} else {
		log.Printf("parse token : %s", token.Signature)
	}
}

// user api
func setUserRouter() {
	userAPI := Router.Group("/user")
	userAPI.POST("", func(ctx *gin.Context) {
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
