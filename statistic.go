package main

import (
	"time"
	"strings"
	"strconv"
	"log"
)

type SenderStatistics struct {
	Username     string
	Images       []ScriptImageDetails
	Accounts     []ScriptAccounts
	Projects     []ScriptProjects
	Builds       []ScriptBuilds
	Registries   []ScriptRegistries
	Tests        []ScriptTests
	Results      []ScriptBuildResults
	Repositories []ScriptRepository
	Day 	     time.Time
}

type ScriptBuilds struct {
	Id        string
	ProjectId string
	TestId    string
	StartTime string
	Status    Status
}

type Status struct {
	Status string
}

type Results struct {
	ResultEntries []string
}

type ScriptProjects struct {
	Id           string
	Name         string
	Author       string
	CreationTime string
	LastRunTime  string
	Status       string
	Images       []ScriptImageDetails
	Tests        []ScriptTests
}

type ScriptAccounts struct {
	Id        string
	FirstName string
	LastName  string
	Username  string
	Password  string
	Roles     []string
}

type ScriptImageDetails struct {
	ProjectId      string
	Id             string   // ---|
 	Name           string   //    | One of these is enough - sorry Denisa
	ImageId        string   // ---|
	Description    string
	Status         string
	RegistryId     string
	Tag            string
	IlmTags        []string
	Location       string
	SkipImageBuild string
}

type ScriptRepository struct {
	Name         string
	Tag          string
	FsLayers     []FsLayer
	Signatures   []Signature
	HasProblems  bool
	Message      string
	RegistryUrl  string
	RegistryName string
}

type FsLayer struct {
	BlobSum string
}
type Signature struct {
	Header    Header
	Signature string
	Protected string
}

type Header struct {
	Algorithm string
}
type ScriptRegistries struct {
	Id   string
	Name string
	Addr string
}
type ScriptTests struct {
	Id        string
	ProjectId string
	Provider  Provider
}
type Provider struct {
	providerType string
}
type ScriptBuildResults struct {
	ID            string
	BuildId       string
	ResultEntries []string
}

type Statistic struct {
	Day 				     time.Time
	Users 		                     int `json:"users"`
	Accounts 	                     int `json:"accounts"`
	AvgAccountPerUser                    float64 `json:"avgaccountperuser"`
	Projects struct {
			 Total               int `json:"total"`
			 ImagesInProjects    int `json:"imagesinprojects"`
			 AvgTestsInProjects  float64 `json:"avgtestsinprojects"`
			 AvgImagesInProjects float64 `json:"avgimagesinprojects"`
			 Passed              int `json:"passed"`
			 Failed              int `json:"failed"`
			 SuccessRate         float64 `json:"successrate"`
			 FailureRate         float64 `json:"failurerate"`
		 } 					`json:"projects"`
	Tests struct {
		Total                        int `json:"total"`
		Passed 	                     int `json:"passed"`
		Failed	                     int `json:"failed"`
	}					 `json:"tests"`

	HourlyActivities 		     map[int]int
	BusiestHours 			     []int
	Registries 	                     int `json:"registries"`
	MostPopularProjects		     map[string]ScriptProjects
	MaxProjectPopularity		     int
	ImagesInProjects		     map[string][]ScriptProjects
	ProjectsSuccess 		     map[string]float64
	ProjectsFailure			     map[string]float64
	ScriptProjects			     []ScriptProjects
	MostUsedImages			     []string
	MostUsedImageOccurence		     int
	LeastUsedImages			     []string
	LeastUsedImageOccurence		     int
}

