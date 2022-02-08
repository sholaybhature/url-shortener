package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shorten/controllers"
	"url-shorten/models"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!", ps.ByName("name"))
}

func main() {

	err := models.NewDatabase(":6379")
	if err != nil {
		log.Fatal("error database")
	}
	router := httprouter.New()
	router.POST("/api/v1/shorten", controllers.CreateShortenedURL)
	router.GET("/api/v1/analytics", controllers.GetURLAnalytics)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
