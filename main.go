package main

import (
	"log"
	"net/http"
	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

func main() {
	log.Println("Adding cronometer")
	// Set timer to calculate and send statistics
	c := gron.New()
	c.AddFunc(gron.Every(1 * xtime.Day).At("00:00"), SendStatistics)
	c.Start()

	log.Println("Cronometer started")

	SendStatistics()

	// Listen on port 8080 for incoming REST calls
	log.Println("Listening on port 8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
