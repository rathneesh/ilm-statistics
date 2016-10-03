package util

import (
	"github.com/ilm-statistics/ilm-statistics/model"
)

func DiffImageList(oldImgs, newImgs []model.Image) ([]model.Image, []model.Image) {

	addedImages := []model.Image{}
	deletedImages := []model.Image{}

	for _, img1 := range newImgs {
		found := false
		for _, img2 := range oldImgs {
			if CmpImages(img1, img2) {
				found = true
				break
			}
		}
		if !found {
			addedImages = append(addedImages, img1)
		}
	}

	for _, img1 := range oldImgs {
		found := false
		for _, img2 := range newImgs {
			if CmpImages(img1, img2) {
				found = true
				break
			}
		}
		if !found {
			deletedImages = append(deletedImages, img1)
		}
	}

	return addedImages, deletedImages
}

func DiffProjectList(oldProjects, newProjects []model.Project) ([]model.Project, []model.Project) {

	addedProjects := []model.Project{}
	deletedProjects := []model.Project{}

	for _, proj1 := range newProjects {
		found := false
		for _, proj2 := range oldProjects {
			if CmpProjects(proj1, proj2) {
				found = true
				break
			}
		}
		if !found {
			addedProjects = append(addedProjects, proj1)
		}
	}

	for _, proj1 := range oldProjects {
		found := false
		for _, proj2 := range newProjects {
			if CmpProjects(proj1, proj2) {
				found = true
				break
			}
		}
		if !found {
			deletedProjects = append(deletedProjects, proj1)
		}
	}

	return addedProjects, deletedProjects
}

func DiffBuildList(oldBuilds, newBuilds []model.Build) ([]model.Build, []model.Build) {

	addedBuilds := []model.Build{}
	deletedBuilds := []model.Build{}

	for _, build1 := range newBuilds {
		found := false
		for _, build2 := range oldBuilds {
			if CmpBuilds(build1, build2) {
				found = true
				break
			}
		}
		if !found {
			addedBuilds = append(addedBuilds, build1)
		}
	}

	for _, build1 := range oldBuilds {
		found := false
		for _, build2 := range newBuilds {
			if CmpBuilds(build1, build2) {
				found = true
				break
			}
		}
		if !found {
			deletedBuilds = append(deletedBuilds, build1)
		}
	}

	return addedBuilds, deletedBuilds
}

func DiffRegistryList(oldRegistries, newRegistries []model.Registry) ([]model.Registry, []model.Registry) {

	addedRegistries := []model.Registry{}
	deletedRegistries := []model.Registry{}

	for _, reg1 := range newRegistries {
		found := false
		for _, reg2 := range oldRegistries {
			if CmpRegistries(reg1, reg2) {
				found = true
				break
			}
		}
		if !found {
			addedRegistries = append(addedRegistries, reg1)
		}
	}

	for _, reg1 := range oldRegistries {
		found := false
		for _, reg2 := range newRegistries {
			if CmpRegistries(reg1, reg2) {
				found = true
				break
			}
		}
		if !found {
			deletedRegistries = append(deletedRegistries, reg1)
		}
	}

	return addedRegistries, deletedRegistries
}

func DiffTestList(oldTests, newTests []model.Test) ([]model.Test, []model.Test) {

	addedTests := []model.Test{}
	deletedTests := []model.Test{}

	for _, test1 := range newTests {
		found := false
		for _, test2 := range oldTests {
			if CmpTests(test1, test2) {
				found = true
				break
			}
		}
		if !found {
			addedTests = append(addedTests, test1)
		}
	}

	for _, test1 := range oldTests {
		found := false
		for _, test2 := range newTests {
			if CmpTests(test1, test2) {
				found = true
				break
			}
		}
		if !found {
			deletedTests = append(deletedTests, test1)
		}
	}

	return addedTests, deletedTests
}

func DiffResultList(oldResults, newResults []model.BuildResult) ([]model.BuildResult, []model.BuildResult) {

	addedResults := []model.BuildResult{}
	deletedResults := []model.BuildResult{}

	for _, res1 := range newResults {
		found := false
		for _, res2 := range oldResults {
			if CmpBuildResults(res1, res2) {
				found = true
				break
			}
		}
		if !found {
			addedResults = append(addedResults, res1)
		}
	}

	for _, res1 := range oldResults{
		found := false
		for _, res2 := range newResults {
			if CmpBuildResults(res1, res2) {
				found = true
				break
			}
		}
		if !found {
			deletedResults = append(deletedResults, res1)
		}
	}

	return addedResults, deletedResults
}

func DiffRepositoryList(oldRepositories, newRepositories []model.Repository) ([]model.Repository, []model.Repository) {

	addedRepositories := []model.Repository{}
	deletedRepositories := []model.Repository{}

	for _, repo1 := range newRepositories {
		found := false
		for _, repo2 := range oldRepositories {
			if CmpRepositories(repo1, repo2) {
				found = true
				break
			}
		}
		if !found {
			addedRepositories = append(addedRepositories, repo1)
		}
	}

	for _, repo1 := range oldRepositories {
		found := false
		for _, repo2 := range newRepositories {
			if CmpRepositories(repo1, repo2) {
				found = true
				break
			}
		}
		if !found {
			deletedRepositories = append(deletedRepositories, repo1)
		}
	}

	return addedRepositories, deletedRepositories
}

func DiffCollectedData(oldData, newData model.CollectedData) model.CollectedDataDiff {
	difference := model.CollectedDataDiff{}
	difference.MAC = newData.MAC

	if newData.Username != oldData.Username {
		difference.NewUserName = newData.Username
	}
	difference.AddedImages, difference.DeletedImages = DiffImageList(oldData.Images, newData.Images)
	difference.AddedProjects, difference.DeletedProjects = DiffProjectList(oldData.Projects, newData.Projects)
	difference.AddedBuilds, difference.DeletedBuilds = DiffBuildList(oldData.Builds, newData.Builds)
	difference.AddedRegistries, difference.DeletedRegistries = DiffRegistryList(oldData.Registries, newData.Registries)
	difference.AddedTests, difference.DeletedTests = DiffTestList(oldData.Tests, newData.Tests)
	difference.AddedResults, difference.DeletedResults = DiffResultList(oldData.Results, newData.Results)
	if newData.Day != oldData.Day {
		difference.NewDay = newData.Day
	}

	return difference
}