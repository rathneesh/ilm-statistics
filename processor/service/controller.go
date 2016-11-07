package service

import (
	"log"
	"github.com/ilm-statistics/ilm-statistics/processor/util"
	"time"
	"github.com/ilm-statistics/ilm-statistics/processor/repository"
	"github.com/ilm-statistics/ilm-statistics/model"
)

func SendStatistics() {

	// Read statistics
	stats := repository.GetDataForEmail()

	// Calculate the average
	log.Println("Calculating the average...")
	s, sforIps := util.StatisticsCalculateAverages(stats)

	// Add the day
	s.Day = time.Now()

	// Email the statistics
	go func() {
		attachment, err := util.SendEmailTemplate(s, sforIps)
		if err != nil {
			log.Println(err)
		} else {
			// Save today's averages to file
			repository.SaveStatisticsToFile(attachment)
		}
	}()


}

func CreateStatistic(stat model.CollectedData) model.CollectedData{
	return repository.CreateStatistic(stat)
}

func GetTodaysData() []model.CollectedData {
	return repository.GetTodaysData()
}