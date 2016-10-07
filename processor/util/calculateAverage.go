package util

import (
	"math"
	"strings"
	"time"
	"strconv"
	"log"
	"github.com/ilm-statistics/ilm-statistics/model"
	"regexp"
	"sort"
)

const (
	SUCCESS = "finished_success"
	FAILURE = "finished_failed"
)

func StatisticsCalculateAverages(dataList []model.CollectedData) model.Statistic{

	idToProject := mapIdToProject(dataList)
	idToTest := mapIdToTest(dataList)
	nameToImage := mapNameToImage(dataList)
	idToBuild := mapIdToBuild(dataList)
	idToRegistry := mapIdToRegistry(dataList)

	//Initializing the required values
	s := model.Statistic{}

	s.Projects.IdToProject = idToProject
	s.Projects.Total = CalculateNumberOfProjects(idToProject)
	s.Projects.ImagesInProjects = CalculateNumberOfImages(nameToImage)
	s.Projects.AvgTestsInProjects = CalculateAverageTestPerProject(idToTest, idToProject)
	s.Projects.AvgImagesInProjects = CalculateAverageImagePerProject(nameToImage, idToProject)
	s.Projects.SuccessRate, s.Projects.FailureRate = CalculateAllProjectsOutcomeRates(idToProject)
	s.Tests.Total = CalculateNumberOfTests(idToTest)
	s.MostUsedImages = CalculateMostUsedImages(idToProject)
	s.HourlyActivities = CalculateNumberOfTestsInEachHour(idToBuild, idToTest)
	s.BusiestHours = CalculateBusiestHours(idToBuild, idToTest)
	s.MostPopularProjects, s.MaxProjectPopularity = CalculateMostExecutedProjects(idToBuild, idToProject)
	s.ImagesInProjects = ShowImagesInProjects(idToProject)
	s.ProjectsSuccess, s.ProjectsFailure = CalculatePerProjectOutcomeRates(idToBuild, idToProject)
	s.MostExecutedTests, s.MostExecutedTestsNr = CalculateMostExecutedTests(idToBuild, idToTest)
	s.LeastExecutedTests, s.LeastExecutedTestsNr = CalculateLeastExecutedTests(idToBuild, idToTest)
	s.Vulnerabilities = CalculateNoOfVulnerabilitiesFound(dataList, idToBuild)
	s.RegistriesAndImages = ShowImagesInRegistries(nameToImage, idToRegistry)


	return s
}

func mapIdToProject(dataList []model.CollectedData) map[string]model.Project {
	idToProject := map[string]model.Project{}
	for _, data := range dataList{
		for _, project := range data.Projects {
			idToProject[project.Id] = project
		}
	}
	return idToProject
}

func mapIdToTest(dataList []model.CollectedData) map[string]model.Test {
	idToTest := map[string]model.Test{}
	for _, data := range dataList{
		for _, test := range data.Tests {
			idToTest[test.Id] = test
		}
	}
	return idToTest
}

func mapNameToImage(dataList []model.CollectedData) map[string]model.Image {
	nameToImage := map[string]model.Image{}
	for _, data := range dataList {
		for _, image := range data.Images {
			nameToImage[image.Name+":"+image.Tag] = image
		}
	}
	return nameToImage
}

func mapIdToBuild(dataList []model.CollectedData) map[string]model.Build {
	idToBuild := map[string]model.Build{}
	for _, data := range dataList {
		for _, build := range data.Builds {
			idToBuild[build.Id] = build
		}
	}
	return idToBuild
}

func mapIdToRegistry(dataList []model.CollectedData) map[string]model.Registry {
	idToRegistry := map[string]model.Registry{}
	for _, data := range dataList {
		for _, registry := range data.Registries {
			idToRegistry[registry.Id] = registry
		}
	}
	return idToRegistry
}


func CalculateNumberOfImages(nameToImage map[string]model.Image) int {
	return len(nameToImage)
}

func CalculateNumberOfTests(idToTest map[string]model.Test) int {
	return len(idToTest)
}

func CalculateNumberOfProjects(idToProject map[string]model.Project) int {
	return len(idToProject)
}

func CalculateAverageImagePerProject(nameToImage map[string]model.Image, idToProject map[string]model.Project) float64 {
	return float64(CalculateNumberOfImages(nameToImage))/float64(CalculateNumberOfProjects(idToProject))
}

