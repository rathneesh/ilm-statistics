package main

import (
	"io/ioutil"
	"io"
	"encoding/json"
	"net/http"
	"time"
	"log"
)

func CreateNewStatistic(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var stat SenderStatistics
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &stat); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	stat.Day = time.Now()
	stat = CreateStatistic(stat)

	if err := json.NewEncoder(w).Encode(stat); err != nil {
		panic(err)
	}
}

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(GetTodaysStatistic()); err != nil {
		panic(err)
	}
}