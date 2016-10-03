package util

import (
	"math"
	"strings"
	"time"
	"strconv"
	"log"
	"github.com/ilm-statistics/ilm-statistics/model"
	"regexp"
)

const (
	SUCCESS = "finished_success"
	FAILURE = "finished_failed"
)

func StatisticsCalculateAverages(stat []model.CollectedData) model.Statistic{

	//Initializing the required values
	var s model.Statistic
	s.Users = len(stat);
	s.Accounts = 0
	s.Projects.Total = 0
	s.Projects.ImagesInProjects = 0
	s.Projects.Passed = 0
	s.Projects.Failed = 0
	s.Registries = 0
	s.Tests.Total = 0
	s.HourlyActivities = map[int]int{}
	projectsPopularity := map[string]int{}
	s.MaxProjectPopularity = 0
	var projectId map[string]bool
	s.ImagesInProjects = map[string][]model.Project{}
	s.ProjectsFailure = map[string]float64{}
	s.ProjectsSuccess = map[string]float64{}
	projectsOtherOutcome := map[string]float64{}
	s.MostExecutedTestsNr = 0
	s.LeastExecutedTestsNr = math.MaxInt32
	buildsToTest := map[string][]model.Build{}
	s.StatisticsPerUsers = map[string][]model.StatPerUser{}


	//iterate through the statistics
	for i:=0; i<len(stat); i++ {

		//Number all the accounts and projects
		s.Projects.Total += len(stat[i].Projects)

		//Overall project success/failure rate
		for j := 0; j < len(stat[i].Projects); j++ {
			if stat[i].Projects[j].Status == SUCCESS {
				s.Projects.Passed++
			} else if stat[i].Projects[j].Status == FAILURE {
				s.Projects.Failed++
			}

			//Images in projects
			for k := 0; k < len(stat[i].Projects[j].Images); k++ {
				s.ImagesInProjects[strings.Join([]string{stat[i].Projects[j].Images[k].Name, ":", stat[i].Projects[j].Images[k].Tag},"")] = append(s.ImagesInProjects[strings.Join([]string{stat[i].Projects[j].Images[k].Name, ":", stat[i].Projects[j].Images[k].Tag},"")], stat[i].Projects[j])
			}
		}


		//Iterate through the builds
		for j := 0; j < len(stat[i].Builds); j++ {

			//Get the build's date and today's date in different variables - for the comparision below
			datetime := strings.Split(stat[i].Builds[j].StartTime,"T")
			date := strings.Split(datetime[0],"-")
			now := time.Now()


			year, err := strconv.Atoi(date[0])
			if err != nil {
				log.Println(err)
			}

			month, err := strconv.Atoi(date[1])
			if err != nil {
				log.Println(err)
			}

			day, err := strconv.Atoi(date[2])
			if err != nil {
				log.Println(err)
			}

			//It is triggered at midnight, has to calculate for the day before
			day += 1


			t := strings.Split(datetime[1],":")
			hour := t[0]
			hr, err := strconv.Atoi(hour);
			if err != nil {
				log.Println(err)
			} else {
				//Number today's hourly activities
				if (year == now.Year() && time.Month(month) == now.Month() && day == now.Day()){
					for k := 0; k < len(stat[i].Tests); k++ {
						if (stat[i].Builds[j].TestId == stat[i].Tests[k].Id) {
							s.HourlyActivities[hr]++
						}
					}
				}
			}

			// Most popular projects
			projectsPopularity[stat[i].Builds[j].ProjectId]++

			//Per project success/failure rate
			if stat[i].Builds[j].Status.Status == SUCCESS {
				s.ProjectsSuccess[stat[i].Builds[j].ProjectId]++
			} else if stat[i].Builds[j].Status.Status == FAILURE {
				s.ProjectsFailure[stat[i].Builds[j].ProjectId]++
			} else {
				projectsOtherOutcome[stat[i].Builds[j].ProjectId]++
			}

			// Most/Least executed tests
			buildsToTest[stat[i].Builds[j].TestId] = append(buildsToTest[stat[i].Builds[j].TestId], stat[i].Builds[j])
		}

		s.Projects.ImagesInProjects += len(stat[i].Images)

		//TODO find unique registries - when statistics regarding registries will be needed
		s.Registries += len(stat[i].Registries)
		s.Tests.Total += len(stat[i].Tests)

		//Most popular projects
		for projid, occurrence := range projectsPopularity {
			if occurrence > s.MaxProjectPopularity {
				projectId = make(map[string]bool)
				projectId[projid] = true
				s.MaxProjectPopularity = occurrence
			} else if occurrence == s.MaxProjectPopularity {
				projectId[projid] = true
			}
		}
	}

	//Most popular projects
	s.MostPopularProjects = make(map[string]model.Project)
	for j := 0; j < len(stat); j++ {
		for k := 0; k < len(stat[j].Projects); k++ {
			s.ScriptProjects = append(s.ScriptProjects, stat[j].Projects[k])

			for id, _ := range projectId {
				if stat[j].Projects[k].Id == id {
					s.MostPopularProjects[id] = stat[j].Projects[k]
				}
			}

			//Project success/failure rate - calculate in percents
			totalNoOfBuilds := s.ProjectsSuccess[stat[j].Projects[k].Id] + s.ProjectsFailure[stat[j].Projects[k].Id] + projectsOtherOutcome[stat[j].Projects[k].Id]
			if totalNoOfBuilds != 0 {
				s.ProjectsSuccess[stat[j].Projects[k].Id] = float64(s.ProjectsSuccess[stat[j].Projects[k].Id] * 100) / float64(totalNoOfBuilds)
				s.ProjectsFailure[stat[j].Projects[k].Id] = float64(s.ProjectsFailure[stat[j].Projects[k].Id] * 100) / float64(totalNoOfBuilds)
			} else {
				s.ProjectsSuccess[stat[j].Projects[k].Id] = 0
				s.ProjectsFailure[stat[j].Projects[k].Id] = 0
			}
		}

		// Most/least executed tests
		for _, test := range stat[j].Tests {
			buildListLength := len(buildsToTest[test.Id])
			if s.MostExecutedTestsNr < buildListLength {
				s.MostExecutedTestsNr = buildListLength
				s.MostExecutedTests = []model.Test{test}
			} else if s.MostExecutedTestsNr == buildListLength {
				s.MostExecutedTests = append(s.MostExecutedTests, test)
			}

			if s.LeastExecutedTestsNr > buildListLength {
				s.LeastExecutedTestsNr = buildListLength
				s.LeastExecutedTests = []model.Test{test}
			} else if s.LeastExecutedTestsNr == buildListLength {
				s.LeastExecutedTests = append(s.LeastExecutedTests, test)
			}
		}
	}


	//Tests/hour - busiest hours
	s.BusiestHours = []int{}
	max := s.HourlyActivities[0]
	for i := 1; i < 24; i++ {
		if s.HourlyActivities[i] > max {
			s.BusiestHours = []int{i}
			max = s.HourlyActivities[i]
		} else if s.HourlyActivities[i] == max {
			s.BusiestHours = append(s.BusiestHours, i)
		}
	}

	// Most/least used images and their occurences
	s.MostUsedImageOccurrence = 0
	s.LeastUsedImageOccurrence = math.MaxInt32
	for imageName, projectList := range s.ImagesInProjects {
		if s.MostUsedImageOccurrence < len(projectList){
			s.MostUsedImageOccurrence = len(projectList)
			s.MostUsedImages = []string{imageName}
		} else if s.MostUsedImageOccurrence == len(projectList) {
			s.MostUsedImages = append(s.MostUsedImages, imageName)
		}

		if s.LeastUsedImageOccurrence > len(projectList) {
			s.LeastUsedImageOccurrence = len(projectList)
			s.LeastUsedImages = []string{imageName}
		} else if s.LeastUsedImageOccurrence == len(projectList) {
			s.LeastUsedImages = append(s.LeastUsedImages, imageName)
		}
	}


	//Statistics per user - days, nr of images, nr of vulnerabilities
	for _, st := range stat {
		auxStatPerUser := model.StatPerUser{}
		auxStatPerUser.Day = st.Day
		auxStatPerUser.Username = st.Username
		auxStatPerUser.NoOfVulnerabilities = 0
		auxStatPerUser.NoOfImages = 0
		auxStatPerUser.Vulnerabilities = map[string]int{}
		for _, results := range st.Results {
			for _, entry := range results.ResultEntries {
				if strings.Contains(entry, "SUCCESS:") || strings.Contains(entry, "FAILURE:") {
					re := regexp.MustCompile("[0-9]+")
					vulnerabilities := re.FindAllString(entry, -1)
					noOfVulnerabilities, err := strconv.Atoi(vulnerabilities[0])
					reImage := regexp.MustCompile("[^ ]+:[^ ]+")
					registryAndImage := reImage.FindAllString(entry, -1)
					if err != nil {
						log.Println(err)
					} else {
						auxStatPerUser.Vulnerabilities[strings.Join(registryAndImage, " ")] = noOfVulnerabilities
						auxStatPerUser.NoOfImages ++
						auxStatPerUser.NoOfVulnerabilities += noOfVulnerabilities
					}
				}
			}
		}
		s.StatisticsPerUsers[st.MAC] = append(s.StatisticsPerUsers[st.MAC], auxStatPerUser)
	}

	s.NumberOfImages = len(s.ImagesInProjects)
	s.Projects.AvgTestsInProjects = float64(s.Tests.Total)/float64(s.Projects.Total)
	s.Projects.AvgImagesInProjects = float64(s.Projects.ImagesInProjects)/float64(s.Projects.Total)
	s.Projects.SuccessRate = float64(s.Projects.Passed*100)/float64(s.Projects.Total)
	s.Projects.FailureRate = float64(s.Projects.Failed*100)/float64(s.Projects.Total)

	// Return all in a statistics object
	return s
}