func CalculateAverageTestPerProject(idToTest map[string]model.Test, idToProject map[string]model.Project) float64 {
	return float64(CalculateNumberOfTests(idToTest))/float64(CalculateNumberOfProjects(idToProject))
}

func CalculateAllProjectsOutcomeRates(idToProject map[string]model.Project) (float64, float64) {
	projectsSuccess := 0
	projectsFailure := 0

	for _, project := range idToProject {
		if project.Status == SUCCESS {
			projectsSuccess ++
		}
		if project.Status == FAILURE {
			projectsFailure ++
		}
	}

	return float64(projectsSuccess * 100)/float64(projectsSuccess+projectsFailure), float64(projectsFailure * 100)/float64(projectsSuccess+projectsFailure)
}

func CalculatePerProjectOutcomeRates(idToBuild map[string]model.Build, idToProject map[string]model.Project) (map[string]float64, map[string]float64){
	projectsSuccess := map[string]float64{}
	projectsFailure := map[string]float64{}

	for _, build := range idToBuild {
		if build.Status.Status == SUCCESS {
			projectsSuccess[build.ProjectId]++
		} else if build.Status.Status == FAILURE {
			projectsFailure[build.ProjectId]++
		}
	}

	for id := range idToProject {
		projectsSuccess[id] = float64(projectsSuccess[id] * 100)/float64(projectsSuccess[id]+projectsFailure[id])
		projectsFailure[id] = float64(projectsFailure[id] * 100)/float64(projectsSuccess[id]+projectsFailure[id])
	}

	return projectsSuccess, projectsFailure
}

func CalculateMostExecutedProjects(idToBuild map[string]model.Build, idToProject map[string]model.Project) ([]model.Project, int) {
	mostExecutedProjects := map[string]int{}

	for _, build := range idToBuild {
		mostExecutedProjects[build.ProjectId]++
	}

	max := 0
	projects := []model.Project{}
	for id, occurrence := range mostExecutedProjects{
		if occurrence > max {
			projects = []model.Project{idToProject[id]}
			max = occurrence
		} else if occurrence == max {
			projects = append(projects, idToProject[id])
		}
	}

	return projects, max
}

func CalculateMostUsedImages(idToProject map[string]model.Project) model.PairList{
	mostUsedImages := map[string]int{}

	for _, project := range idToProject {
		for _, image := range project.Images {
			mostUsedImages[image.Name+":"+image.Tag]++
		}
	}

	return rankByValue(mostUsedImages)
}

func ShowImagesInProjects(idToProject map[string]model.Project) map[string][]model.Project {
	imagesInProjects := map[string][]model.Project{}

	for _, project := range idToProject{
		for _, image := range project.Images {
			imagesInProjects[image.Name+":"+image.Tag] = append(imagesInProjects[image.Name+":"+image.Tag], project)
		}
	}

	return imagesInProjects
}

func CalculateMostExecutedTests(idToBuild map[string]model.Build, idToTest map[string]model.Test) ([]model.Test, int){
	buildsToTest := map[string][]model.Build{}
	mostExecutedTestsNr := 0
	mostExecutedTests := []model.Test{}

	for _, build := range idToBuild {
		buildsToTest[build.TestId] = append(buildsToTest[build.TestId], build)
	}

	for _, test := range idToTest {
		buildListLength := len(buildsToTest[test.Id])
		if mostExecutedTestsNr < buildListLength {
			mostExecutedTestsNr = buildListLength
			mostExecutedTests = []model.Test{test}
		} else if mostExecutedTestsNr == buildListLength {
			mostExecutedTests = append(mostExecutedTests, test)
		}
	}

	return mostExecutedTests, mostExecutedTestsNr
}

func CalculateLeastExecutedTests(idToBuild map[string]model.Build, idToTest map[string]model.Test) ([]model.Test, int){
	buildsToTest := map[string][]model.Build{}
	leastExecutedTestsNr := math.MaxInt32
	leastExecutedTests := []model.Test{}

	for _, build := range idToBuild {
		buildsToTest[build.TestId] = append(buildsToTest[build.TestId], build)
	}

	for _, test := range idToTest {
		buildListLength := len(buildsToTest[test.Id])
		if leastExecutedTestsNr > buildListLength {
			leastExecutedTestsNr = buildListLength
			leastExecutedTests = []model.Test{test}
		} else if leastExecutedTestsNr == buildListLength {
			leastExecutedTests = append(leastExecutedTests, test)
		}
	}

	return leastExecutedTests, leastExecutedTestsNr
}

