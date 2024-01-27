package main

import (
	"multi-tv/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.HandleMainPage)

	router.Run() // Listen and serve on :8080
}
