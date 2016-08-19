package model

import (
	"time"
)

type CollectedData struct {
	Username     string
	Images       []Image
	Accounts     []Account
	Projects     []Project
	Builds       []Build
	Registries   []Registry
	Tests        []Test
	Results      []BuildResult
	Repositories []Repository
	Day 	     time.Time
}

type Build struct {
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

type Project struct {
	Id           string
	Name         string
	Author       string
	CreationTime string
	LastRunTime  string
	Status       string
	Images       []Image
	Tests        []Test
}

type Account struct {
	Id        string
	FirstName string
	LastName  string
	Username  string
	Password  string
	Roles     []string
}

type Image struct {
	ProjectId      string
	Id             string
 	Name           string
	ImageId        string
	Description    string
	Status         string
	RegistryId     string
	Tag            string
	IlmTags        []string
	Location       string
	SkipImageBuild string
}

type Repository struct {
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
type Registry struct {
	Id   string
	Name string
	Addr string
}
type Test struct {
	Id        string
	ProjectId string
	Provider  Provider
}
type Provider struct {
	ProviderType string
}
type BuildResult struct {
	ID            string
	BuildId       string
	ResultEntries []string
}

type Statistic struct {
	Day 				     time.Time
	Users 		                     int
	Accounts 	                     int
	AvgAccountPerUser                    float64
	Projects struct {
			 Total               int
			 ImagesInProjects    int
			 AvgTestsInProjects  float64
			 AvgImagesInProjects float64
			 Passed              int
			 Failed              int
			 SuccessRate         float64
			 FailureRate         float64
		 }
	Tests struct {
		Total                        int
		Passed 	                     int
		Failed	                     int
	}

	HourlyActivities 		     map[int]int
	BusiestHours 			     []int
	Registries 	                     int
	MostPopularProjects		     map[string]Project
	MaxProjectPopularity		     int
	ImagesInProjects		     map[string][]Project
	ProjectsSuccess 		     map[string]float64
	ProjectsFailure			     map[string]float64
	ScriptProjects			     []Project
	MostUsedImages			     []string
	MostUsedImageOccurrence		     int
	LeastUsedImages			     []string
	LeastUsedImageOccurrence		     int
	NumberOfImages			     int
	MostExecutedTests		     []Test
	MostExecutedTestsNr		     int
	LeastExecutedTests		     []Test
	LeastExecutedTestsNr		     int
}