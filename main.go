package main

import (
	"log"
	"net/http"
	"github.com/robfig/cron"
	"github.com/ilm-statistics/ilm-statistics/processor/resource"
)

func main() {
	log.Println("Adding cronometer")
	// Set timer to calculate and send statistics
	c := cron.New()
	c.AddFunc("@midnight", resource.SendStatistics)
	c.Start()
	log.Println("Cronometer started")

	// resource.SendStatistics()

	// Listen on port 8080 for incoming REST calls
	log.Println("Listening on port 8080")
	router := resource.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
