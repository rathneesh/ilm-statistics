package service

import (
	"github.com/ilm-statistics/ilm-statistics/model"
	"github.com/ilm-statistics/ilm-statistics/processor/repository"
	"github.com/ilm-statistics/ilm-statistics/processor/util"
	"log"
	"time"
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

func CreateStatistic(stat model.CollectedData) model.CollectedData {
	return repository.CreateStatistic(stat)
}

func GetYesterdaysData() []byte {
	return repository.GetYesterdaysData()
}
