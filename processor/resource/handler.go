package resource

import (
	"io/ioutil"
	"io"
	"encoding/json"
	"net/http"
	"time"
	"log"
	"github.com/ilm-statistics/ilm-statistics/model"
	"github.com/ilm-statistics/ilm-statistics/processor/service"
	"math"
	"strings"
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
	var stat model.CollectedData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, math.MaxInt64))
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
	stat.Ip = strings.Split(r.RemoteAddr, ":")[0]
	stat.Day = time.Now()
	stat = service.CreateStatistic(stat)

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
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write(service.GetYesterdaysData())
}

func GetIp(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	log.Printf("IP of sender: %s", r.RemoteAddr)
}

func SendStatistics(){
	service.SendStatistics()
}

func SendStatisticsForced(w http.ResponseWriter, r *http.Request) {
	service.SendStatistics()
}