package util

import (
	"github.com/ilm-statistics/ilm-statistics/model"
	"log"
)

func MergeDiff(oldData model.CollectedData, newData model.CollectedDataDiff) model.CollectedData {

	data := model.CollectedData{}
	if oldData.MAC != newData.MAC {
		log.Println("Data from different machines, not possible to merge the differences")
		return model.CollectedData{}
	}
	data.MAC = newData.MAC
	data.Ip = newData.Ip
	if newData.NewUserName != "" {
		data.Username = newData.NewUserName
	} else {
		data.Username = oldData.Username
	}
	data.Images = MergeImageLists(oldData.Images, newData.AddedImages)
	data.Projects = MergeProjectLists(oldData.Projects, newData.AddedProjects)
	data.Builds = MergeBuildLists(oldData.Builds, newData.AddedBuilds)
	data.Registries = MergeRegistryLists(oldData.Registries, newData.AddedRegistries)
	data.Tests = MergeTestLists(oldData.Tests, newData.AddedTests)
	data.Results = MergeResultLists(oldData.Results, newData.AddedResults)
	data.Day = newData.NewDay

	return data
}

func MergeImageLists(starterImages []model.Image, addedImages []model.Image) []model.Image {

	for _, image := range addedImages {
		starterImages = appendIfMissingImage(starterImages, image)
	}

	return starterImages
}

func appendIfMissingImage(imageList []model.Image, image model.Image) []model.Image{
	for _, e := range imageList {
		if CmpImages(e, image) {
			return imageList
		}
	}
	return append(imageList, image)
}

func MergeProjectLists(starterProjects []model.Project, addedProjects []model.Project) []model.Project {
	for _, proj := range addedProjects {
		starterProjects = appendIfMissingProject(starterProjects, proj)
	}

	return starterProjects
}

func appendIfMissingProject(projectList []model.Project, project model.Project) []model.Project{
	for _, e := range projectList {
		if CmpProjects(e, project) {
			return projectList
		}
	}
	return append(projectList, project)
}

func MergeBuildLists(starterBuilds []model.Build, addedBuilds []model.Build) []model.Build {
	for _, bld := range addedBuilds {
		starterBuilds = appendIfMissingBuild(starterBuilds, bld)
	}

	return starterBuilds
}

func appendIfMissingBuild(buildList []model.Build, build model.Build) []model.Build{
	for _, e := range buildList {
		if CmpBuilds(e, build) {
			return buildList
		}
	}
	return append(buildList, build)
}

func MergeRegistryLists(starterRegistries []model.Registry, addedRegistries []model.Registry) []model.Registry {
	for _, reg := range addedRegistries {
		starterRegistries = appendIfMissingRegistry(starterRegistries, reg)
	}

	return starterRegistries
}

func appendIfMissingRegistry(registryList []model.Registry, registry model.Registry) []model.Registry{
	for _, e := range registryList {
		if CmpRegistries(e, registry) {
			return registryList
		}
	}
	return append(registryList, registry)
}

func MergeTestLists(starterTests []model.Test, addedTests []model.Test) []model.Test {
	for _, test := range addedTests {
		starterTests = appendIfMissingTest(starterTests, test)
	}

	return starterTests
}

func appendIfMissingTest(testList []model.Test, test model.Test) []model.Test{
	for _, e := range testList {
		if CmpTests(e, test) {
			return testList
		}
	}
	return append(testList, test)
}

func MergeResultLists(starterResults []model.BuildResult, addedResults []model.BuildResult) []model.BuildResult {
	for _, res := range addedResults {
		starterResults = appendIfMissingResult(starterResults, res)
	}

	return starterResults
}

func appendIfMissingResult(resultList []model.BuildResult, result model.BuildResult) []model.BuildResult{
	for _, e := range resultList {
		if CmpBuildResults(e, result) {
			return resultList
		}
	}
	return append(resultList, result)
}
