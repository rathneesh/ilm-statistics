package main

import (
	"github.com/ilm-statistics/ilm-statistics/processor/resource"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

func main() {
	log.Println("Adding cronometer")
	// Set timer to calculate and send statistics
	c := cron.New()
	c.AddFunc("@midnight", resource.SendStatistics)
	c.Start()
	log.Println("Cronometer started")

	// Listen on port 8084 for incoming REST calls
	log.Println("Listening on port 8084")
	router := resource.NewRouter()
	log.Fatal(http.ListenAndServe(":8084", router))
}
