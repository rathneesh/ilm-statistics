package service

import (
	"log"
	"github.com/ilm-statistics/ilm-statistics/processor/util"
	"time"
	"github.com/ilm-statistics/ilm-statistics/processor/repository"
	"github.com/ilm-statistics/ilm-statistics/model"
)

func SendStatistics() {

	if !repository.IsDataForToday() {
		log.Println("There is no data for today")
		return
	}

	// Read statistics
	stats := repository.GetTodaysStatistic()

	// Calculate the average
	s := util.StatisticsCalculateAverages(stats)

	// Add the day
	s.Day = time.Now()

	// Email the statistics
	util.SendEmailTemplate(s)

	// Save today's averages to file
	repository.SaveStatisticsToFile(s)
}

func CreateStatistic(stat model.CollectedData) model.CollectedData{
	return repository.CreateStatistic(stat)
}