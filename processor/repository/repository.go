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
	"path/filepath"
)

const (
	FOLDER = "data/"
	HTMLFOLDER = "web/"
	EXTENSION = ".json"
	HTMLEXTENSION = ".html"
	DELIMITER = "-"
)
var tmpfilename string
var tmpHtmlFileName string
var statistics map[string]model.CollectedData
var diffData []model.CollectedDataDiff
var fileMutex sync.Mutex

func init() {
	log.Println("Initializing repository")
	statistics = map[string]model.CollectedData{}
	UpdateTmpFileName()
	InitFromFile()
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

func GetYesterdaysData() []byte {
	log.Println("Getting today's statistics")

	yesterday := time.Now().Add(-24 * time.Hour)
	tmpDataHtmlAux := []string{HTMLFOLDER, strconv.Itoa(int(yesterday.Month())), DELIMITER, strconv.Itoa(int(yesterday.Day())), DELIMITER, strconv.Itoa(int(yesterday.Year())), HTMLEXTENSION}
	tmpDataHtml := strings.Join(tmpDataHtmlAux, "")

	data, err := ioutil.ReadFile(tmpDataHtml)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	return data
}

func UpdateTmpFileName(){
	tmpData := []string{FOLDER, strconv.Itoa(int(time.Now().Month())), DELIMITER, strconv.Itoa(int(time.Now().Day())), DELIMITER, strconv.Itoa(int(time.Now().Year())), EXTENSION}
	tmpDataHtml := []string{HTMLFOLDER, strconv.Itoa(int(time.Now().Month())), DELIMITER, strconv.Itoa(int(time.Now().Day())), DELIMITER, strconv.Itoa(int(time.Now().Year())), HTMLEXTENSION}
	if tmpfilename != strings.Join(tmpData, "") {
		statistics = map[string]model.CollectedData{}
		tmpfilename = strings.Join(tmpData, "")
	}
	if tmpHtmlFileName != strings.Join(tmpDataHtml, ""){
		tmpHtmlFileName = strings.Join(tmpDataHtml, "")
	}
}

func SaveStatisticsToFile(attachment []byte){
	log.Println("Saving the sent statistics to a html file")
	f, err := os.OpenFile(tmpHtmlFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	fileMutex.Lock()
	if _, err = f.Write(attachment); err != nil {
		log.Println(err)
		return
	}
	defer fileMutex.Unlock()
	log.Println("Saved the statistics to ./"+tmpHtmlFileName)
	UpdateTmpFileName()
}

func InitFromFile(){
	log.Println("Loading past data into memory from", tmpfilename)
	os.Mkdir("." + string(filepath.Separator) + "data", 0777)
	os.Mkdir("." + string(filepath.Separator) + "web", 0777)


	statistics = map[string]model.CollectedData{}
	diffData = []model.CollectedDataDiff{}

	data, err := ioutil.ReadFile(tmpfilename)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(data, &diffData)
	if err != nil {
		log.Println(err)
		return
	}

	for _, collData := range diffData {
		if statistics[collData.MAC].MAC == ""{
			statistics[collData.MAC] = model.CollectedData{MAC:collData.MAC}
		}

		statistics[collData.MAC] = util.MergeDiff(statistics[collData.MAC], collData)
	}
}

func GetDataForEmail() []model.CollectedData {
	log.Println("Loading past data into memory from the data folder")

	var statsDiff, statsDiffAux []model.CollectedDataDiff

	files, _ := ioutil.ReadDir("."+string(filepath.Separator)+FOLDER)

	for _, f := range files {
		statsDiffAux = readDiffsFromFile("."+string(filepath.Separator)+FOLDER + f.Name())
		statsDiff = append(statsDiff, statsDiffAux...)
	}

	statisticsMap := map[string]model.CollectedData{}

	stats := []model.CollectedData{}

	for _, collData := range statsDiff {
		if statisticsMap[collData.MAC].MAC == ""{
			statisticsMap[collData.MAC] = model.CollectedData{MAC:collData.MAC}
		}

		statisticsMap[collData.MAC] = util.MergeDiff(statisticsMap[collData.MAC], collData)
		stats = append(stats, statisticsMap[collData.MAC])
	}

	return stats
}

func readDiffsFromFile(fileName string) []model.CollectedDataDiff{
	log.Println("Loading data from "+fileName)
	diffData = []model.CollectedDataDiff{}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return diffData
	}
	err = json.Unmarshal(data, &diffData)
	if err != nil {
		log.Println(err)
		return diffData
	}

	return diffData
}