func CalculateNumberOfTestsInEachHour(idToBuild map[string]model.Build, idToTest map[string]model.Test) map[int]int {
	hourlyActivities := map[int]int{}

	for _, build := range idToBuild {
		//Get the build's date and today's date in different variables - for the comparision below
		datetime := strings.Split(build.StartTime, "T")
		date := strings.Split(datetime[0], "-")
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

		t := strings.Split(datetime[1], ":")
		hour := t[0]
		hr, err := strconv.Atoi(hour);
		if err != nil {
			log.Println(err)
		} else {
			//Number today's hourly activities
			if (year == now.Year() && time.Month(month) == now.Month() && day == now.Day()) {
				for id := range idToTest{
					if (build.TestId == id) {
						hourlyActivities[hr]++
					}
				}
			}
		}
	}

	return hourlyActivities
}

func CalculateBusiestHours(idToBuild map[string]model.Build, idToTest map[string]model.Test) []int {
	hourlyActivities := CalculateNumberOfTestsInEachHour(idToBuild, idToTest)

	busiestHours := []int{}
	max := hourlyActivities[0]
	for i := 1; i < 24; i++ {
		if hourlyActivities[i] > max {
			busiestHours = []int{i}
			max = hourlyActivities[i]
		} else if hourlyActivities[i] == max {
			busiestHours = append(busiestHours, i)
		}
	}

	return busiestHours
}

func CalculateNoOfVulnerabilitiesFound(dataList []model.CollectedData, idToBuild map[string]model.Build) model.NoOfVulnerabilitiesWithLinksList {
	imageToNoOfVulnerability := map[string]int{}
	imageToReport := map[string]string{}

	for _, data := range dataList {
		for _, results := range data.Results {
			for _, entry := range results.ResultEntries {
				if strings.Contains(entry, "SUCCESS:") || strings.Contains(entry, "FAILURE:") {
					re := regexp.MustCompile("[0-9]+")
					vulnerabilitiesAux := re.FindAllString(entry, -1)
					noOfVulnerabilities, err := strconv.Atoi(vulnerabilitiesAux[0])
					reImage := regexp.MustCompile("[^ ]+:[^ ]+")
					registryAndImage := reImage.FindAllString(entry, -1)
					if err != nil {
						log.Println(err)
					} else {
						imageToNoOfVulnerability[strings.Join(registryAndImage, "")] = noOfVulnerabilities
						imageToReport[strings.Join(registryAndImage, "")] = strings.Join([]string{"/projects/", idToBuild[results.BuildId].ProjectId, "/tests/", idToBuild[results.BuildId].TestId, "/results", results.TargetArtifact.Artifact.ImageId}, "")
					}
				}
			}
		}
	}

	sortedImageToNoOfVulnerability := rankByValue(imageToNoOfVulnerability)
	vulnerabilities := make(model.NoOfVulnerabilitiesWithLinksList, len(sortedImageToNoOfVulnerability))

	i := 0
	for _, pair := range sortedImageToNoOfVulnerability {
		vulnerabilities[i] = model.NoOfVulnerabilitiesWithLinks{pair.Key, model.Pair{imageToReport[pair.Key], pair.Value}}
		i++
	}

	return vulnerabilities
}

//As seen in http://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values

func rankByValue(stringToInt map[string]int) model.PairList{
	pl := make(model.PairList, len(stringToInt))
	i := 0
	for k, v := range stringToInt {
		pl[i] = model.Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func ShowImagesInRegistries(nameToImage map[string]model.Image, idToRegistry map[string]model.Registry) map[string][]string {
	imagesInRegistries := map[string][]string{}

	for name, image := range nameToImage {
		if (image.Location == "Public Registry") {
			imagesInRegistries["Public Registry"] = append(imagesInRegistries["Public Registry"], name)
		} else {
			imagesInRegistries[idToRegistry[image.RegistryId].Name+"("+idToRegistry[image.RegistryId].Addr+")"] = append(imagesInRegistries[image.RegistryId], name)
		}
	}

	return imagesInRegistries
}