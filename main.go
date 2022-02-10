package main

import (
	"log"
	"net/http"
	"os"
	"url-shorten/controllers"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	redisPort := os.Getenv("REDIS_PORT")
	httpPort := os.Getenv("HTTP_PORT")
	// Connect to database
	err := models.NewDatabase(redisPort)
	if err != nil {
		log.Fatal("error database")
	}
	router := httprouter.New()
	router.GET("/", controllers.HomePage)
	router.GET("/id/:id", controllers.RedirectShortenedURL)
	router.POST("/api/v1/shorten", controllers.CreateShortenedURL)
	router.GET("/api/v1/analytics", controllers.GetURLAnalytics)
	log.Fatal(http.ListenAndServe(httpPort, router))
}
