package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // test mode
		// AllowOrigins:    []string{"localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
	}))

	router.GET("/api", func(ctx *gin.Context) {

		testMap := map[string]any{
			"this": "possible?",
		}
		ctx.JSON(http.StatusOK, testMap)
	})

	router.Run()
}
