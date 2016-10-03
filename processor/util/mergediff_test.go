package util

import (
	"testing"
	"github.com/ilm-statistics/ilm-statistics/model"
	"time"
)

func TestMergeImageLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Image{}
	addedImages := []model.Image{}
	deletedImages := []model.Image{}

	if len(MergeImageLists(initList, addedImages, deletedImages)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added image

	addedImages = []model.Image{{Id: "1", Name: "image1"}}
	add, del := DiffImageList(MergeImageLists(initList, addedImages, deletedImages), addedImages)
	if len(MergeImageLists(initList, addedImages, deletedImages)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted image

	initList = MergeImageLists(initList, addedImages, deletedImages)
	addedImages = []model.Image{}
	deletedImages = []model.Image{{Id: "1", Name: "image1"}}

	add, del = DiffImageList(MergeImageLists(initList, addedImages, deletedImages), []model.Image{})
	if len(MergeImageLists(initList, addedImages, deletedImages)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Deleting a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 image, add another 1

	initList = []model.Image{{Id: "1", Name: "image1"}}
	deletedImages = []model.Image{{Id: "1", Name: "image1"}}
	addedImages = []model.Image{{Id: "1", Name: "image1"}}
	add, del = DiffImageList(MergeImageLists(initList, addedImages, deletedImages), []model.Image{{Id: "1", Name: "image1"}})

	if len(MergeImageLists(initList, addedImages, deletedImages)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding and deleting a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeProjectLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Project{}
	addedProjects := []model.Project{}
	deletedProjects := []model.Project{}

	if len(MergeProjectLists(initList, addedProjects, deletedProjects)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added project

	addedProjects = []model.Project{{Id: "1"}}
	add, del := DiffProjectList(MergeProjectLists(initList, addedProjects, deletedProjects), addedProjects)
	if len(MergeProjectLists(initList, addedProjects, deletedProjects)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted project

	initList = MergeProjectLists(initList, addedProjects, deletedProjects)
	addedProjects = []model.Project{}
	deletedProjects = []model.Project{{Id: "1"}}

	add, del = DiffProjectList(MergeProjectLists(initList, addedProjects, deletedProjects), []model.Project{})
	if len(MergeProjectLists(initList, addedProjects, deletedProjects)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 project, add another 1

	initList = []model.Project{{Id: "1"}}
	deletedProjects = []model.Project{{Id: "1"}}
	addedProjects = []model.Project{{Id: "1"}}
	add, del = DiffProjectList(MergeProjectLists(initList, addedProjects, deletedProjects), []model.Project{{Id: "1"}})

	if len(MergeProjectLists(initList, addedProjects, deletedProjects)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeBuildLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Build{}
	addedBuilds := []model.Build{}
	deletedBuilds := []model.Build{}

	if len(MergeBuildLists(initList, addedBuilds, deletedBuilds)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added build

	addedBuilds = []model.Build{{ Id: "1"}}
	add, del := DiffBuildList(MergeBuildLists(initList, addedBuilds, deletedBuilds), addedBuilds)
	if len(MergeBuildLists(initList, addedBuilds, deletedBuilds)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted build

	initList = MergeBuildLists(initList, addedBuilds, deletedBuilds)
	addedBuilds = []model.Build{}
	deletedBuilds = []model.Build{{Id: "1"}}

	add, del = DiffBuildList(MergeBuildLists(initList, addedBuilds, deletedBuilds), []model.Build{})
	if len(MergeBuildLists(initList, addedBuilds, deletedBuilds)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 build, add another 1

	initList = []model.Build{{Id: "1"}}
	deletedBuilds = []model.Build{{Id: "1"}}
	addedBuilds = []model.Build{{Id: "1"}}
	add, del = DiffBuildList(MergeBuildLists(initList, addedBuilds, deletedBuilds), []model.Build{{Id: "1"}})

	if len(MergeBuildLists(initList, addedBuilds, deletedBuilds)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeRegistryLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Registry{}
	addedRegistries := []model.Registry{}
	deletedRegistries := []model.Registry{}

	if len(MergeRegistryLists(initList, addedRegistries, deletedRegistries)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added registry

	addedRegistries = []model.Registry{{Id: "1"}}
	add, del := DiffRegistryList(MergeRegistryLists(initList, addedRegistries, deletedRegistries), addedRegistries)
	if len(MergeRegistryLists(initList, addedRegistries, deletedRegistries)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted registry

	initList = MergeRegistryLists(initList, addedRegistries, deletedRegistries)
	addedRegistries = []model.Registry{}
	deletedRegistries = []model.Registry{{Id: "1"}}

	add, del = DiffRegistryList(MergeRegistryLists(initList, addedRegistries, deletedRegistries), []model.Registry{})
	if len(MergeRegistryLists(initList, addedRegistries, deletedRegistries)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 registry, add another 1

	initList = []model.Registry{{Id: "1"}}
	deletedRegistries = []model.Registry{{Id: "1"}}
	addedRegistries = []model.Registry{{Id: "1"}}
	add, del = DiffRegistryList(MergeRegistryLists(initList, addedRegistries, deletedRegistries), []model.Registry{{Id: "1"}})

	if len(MergeRegistryLists(initList, addedRegistries, deletedRegistries)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeTestLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Test{}
	addedTests := []model.Test{}
	deletedTests := []model.Test{}

	if len(MergeTestLists(initList, addedTests, deletedTests)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added test

	addedTests = []model.Test{{Id: "1"}}
	add, del := DiffTestList(MergeTestLists(initList, addedTests, deletedTests), addedTests)
	if len(MergeTestLists(initList, addedTests, deletedTests)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted test

	initList = MergeTestLists(initList, addedTests, deletedTests)
	addedTests = []model.Test{}
	deletedTests = []model.Test{{Id: "1"}}

	add, del = DiffTestList(MergeTestLists(initList, addedTests, deletedTests), []model.Test{})
	if len(MergeTestLists(initList, addedTests, deletedTests)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 test, add another 1

	initList = []model.Test{{Id: "1"}}
	deletedTests = []model.Test{{Id: "1"}}
	addedTests = []model.Test{{Id: "1"}}
	add, del = DiffTestList(MergeTestLists(initList, addedTests, deletedTests), []model.Test{{Id: "1"}})

	if len(MergeTestLists(initList, addedTests, deletedTests)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeResultLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.BuildResult{}
	addedResults := []model.BuildResult{}
	deletedResults := []model.BuildResult{}

	if len(MergeResultLists(initList, addedResults, deletedResults)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added build result

	addedResults = []model.BuildResult{{Id: "1"}}
	add, del := DiffResultList(MergeResultLists(initList, addedResults, deletedResults), addedResults)
	if len(MergeResultLists(initList, addedResults, deletedResults)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted build result

	initList = MergeResultLists(initList, addedResults, deletedResults)
	addedResults = []model.BuildResult{}
	deletedResults = []model.BuildResult{{Id: "1"}}

	add, del = DiffResultList(MergeResultLists(initList, addedResults, deletedResults), []model.BuildResult{})
	if len(MergeResultLists(initList, addedResults, deletedResults)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 build result, add another 1

	initList = []model.BuildResult{{Id: "1"}}
	deletedResults = []model.BuildResult{{Id: "1"}}
	addedResults = []model.BuildResult{{Id: "1"}}
	add, del = DiffResultList(MergeResultLists(initList, addedResults, deletedResults), []model.BuildResult{{Id: "1"}})

	if len(MergeResultLists(initList, addedResults, deletedResults)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeRepositoryLists(t *testing.T) {

	// Merge 3 empty lists

	initList := []model.Repository{}
	addedRepositories := []model.Repository{}
	deletedRepositories := []model.Repository{}

	if len(MergeRepositoryLists(initList, addedRepositories, deletedRepositories)) != 0 {
		t.Error("Merging three empty list resulted in a non-empty list")
	}

	// Merge 1 added repository

	addedRepositories = []model.Repository{{Name: "repo1"}}
	add, del := DiffRepositoryList(MergeRepositoryLists(initList, addedRepositories, deletedRepositories), addedRepositories)
	if len(MergeRepositoryLists(initList, addedRepositories, deletedRepositories)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}


	// Merge 1 deleted repository

	initList = MergeRepositoryLists(initList, addedRepositories, deletedRepositories)
	addedRepositories = []model.Repository{}
	deletedRepositories = []model.Repository{{Name: "repo1"}}

	add, del = DiffRepositoryList(MergeRepositoryLists(initList, addedRepositories, deletedRepositories), []model.Repository{})
	if len(MergeRepositoryLists(initList, addedRepositories, deletedRepositories)) != 0 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Delete 1 repository, add another 1

	initList = []model.Repository{{Name: "repo1"}}
	deletedRepositories = []model.Repository{{Name: "repo1"}}
	addedRepositories = []model.Repository{{Name: "repo1"}}
	add, del = DiffRepositoryList(MergeRepositoryLists(initList, addedRepositories, deletedRepositories), []model.Repository{{Name: "repo1"}})

	if len(MergeRepositoryLists(initList, addedRepositories, deletedRepositories)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}
}

func TestMergeDiff(t *testing.T) {

	// Merge 2 empty objects

	initData := model.CollectedData{}
	finalData := model.CollectedDataDiff{}

	dataDiff := MergeDiff(initData, finalData)

	if dataDiff.MAC != "" {
		t.Error("Empty MACs give error when merging")
	}

	if len(dataDiff.Images) != 0 {
		t.Error("Empty image lists' merging does not result in empty list")
	}

	if len(dataDiff.Projects) != 0 {
		t.Error("Empty project lists' merging does not result in empty list")
	}

	if len(dataDiff.Builds) != 0 {
		t.Error("Empty build lists' merging does not result in empty list")
	}

	if len(dataDiff.Registries) != 0 {
		t.Error("Empty registry lists' merging does not result in empty list")
	}

	if len(dataDiff.Tests) != 0 {
		t.Error("Empty test lists' merging does not result in empty list")
	}

	if len(dataDiff.Results) != 0 {
		t.Error("Empty build result lists' merging does not result in empty list")
	}

	if !dataDiff.Day.IsZero() {
		t.Error("Null time value merging results in non-null value")
	}

	// Merge 2 non-empty objects

	initData = model.CollectedData {
		MAC: "co:mp:ut:er",
		Username: "uname",
		Images: []model.Image{{
			Id: "1",
			Name: "Image1",
			ImageId: "1",
			Description: "A cool new image with all that fancy stuff",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		}, {
			Id: "2",
			Name: "Image2",
			ImageId: "2",
			Description: "A cool new image with all that fancy stuff",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		}},
		Projects: []model.Project{{
			Id: "1",
			Name: "Project1",
			CreationTime: "2002",
			Status: "new",
			Images: []model.Image{},
			Tests: []model.Test{},
		},
			{
				Id: "2",
				Name: "Project2",
				CreationTime: "2002",
				Status: "new",
				Images: []model.Image{},
				Tests: []model.Test{},
			}},
		Builds: []model.Build{{
			Id: "1",
			ProjectId: "1",
			TestId: "1",
			StartTime: "2002",
			Status: model.Status{Status:"new"},
		},
			{
				Id: "2",
				ProjectId: "2",
				TestId: "2",
				StartTime: "2002",
				Status: model.Status{Status:"new"},
			}},
		Registries: []model.Registry{{
			Id: "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
			{
				Id: "2",
				Name: "Private2",
				Addr: "localhost:5002",
			}},
		Tests: []model.Test{{
			Id: "1",
			Provider: model.Provider{
				ProviderType: "clair",
			},
		},
			{
				Id: "2",
				Provider: model.Provider{
					ProviderType: "clair",
				},
			}},
		Results: []model.BuildResult{{
			Id: "1",
			BuildId: "1",
			ResultEntries: []string{"1","2"},
			},
			{
				Id: "2",
				BuildId: "2",
				ResultEntries: []string{"1","2"},
			}},
		Day: time.Now(),
	}

	finalData = model.CollectedDataDiff {
		MAC: "co:mp:ut:er",
		NewUserName: "uname",
		AddedImages: []model.Image{},
		DeletedImages: []model.Image{{
			Id: "1",
			Name: "Image1",
			ImageId: "1",
			Description: "A cool new image with all that fancy stuff",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		}, {
			Id: "2",
			Name: "Image2",
			ImageId: "2",
			Description: "A cool new image with all that fancy stuff",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		}},
		AddedProjects: []model.Project{},
		DeletedProjects: []model.Project{{
			Id: "1",
			Name: "Project1",
			CreationTime: "2002",
			Status: "new",
			Images: []model.Image{},
			Tests: []model.Test{},
		},
			{
				Id: "2",
				Name: "Project2",
				CreationTime: "2002",
				Status: "new",
				Images: []model.Image{},
				Tests: []model.Test{},
			},},
		AddedBuilds: []model.Build{},
		DeletedBuilds: []model.Build{{
			Id: "1",
			ProjectId: "1",
			TestId: "1",
			StartTime: "2002",
			Status: model.Status{Status:"new"},
		},
			{
				Id: "2",
				ProjectId: "2",
				TestId: "2",
				StartTime: "2002",
				Status: model.Status{Status:"new"},
			}},
		AddedRegistries: []model.Registry{},
		DeletedRegistries: []model.Registry{{
			Id: "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
			{
				Id: "2",
				Name: "Private2",
				Addr: "localhost:5002",
			}},
		AddedTests: []model.Test{},
		DeletedTests: []model.Test{{
			Id: "1",
			Provider: model.Provider{
				ProviderType: "clair",
			},
		},
			{
				Id: "2",
				Provider: model.Provider{
					ProviderType: "clair",
				},
			}},
		AddedResults: []model.BuildResult{},
		DeletedResults: []model.BuildResult{{
			Id: "1",
			BuildId: "1",
			ResultEntries: []string{"1","2"},
		},
			{
				Id: "2",
				BuildId: "2",
				ResultEntries: []string{"1","2"},
			}},
		AddedRepositories: []model.Repository{},
		DeletedRepositories: []model.Repository{{
			Name: "Repo1",
			Tag: "private",
			FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
			Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
			HasProblems: false,
			Message: "msg",
			RegistryUrl: "localhost:5000",
			RegistryName: "Reggy",
		},
			{
				Name: "Repo2",
				Tag: "public",
				FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
				Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
				HasProblems: false,
				Message: "msg",
				RegistryUrl: "localhost:5000",
				RegistryName: "Repsy",
			}},
		NewDay: time.Now(),
	}


	dataDiff = MergeDiff(initData, finalData)

	if dataDiff.MAC != "co:mp:ut:er" {
		t.Error("MACs give an error when merging")
	}

	if len(dataDiff.Images) != 0 {
		t.Error("Image lists are not merging correctly")
	}

	if len(dataDiff.Projects) != 0 {
		t.Error("Project lists are not merging correctly")
	}

	if len(dataDiff.Builds) != 0 {
		t.Error("Build lists are not merging correctly")
	}

	if len(dataDiff.Registries) != 0 {
		t.Error("Registry lists are not merging  correctly")
	}

	if len(dataDiff.Tests) != 0 {
		t.Error("Test lists are not merging correctly")
	}

	if len(dataDiff.Results) != 0 {
		t.Error("Build result lists are not merging correctly")
	}

	finalData = model.CollectedDataDiff{}
	finalData.MAC = "co:mp:ut:er"

	dataDiff = MergeDiff(initData, finalData)

	if dataDiff.MAC != "co:mp:ut:er" {
		t.Error("MACs give an error when merging")
	}

	if len(dataDiff.Images) != 2 {
		t.Error("Image lists are not merging correctly")
	}

	if len(dataDiff.Projects) != 2 {
		t.Error("Project lists are not merging correctly")
	}

	if len(dataDiff.Builds) != 2 {
		t.Error("Build lists are not merging correctly")
	}

	if len(dataDiff.Registries) != 2 {
		t.Error("Registry lists are not merging  correctly")
	}

	if len(dataDiff.Tests) != 2 {
		t.Error("Test lists are not merging correctly")
	}

	if len(dataDiff.Results) != 2 {
		t.Error("Build result lists are not merging correctly")
	}
}