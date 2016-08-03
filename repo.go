package main

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"net/smtp"
	"strings"
	"strconv"
	"os"
	"log"
)

const FILENAME = "data/statistics.json"
var tmpfilename string
var statistics []SenderStatistics

func init() {
	log.Println("Initializing repository")
	a := []string{"data/",strconv.Itoa(int(time.Now().Month())), "-", strconv.Itoa(int(time.Now().Day())), "-", strconv.Itoa(int(time.Now().Year())), ".json"}
	tmpfilename = strings.Join(a, "")
}

// Save the incoming stats into a file
func CreateStatistic(s SenderStatistics) SenderStatistics {
	log.Println("Saving the statistic from user", s.Username)
	statistics = append(statistics, s)
	data, err := json.Marshal(statistics)
	if err != nil {
		log.Println(err)
		return SenderStatistics{}
	}
	err = ioutil.WriteFile(tmpfilename, data, 0666)
	if err != nil {
		log.Println(err)
		return SenderStatistics{}
	}
	return s
}

func GetTodaysStatistic() []SenderStatistics {
	log.Println("Getting today's statistics")

	data, err := ioutil.ReadFile(tmpfilename)
	if err != nil {
		log.Println(err)
		return []SenderStatistics{}
	}
	var stats []SenderStatistics
	err = json.Unmarshal(data, &stats)
	if err != nil {
		log.Println(err)
		return []SenderStatistics{}
	}
	return stats
}

func YesterdaysStatistics() Statistic {
	// Get the averages of yesterday to calculate a difference
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	var stats []Statistic
	err = json.Unmarshal(data, &stats)

	yesterday := time.Now().AddDate(0,0,-1)
	for _, s := range stats {
		if (s.Day.Day() == yesterday.Day()) && (s.Day.Month() == yesterday.Month()) && (s.Day.Year() == yesterday.Year()) {
			return s
		}
	}
	return Statistic{}
}

func SendStatistics() {

	// Check if data exists for today
	if _, err := os.Stat(tmpfilename); os.IsNotExist(err) {
		log.Println("File for today's data does not exists")
		return
	}

	// Read statistics
	stats := GetTodaysStatistic()

	// Calculate the average
	s := StatisticsCalculateAverages(stats)

	// Calculate the difference
	//s = StatisticCalculateDifference(YesterdaysStatistics(), s)

	// Add the day
	s.Day = time.Now()

	// Email the statistics
	//EmailStatistics(s)
	SendEmailTemplate(s)

	// Save today's averages to the file
	f, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	data, err := json.Marshal(s)
	if _, err = f.Write(data); err != nil {
		log.Println(err)
		return
	}

	UpdateTmpFileName()
}

func UpdateTmpFileName(){
	a := []string{"data/",strconv.Itoa(int(time.Now().Month())), "-", strconv.Itoa(int(time.Now().Day())), "-", strconv.Itoa(int(time.Now().Year())),".json"}
	tmpfilename = strings.Join(a, "")
}