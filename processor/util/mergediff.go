package util

import (
	"github.com/ilm-statistics/ilm-statistics/model"
	"log"
)

func MergeDiff(oldData model.CollectedData, newData model.CollectedDataDiff) model.CollectedData {

	data := model.CollectedData{}
	if oldData.MAC != newData.MAC {
		log.Panicln("Data from different machines, not possible to merge the differences")
		return model.CollectedData{}
	}
	data.MAC = newData.MAC
	if newData.NewUserName != "" {
		data.Username = newData.NewUserName
	} else {
		data.Username = oldData.Username
	}
	data.Images = MergeImageLists(oldData.Images, newData.AddedImages, newData.DeletedImages)
	data.Projects = MergeProjectLists(oldData.Projects, newData.AddedProjects, newData.DeletedProjects)
	data.Builds = MergeBuildLists(oldData.Builds, newData.AddedBuilds, newData.DeletedBuilds)
	data.Registries = MergeRegistryLists(oldData.Registries, newData.AddedRegistries, newData.DeletedRegistries)
	data.Tests = MergeTestLists(oldData.Tests, newData.AddedTests, newData.DeletedTests)
	data.Results = MergeResultLists(oldData.Results, newData.AddedResults, newData.DeletedResults)
	data.Day = newData.NewDay

	return data
}

func MergeImageLists(starterImages []model.Image, addedImages []model.Image, deletedImages []model.Image) []model.Image {
	result := []model.Image{}

	isDeleted := map[string]model.Image{}

	for _, img := range deletedImages {
		isDeleted[img.Id] = img
	}

	for _, img := range starterImages {
		if !CmpImages(isDeleted[img.Id], img) {
			result = append(result, img)
		}
	}

	for _, img := range addedImages {
		result = append(result, img)
	}

	return result
}

func MergeProjectLists(starterProjects []model.Project, addedProjects []model.Project, deletedProjects []model.Project) []model.Project {
	result := []model.Project{}

	isDeleted := map[string]model.Project{}

	for _, proj := range deletedProjects {
		isDeleted[proj.Id] = proj
	}

	for _, proj := range starterProjects {
		if !CmpProjects(isDeleted[proj.Id], proj) {
			result = append(result, proj)
		}
	}

	for _, proj := range addedProjects {
		result = append(result, proj)
	}

	return result
}

func MergeBuildLists(starterBuilds []model.Build, addedBuilds []model.Build, deletedBuilds []model.Build) []model.Build {
	result := []model.Build{}

	isDeleted := map[string]model.Build{}

	for _, bld := range deletedBuilds {
		isDeleted[bld.Id] = bld
	}

	for _, bld := range starterBuilds {
		if !CmpBuilds(isDeleted[bld.Id], bld) {
			result = append(result, bld)
		}
	}

	for _, bld := range addedBuilds {
		result = append(result, bld)
	}

	return result
}

func MergeRegistryLists(starterRegistries []model.Registry, addedRegistries []model.Registry, deletedRegistries []model.Registry) []model.Registry {
	result := []model.Registry{}

	isDeleted := map[string]model.Registry{}

	for _, reg := range deletedRegistries {
		isDeleted[reg.Id] = reg
	}

	for _, reg := range starterRegistries{
		if !CmpRegistries(isDeleted[reg.Id],reg) {
			result = append(result, reg)
		}
	}

	for _, reg := range addedRegistries {
		result = append(result, reg)
	}

	return result
}

func MergeTestLists(starterTests []model.Test, addedTests []model.Test, deletedTests []model.Test) []model.Test {
	result := []model.Test{}

	isDeleted := map[string]model.Test{}

	for _, test := range deletedTests {
		isDeleted[test.Id] = test
	}

	for _, test := range starterTests {
		if !CmpTests(isDeleted[test.Id], test) {
			result = append(result, test)
		}
	}

	for _, test := range addedTests {
		result = append(result, test)
	}

	return result
}

func MergeResultLists(starterResults []model.BuildResult, addedResults []model.BuildResult, deletedResults []model.BuildResult) []model.BuildResult {
	result := []model.BuildResult{}

	isDeleted := map[string]model.BuildResult{}

	for _, res := range deletedResults {
		isDeleted[res.Id] = res
	}

	for _, res := range starterResults {
		if !CmpBuildResults(isDeleted[res.Id], res) {
			result = append(result, res)
		}
	}

	for _, res := range addedResults {
		result = append(result, res)
	}

	return result
}

func MergeRepositoryLists(starterRepositories []model.Repository, addedRepositories []model.Repository, deletedRepositories []model.Repository) []model.Repository {
	result := []model.Repository{}

	isDeleted := map[string]model.Repository{}

	for _, repo := range deletedRepositories {
		isDeleted[repo.Name] = repo
	}

	for _, repo := range starterRepositories {
		if !CmpRepositories(isDeleted[repo.Name], repo) {
			result = append(result, repo)
		}
	}

	for _, repo := range addedRepositories {
		result = append(result, repo)
	}

	return result
}