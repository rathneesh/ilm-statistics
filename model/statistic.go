package model

import (
	"time"
)

type CollectedData struct {
	MAC          string
	Username     string
	Images       []Image
	Projects     []Project
	Builds       []Build
	Registries   []Registry
	Tests        []Test
	Results      []BuildResult
	Day          time.Time
}

type Build struct {
	Id        string `json:"id,omitempty"`
	ProjectId string `json:"projectId"`
	TestId    string `json:"testId"`
	StartTime string `json:"startTime,omitempty"`
	Status    Status `json:"status,omitempty"`
}

type Status struct {
	Status string `json:"status"`
}

type Results struct {
	ResultEntries []string
}

type Project struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name"`
	CreationTime string `json:"creationTime"`
	Status       string `json:"status"`
	ImageIds     []string `json:"imageids"`
	TestIds      []string `json:"testids"`
	Images       []Image `json:"images,omitempty"`
	Tests        []Test `json:"tests,omitempty"`
}

type Image struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name"`
	ImageId        string `json:"imageId"`
	Description    string `json:"description"`
	RegistryId     string `json:"registryId"`
	Tag            string `json:"tag"`
	IlmTags        []string `json:"ilmTags"`
	Location       string `json:"location"`
	SkipImageBuild  bool `json:"skipImageBuild"`
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
	Id   string `json:"id,omitempty" gorethink:"id,omitempty"`
	Name string `json:"name,omitempty" gorethink:"name,omitempty"`
	Addr string `json:"addr,omitempty" gorethink:"addr,omitempty"`
}

type Test struct {
	Id       string `json:"id,omitempty" gorethink:"id,omitempty"`
	Name     string `json:"name" gorethink:"name"`
	Provider Provider `json:"provider" gorethink:"provider"`
}

type Provider struct {
	ProviderType string `json:"providerType" gorethink:"providerType"`
}

type BuildResult struct {
	Id            string `json:"-" gorethink:"id,omitempty"`
	BuildId        string `json:"buildId" gorethink:"buildId"`
	ResultEntries  []string `json:"resultEntries" gorethink:"resultEntries"`
	TargetArtifact TargetArtifact `json:"targetArtifact" gorethink:"targetArtifact"`
}
type TargetArtifact struct {
	Artifact Artifact
}

type Artifact struct {
	ImageId string
	Link    string
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authentication struct {
	AuthToken string `json:"auth_token"`
}

type Statistic struct {
	Day               time.Time
	Users             int
	Accounts          int
	AvgAccountPerUser float64
	Projects          struct {
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
		Total  int
		Passed int
		Failed int
	}
	HourlyActivities         map[int]int
	BusiestHours             []int
	Registries               int
	MostPopularProjects      map[string]Project
	MaxProjectPopularity     int
	ImagesInProjects         map[string][]Project
	ProjectsSuccess          map[string]float64
	ProjectsFailure          map[string]float64
	ScriptProjects           []Project
	MostUsedImages           []string
	MostUsedImageOccurrence  int
	LeastUsedImages          []string
	LeastUsedImageOccurrence int
	NumberOfImages           int
	MostExecutedTests        []Test
	MostExecutedTestsNr      int
	LeastExecutedTests       []Test
	LeastExecutedTestsNr     int
	StatisticsPerUsers       map[string][]StatPerUser
}

type CollectedDataDiff struct {
	MAC                 string
	NewUserName         string
	AddedImages         []Image
	DeletedImages       []Image
	AddedProjects       []Project
	DeletedProjects     []Project
	AddedBuilds         []Build
	DeletedBuilds       []Build
	AddedRegistries     []Registry
	DeletedRegistries   []Registry
	AddedTests          []Test
	DeletedTests        []Test
	AddedResults        []BuildResult
	DeletedResults      []BuildResult
	AddedRepositories   []Repository
	DeletedRepositories []Repository
	NewDay 		    time.Time
}

type StatPerUser struct {
	Username string
	Day time.Time
	Vulnerabilities map[string]int
	NoOfVulnerabilities int
	NoOfImages int
}
