package main

import (
	"log"
	"net/http"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Adding cronometer")
	// Set timer to calculate and send statistics
	c := cron.New()
	c.AddFunc("@midnight", SendStatistics)
	c.Start()

	log.Println("Cronometer started")

	SendStatistics()

	// Listen on port 8080 for incoming REST calls
	log.Println("Listening on port 8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
