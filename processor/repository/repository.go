package repository

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
	"sync"
	"github.com/ilm-statistics/ilm-statistics/model"
	"github.com/ilm-statistics/ilm-statistics/processor/util"
	"os"
)

const (
	FILENAME = "data/statistics.json"
	FOLDER = "data/"
	EXTENSION = ".json"
	DELIMITER = "-"
)
var tmpfilename string
var statistics map[string]model.CollectedData
var diffData []model.CollectedDataDiff
var fileMutex sync.Mutex

func init() {
	log.Println("Initializing repository")
	UpdateTmpFileName()
	statistics = map[string]model.CollectedData{}
	diffData = []model.CollectedDataDiff{}
}

// Save the incoming data into a file
func CreateStatistic(s model.CollectedData) model.CollectedData {
	log.Println("Saving the statistic from user: ", s.Username)

	if util.CmpCollectedData(s, model.CollectedData{}) {
		log.Println("Empty data was posted")
		return model.CollectedData{}
	}

	diffData = append(diffData, util.DiffCollectedData(statistics[s.MAC], s))

	statistics[s.MAC] = s

	data, err := json.MarshalIndent(diffData, "", "  ")
	if err != nil {
		log.Println(err)
		return model.CollectedData{}
	}

	fileMutex.Lock()
	err = ioutil.WriteFile(tmpfilename, data, 0666)
	defer fileMutex.Unlock()
	if err != nil {
		log.Println(err)
		return model.CollectedData{}
	}
	return s
}

func GetTodaysStatistic() []model.CollectedData {
	log.Println("Getting today's statistics")

	data, err := ioutil.ReadFile(tmpfilename)
	if err != nil {
		log.Println(err)
		return []model.CollectedData{}
	}
	var statsDiff []model.CollectedDataDiff
	err = json.Unmarshal(data, &statsDiff)
	if err != nil {
		log.Println(err)
		return []model.CollectedData{}
	}

	statisticsMap := map[string]model.CollectedData{}
	stats := []model.CollectedData{}

	// Suppose the data was saved ordered

	for _, collData := range statsDiff {
		stats = append(stats, util.MergeDiff(statisticsMap[collData.MAC], collData))
		statisticsMap[collData.MAC] = util.MergeDiff(statisticsMap[collData.MAC], collData)
	}

	return stats
}
func UpdateTmpFileName(){
	a := []string{FOLDER,strconv.Itoa(int(time.Now().Month())), DELIMITER, strconv.Itoa(int(time.Now().Day())), DELIMITER, strconv.Itoa(int(time.Now().Year())),EXTENSION}
	tmpfilename = strings.Join(a, "")
}

func IsDataForToday() bool{
	// Check if data exists for today
	if _, err := os.Stat(tmpfilename); os.IsNotExist(err) {
		log.Println("File for today's data does not exists")
		return false
	}
	return true
}

func SaveStatisticsToFile(s model.Statistic){
	f, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	data, err := json.MarshalIndent(s, "", "  ")
	if _, err = f.Write(data); err != nil {
		log.Println(err)
		return
	}

	UpdateTmpFileName()
}