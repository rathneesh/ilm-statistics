package util

import (
	"github.com/ilm-statistics/ilm-statistics/model"
)

// Compare two images
func CmpImages(img1, img2 model.Image) bool {
	if img1.Name != img2.Name {
		return false
	}
	if img1.ImageId != img2.ImageId {
		return false
	}
	if img1.Description != img2.Description {
		return false
	}
	if img1.RegistryId != img2.RegistryId {
		return false
	}
	if img1.Tag != img2.Tag {
		return false
	}
	if len(img1.IlmTags) != len(img2.IlmTags) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(img1.IlmTags); i++ {
			for j := 0; j < len(img2.IlmTags); j++ {
				if img1.IlmTags[i] == img2.IlmTags[j] {
					similar ++
				}
			}
		}
		if similar != len(img1.IlmTags) {
			return false
		}
	}
	if img1.Location != img2.Location {
		return false
	}

	if img1.SkipImageBuild != img2.SkipImageBuild {
		return false
	}

	return true
}

//Compare two providers
func CmpProviders(prov1, prov2 model.Provider) bool {
	return prov1 == prov2
}

//Compare two tests
func CmpTests(test1, test2 model.Test) bool {
	return test1 == test2
}

//Compare two projects
func CmpProjects(proj1, proj2 model.Project) bool {
	if proj1.Id != proj2.Id {
		return false
	}
	if proj1.Name != proj2.Name {
		return false
	}
	if proj1.CreationTime != proj2.CreationTime {
		return false
	}
	if proj1.Status != proj2.Status {
		return false
	}

	if len(proj1.Images) != len(proj2.Images) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(proj1.Images); i++ {
			for j := 0; j < len(proj2.Images); j++ {
				if CmpImages(proj1.Images[i], proj2.Images[j]) {
					similar ++
				}
			}
		}
		if similar != len(proj1.Images) {
			return false
		}
	}

	if len(proj1.Tests) != len(proj2.Tests) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(proj1.Tests); i++ {
			for j := 0; j < len(proj2.Tests); j++ {
				if CmpTests(proj1.Tests[i], proj2.Tests[j]) {
					similar ++
				}
			}
		}
		if similar != len(proj1.Tests) {
			return false
		}
	}

	return true
}

// Compare two builds
func CmpBuilds(build1, build2 model.Build) bool{
	return build1 == build2
}

//Compare two Registries
func CmpRegistries(reg1, reg2 model.Registry) bool {
	return reg1 == reg2
}

//Compre two BuildResults
func CmpBuildResults(res1, res2 model.BuildResult) bool {
	if res1.Id != res2.Id{
		return false
	}
	if res1.BuildId != res2.BuildId {
		return false
	}
	if len(res1.ResultEntries) != len(res2.ResultEntries) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(res1.ResultEntries); i++ {
			for j := 0; j < len(res2.ResultEntries); j++ {
				if res1.ResultEntries[i] == res2.ResultEntries[j] {
					similar ++
				}
			}
		}
		if similar != len(res1.ResultEntries) {
			return false
		}
	}
	return true
}

//Compare two repositories
func CmpRepositories(repo1, repo2 model.Repository) bool{
	if repo1.Name != repo2.Name {
		return false
	}
	if repo1.Tag != repo2.Tag {
		return false
	}
	if len(repo1.FsLayers) != len(repo2.FsLayers) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(repo1.FsLayers); i++ {
			for j := 0; j < len(repo2.FsLayers); j++ {
				if repo1.FsLayers[i] == repo2.FsLayers[j] {
					similar ++
				}
			}
		}
		if similar != len(repo1.FsLayers) {
			return false
		}
	}

	if len(repo1.Signatures) != len(repo2.Signatures) {
		return false
	} else {
		similar := 0
		for i := 0; i < len(repo1.Signatures); i++ {
			for j := 0; j < len(repo2.Signatures); j++ {
				if repo1.Signatures[i] == repo2.Signatures[j] {
					similar ++
				}
			}
		}
		if similar != len(repo1.Signatures) {
			return false
		}
	}

	if repo1.HasProblems != repo2.HasProblems {
		return false;
	}
	if repo1.Message != repo2.Message {
		return false;
	}
	if repo1.RegistryUrl != repo2.RegistryUrl {
		return false;
	}
	if repo1.RegistryName != repo2.RegistryName {
		return false;
	}
	return true
}

// Compare two collected data entities
func CmpCollectedData(data1 model.CollectedData, data2 model.CollectedData) bool {

	addedImages, deletedImages := DiffImageList(data1.Images, data2.Images)
	if !(len(addedImages) == 0 && len(deletedImages) == 0) {
		return false
	}

	addedProjects, deletedProjects := DiffProjectList(data1.Projects, data2.Projects)
	if !(len(addedProjects) == 0 && len(deletedProjects) == 0) {
		return false
	}

	addedBuilds, deletedBuilds := DiffBuildList(data1.Builds, data2.Builds)
	if !(len(addedBuilds) == 0 && len(deletedBuilds) == 0) {
		return false
	}

	addedRegistries, deletedRegistries := DiffRegistryList(data1.Registries, data2.Registries)
	if !(len(addedRegistries) == 0 && len(deletedRegistries) == 0) {
		return false
	}

	addedTests, deletedTests := DiffTestList(data1.Tests, data2.Tests)
	if !(len(addedTests) == 0 && len(deletedTests) == 0) {
		return false
	}

	addedResults, deletedResults := DiffResultList(data1.Results, data2.Results)
	if !(len(addedResults) == 0 && len(deletedResults) == 0) {
		return false
	}

	return (data1.Ip == data2.Ip && data1.MAC == data2.MAC && data1.Username == data2.Username && data1.Day.Equal(data2.Day))
}