package main

import "time"


type SenderStatistics struct {
	Username   string
	Images     []ScriptImageDetails
	Accounts   []ScriptAccounts
	Projects   []ScriptProjects
	Builds     []ScriptBuilds
	Registries []ScriptRegistries
	Tests      []ScriptTests
	Results    []ScriptBuildResults
	Day 	   time.Time
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
	Name         string
	Author       string
	CreationTime string
	LastRunTime  string
	Status       string
}

type ScriptAccounts struct {
	Id        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
}

type ScriptImageDetails struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ImageId     string `json:"imageId"`
	Description string `json:"description"`
	Status      string `json:"status"`
	RegistryId  string `json:"registryId"`

	Location       string `json:"location"`
	SkipImageBuild string `json:"skipImageBuild"`

	ProjectId string `json:"projectId"`
}

type ScriptRegistries struct {
	Name string
	Addr string
}
type ScriptTests struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
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
	Users 		                    int `json:"users"`
	Accounts 	                    int `json:"accounts"`
	AvgAccountPerUser                   float64 `json:"avgaccountperuser"`
	Projects struct {
		Total		            int `json:"total"`
		ImagesInProjects            int `json:"imagesinprojects"`
		AvgTestsInProjects          float64 `json:"avgtestsinprojects"`
		AvgImagesInProjects         float64 `json:"avgimagesinprojects"`
		Passed			    int `json:"passed"`
		Failed			    int `json:"failed"`
		SuccessRate		    float64 `json:"successrate"`
		FailureRate		    float64 `json:"failurerate"`
		NrProj0 int
		NrProj1 int
		NrProj2 int
		NrProj3 int
		NrProj4 int
		NrProj5 int
		NrProj6 int
		NrProj7 int
		NrProj8 int
		NrProj9 int
		NrProj10 int
		NrProj11 int
		NrProj12 int
		NrProj13 int
		NrProj14 int
		NrProj15 int
		NrProj16 int
		NrProj17 int
		NrProj18 int
		NrProj19 int
		NrProj20 int
		NrProj21 int
		NrProj22 int
		NrProj23 int
		BusiestHour string
	} 					`json:"projects"`
	Registries 	                    int `json:"registries"`
	Tests struct {
		Total                       int `json:"total"`
		Passed 	                    int `json:"passed"`
		Failed	                    int `json:"failed"`
//		SuccessRate                 float64 `json:"successrate"`
//		FailureRate                 float64 `json:"failurerate"`
	}					`json:"tests"`
	Difference struct {
		Users 		            int `json:"users"`
		Accounts 	            int `json:"accounts"`
		AvgAccountPerUser           float64 `json:"avgaccountperuser"`
		Projects struct {
			Total		    int `json:"total"`
			ImagesInProjects    int `json:"imagesinprojects"`
			AvgTestsInProjects  float64 `json:"avgtestsinprojects"`
			AvgImagesInProjects float64 `json:"avgimagesinprojects"`
//			Passed		    int `json:passed`
//			Failed		    int `json:failed`
		} 				`json:"projects"`
		Registries 	            int `json:"registries"`
		Tests struct {
			Total               int `json:"total"`
			Passed 	            int `json:"passed"`
			Failed	            int `json:"failed"`
			SuccessRate         float64 `json:"successrate"`
			FailureRate         float64 `json:"failurerate"`
		}				`json:"tests"`
	} 					`json:"difference"`
	Day time.Time
}

func StatisticCalculateDifference(yesterdayStat Statistic, statToday Statistic) Statistic {
	//TODO: calculate difference based on the final model

	return Statistic{}
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

	for i:=0; i<len(stat); i++ {
		s.Accounts += len(stat[i].Accounts)
		s.Projects.Total += len(stat[i].Projects)
		for j := 0; j < len(stat[i].Projects); j++ {
			if stat[i].Projects[j].Status == "finished_success" {
				s.Projects.Passed += 1
			} else if stat[i].Projects[j].Status == "finished_failed" {
				s.Projects.Failed += 1
			}
		}
		s.Projects.ImagesInProjects += len(stat[i].Images)
		//TODO find unique registries
		s.Registries += len(stat[i].Registries)
		s.Tests.Total += len(stat[i].Tests)
	}

	s.AvgAccountPerUser = float64(s.Accounts)/float64(s.Users)
	s.Projects.AvgTestsInProjects = float64(s.Tests.Total)/float64(s.Projects.Total)
	s.Projects.AvgImagesInProjects = float64(s.Projects.ImagesInProjects)/float64(s.Projects.Total)
	s.Projects.SuccessRate = float64(s.Projects.Passed*100)/float64(s.Projects.Total)
	s.Projects.FailureRate = float64(s.Projects.Failed*100)/float64(s.Projects.Total)

	return s
}