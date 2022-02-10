package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shorten/controllers"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Hello, World.")
	redisPort := os.Getenv("REDIS_PORT")
	httpPort := os.Getenv("HTTP_PORT")
	fmt.Println(redisPort, httpPort)
	// err := models.NewDatabase(redisPort)
	err := models.NewDatabase("redis:6379")
	if err != nil {
		fmt.Print(err)
		log.Fatal("error database")
	}
	router := httprouter.New()
	router.GET("/", controllers.HomePage)
	router.GET("/id/:id", controllers.RedirectShortenedURL)
	router.POST("/api/v1/shorten", controllers.CreateShortenedURL)
	router.GET("/api/v1/analytics", controllers.GetURLAnalytics)
	// log.Fatal(http.ListenAndServe(httpPort, router))
	log.Fatal(http.ListenAndServe(":8080", router))
}