func StatisticsCalculateAverages(stat []SenderStatistics) Statistic{
	var s Statistic
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
	s.ImagesInProjects = map[string][]ScriptProjects{}
	s.ProjectsFailure = map[string]float64{}
	s.ProjectsSuccess = map[string]float64{}
	projectsOtherOutcome := map[string]float64{}

	for i:=0; i<len(stat); i++ {
		s.Accounts += len(stat[i].Accounts)
		s.Projects.Total += len(stat[i].Projects)
		for j := 0; j < len(stat[i].Projects); j++ {
			if stat[i].Projects[j].Status == "finished_success" {
				s.Projects.Passed++
			} else if stat[i].Projects[j].Status == "finished_failed" {
				s.Projects.Failed++
			}

			//Images in projects
			for k := 0; k < len(stat[i].Projects[j].Images); k++ {
				s.ImagesInProjects[strings.Join([]string{stat[i].Projects[j].Images[k].Name, ":", stat[i].Projects[j].Images[k].Tag},"")] = append(s.ImagesInProjects[strings.Join([]string{stat[i].Projects[j].Images[k].Name, ":", stat[i].Projects[j].Images[k].Tag},"")], stat[i].Projects[j])
			}
		}


		for j := 0; j < len(stat[i].Builds); j++ {
			datetime := strings.Split(stat[i].Builds[j].StartTime,"T")
			time := strings.Split(datetime[1],":")
			hour := time[0]
			hr, err := strconv.Atoi(hour);
			if err != nil {
				log.Println(err)
				log.Println("Invalid input data")
			} else {
				for k := 0; k < len(stat[i].Tests); k++ {
					if (stat[i].Builds[j].ProjectId == stat[i].Tests[k].ProjectId) {
						s.HourlyActivities[hr]++
					}
				}
			}

			// Most popular projects
			projectsPopularity[stat[i].Builds[j].ProjectId]++

			// Project success/failure rate
			if stat[i].Builds[j].Status.Status == "finished_success" {
				s.ProjectsSuccess[stat[i].Builds[j].ProjectId]++
			} else if stat[i].Builds[j].Status.Status == "finished_failed" {
				s.ProjectsFailure[stat[i].Builds[j].ProjectId]++
			} else {
				projectsOtherOutcome[stat[i].Builds[j].ProjectId]++
			}
		}

		s.Projects.ImagesInProjects += len(stat[i].Images)
		//TODO find unique registries
		s.Registries += len(stat[i].Registries)
		s.Tests.Total += len(stat[i].Tests)

		//Most popular projects
		for projid, occurence := range projectsPopularity {
			if occurence > s.MaxProjectPopularity {
				projectId = make(map[string]bool)
				projectId[projid] = true
				s.MaxProjectPopularity = occurence
			} else if occurence == s.MaxProjectPopularity {
				projectId[projid] = true
			}
		}
	}

	//Most popular projects
	s.MostPopularProjects = make(map[string]ScriptProjects)
	for j := 0; j < len(stat); j++ {
		for k := 0; k < len(stat[j].Projects); k++ {
			s.ScriptProjects = append(s.ScriptProjects, stat[j].Projects[k])

			for id, _ := range projectId {
				if stat[j].Projects[k].Id == id {
					s.MostPopularProjects[id] = stat[j].Projects[k]
				}
			}

			//Project success/failure rate
			totalNoOfBuilds := s.ProjectsSuccess[stat[j].Projects[k].Id] + s.ProjectsFailure[stat[j].Projects[k].Id] + projectsOtherOutcome[stat[j].Projects[k].Id]
			if totalNoOfBuilds != 0 {
				s.ProjectsSuccess[stat[j].Projects[k].Id] = (s.ProjectsSuccess[stat[j].Projects[k].Id] * 100) / totalNoOfBuilds
				s.ProjectsFailure[stat[j].Projects[k].Id] = (s.ProjectsFailure[stat[j].Projects[k].Id] * 100) / totalNoOfBuilds
			} else {
				s.ProjectsSuccess[stat[j].Projects[k].Id] = 0
				s.ProjectsFailure[stat[j].Projects[k].Id] = 0
			}
		}
	}


	//Tests/hour
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

	s.MostUsedImageOccurence = 0
	s.LeastUsedImageOccurence = 10000000
	for imageName, projectList := range s.ImagesInProjects {
		if s.MostUsedImageOccurence < len(projectList){
			s.MostUsedImageOccurence = len(projectList)
			s.MostUsedImages = []string{imageName}
		} else if s.MostUsedImageOccurence == len(projectList) {
			s.MostUsedImages = append(s.MostUsedImages, imageName)
		}

		if s.LeastUsedImageOccurence > len(projectList) {
			s.LeastUsedImageOccurence = len(projectList)
			s.LeastUsedImages = []string{imageName}
		} else if s.LeastUsedImageOccurence == len(projectList) {
			s.LeastUsedImages = append(s.LeastUsedImages, imageName)
		}
	}


	s.AvgAccountPerUser = float64(s.Accounts)/float64(s.Users)
	s.Projects.AvgTestsInProjects = float64(s.Tests.Total)/float64(s.Projects.Total)
	s.Projects.AvgImagesInProjects = float64(s.Projects.ImagesInProjects)/float64(s.Projects.Total)
	s.Projects.SuccessRate = float64(s.Projects.Passed*100)/float64(s.Projects.Total)
	s.Projects.FailureRate = float64(s.Projects.Failed*100)/float64(s.Projects.Total)

	return s
}