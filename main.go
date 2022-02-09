package main

import (
	"log"
	"net/http"
	"url-shorten/controllers"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func main() {

	err := models.NewDatabase(":6379")
	if err != nil {
		log.Fatal("error database")
	}
	router := httprouter.New()
	router.GET("/", controllers.HomePage)
	router.GET("/id/:id", controllers.RedirectShortenedURL)
	router.POST("/api/v1/shorten", controllers.CreateShortenedURL)
	router.GET("/api/v1/analytics", controllers.GetURLAnalytics)
	log.Fatal(http.ListenAndServe(":8080", router))
}
