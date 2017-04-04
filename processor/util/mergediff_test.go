package util

import (
	"github.com/ilm-statistics/ilm-statistics/model"
	"testing"
	"time"
)

func TestMergeImageLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeImageLists: "

	// Merge 2 empty lists

	initList := []model.Image{}
	addedImages := []model.Image{}
	if len(MergeImageLists(initList, addedImages)) != 0 {
		t.Error(testedFuncName + "Merging two empty lists resulted in a non-empty list")
	}

	// Merge 1 added image

	addedImages = []model.Image{{Id: "1", Name: "image1"}}
	add, del := DiffImageList(MergeImageLists(initList, addedImages), addedImages)
	if len(MergeImageLists(initList, addedImages)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName + "Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present image

	initList = addedImages
	add, del = DiffImageList(MergeImageLists(initList, addedImages), addedImages)
	if len(MergeImageLists(initList, addedImages)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName + "Adding a the same element again to the initial list gave an unexpected result")
	}

}

func TestMergeProjectLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeProjectLists: "

	// Merge 2 empty lists

	initList := []model.Project{}
	addedProjects := []model.Project{}

	if len(MergeProjectLists(initList, addedProjects)) != 0 {
		t.Error(testedFuncName + "Merging two empty list resulted in a non-empty list")
	}

	// Merge 1 added project

	addedProjects = []model.Project{{Id: "1"}}
	add, del := DiffProjectList(MergeProjectLists(initList, addedProjects), addedProjects)
	if len(MergeProjectLists(initList, addedProjects)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present project

	initList = []model.Project{{Id: "1"}}
	addedProjects = []model.Project{{Id: "1"}}
	add, del = DiffProjectList(MergeProjectLists(initList, addedProjects), []model.Project{{Id: "1"}})

	if len(MergeProjectLists(initList, addedProjects)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a the same element again to the initial list gave an unexpected result")
	}
}

func TestMergeBuildLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeBuildLists: "

	// Merge 2 empty lists

	initList := []model.Build{}
	addedBuilds := []model.Build{}

	if len(MergeBuildLists(initList, addedBuilds)) != 0 {
		t.Error(testedFuncName, "Merging two empty list resulted in a non-empty list")
	}

	// Merge 1 added build

	addedBuilds = []model.Build{{Id: "1"}}
	add, del := DiffBuildList(MergeBuildLists(initList, addedBuilds), addedBuilds)
	if len(MergeBuildLists(initList, addedBuilds)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present build

	initList = []model.Build{{Id: "1"}}
	addedBuilds = []model.Build{{Id: "1"}}
	add, del = DiffBuildList(MergeBuildLists(initList, addedBuilds), []model.Build{{Id: "1"}})

	if len(MergeBuildLists(initList, addedBuilds)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a the same element again to the initial list gave an unexpected result")
	}
}

func TestMergeRegistryLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeRegistryLists: "

	// Merge 2 empty lists

	initList := []model.Registry{}
	addedRegistries := []model.Registry{}

	if len(MergeRegistryLists(initList, addedRegistries)) != 0 {
		t.Error(testedFuncName, "Merging two empty list resulted in a non-empty list")
	}

	// Merge 1 added registry

	addedRegistries = []model.Registry{{Id: "1"}}
	add, del := DiffRegistryList(MergeRegistryLists(initList, addedRegistries), addedRegistries)
	if len(MergeRegistryLists(initList, addedRegistries)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present registry

	initList = []model.Registry{{Id: "1"}}
	addedRegistries = []model.Registry{{Id: "1"}}
	add, del = DiffRegistryList(MergeRegistryLists(initList, addedRegistries), []model.Registry{{Id: "1"}})

	if len(MergeRegistryLists(initList, addedRegistries)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a the same element again to the initial list gave an unexpected result")
	}
}

func TestMergeTestLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeTestLists: "

	// Merge 2 empty lists

	initList := []model.Test{}
	addedTests := []model.Test{}

	if len(MergeTestLists(initList, addedTests)) != 0 {
		t.Error(testedFuncName, "Merging two empty list resulted in a non-empty list")
	}

	// Merge 1 added test

	addedTests = []model.Test{{Id: "1"}}
	add, del := DiffTestList(MergeTestLists(initList, addedTests), addedTests)
	if len(MergeTestLists(initList, addedTests)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present test

	initList = []model.Test{{Id: "1"}}
	addedTests = []model.Test{{Id: "1"}}
	add, del = DiffTestList(MergeTestLists(initList, addedTests), []model.Test{{Id: "1"}})

	if len(MergeTestLists(initList, addedTests)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a the same element again to the initial list gave an unexpected result")
	}
}

func TestMergeResultLists(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeResultLists: "

	// Merge 2 empty lists

	initList := []model.BuildResult{}
	addedResults := []model.BuildResult{}
	if len(MergeResultLists(initList, addedResults)) != 0 {
		t.Error("Merging two empty list resulted in a non-empty list")
	}

	// Merge 1 added build result

	addedResults = []model.BuildResult{{Id: "1"}}
	add, del := DiffResultList(MergeResultLists(initList, addedResults), addedResults)
	if len(MergeResultLists(initList, addedResults)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error("Adding a one-elemented list to the initial list gave an unexpected result")
	}

	// Merge already present result

	initList = []model.BuildResult{{Id: "1"}}
	addedResults = []model.BuildResult{{Id: "1"}}
	add, del = DiffResultList(MergeResultLists(initList, addedResults), []model.BuildResult{{Id: "1"}})

	if len(MergeResultLists(initList, addedResults)) != 1 || len(add) != 0 || len(del) != 0 {
		t.Error(testedFuncName, "Adding a the same element again to the initial list gave an unexpected result")
	}
}

func TestMergeDiff(t *testing.T) {
	t.Parallel()
	testedFuncName := "MergeDiff: "

	// Merge 2 empty objects

	initData := model.CollectedData{}
	finalData := model.CollectedDataDiff{}

	dataDiff := MergeDiff(initData, finalData)

	if dataDiff.MAC != "" {
		t.Error(testedFuncName, "Empty MACs give error when merging")
	}

	if len(dataDiff.Images) != 0 {
		t.Error(testedFuncName, "Empty image lists' merging does not result in empty list")
	}

	if len(dataDiff.Projects) != 0 {
		t.Error(testedFuncName, "Empty project lists' merging does not result in empty list")
	}

	if len(dataDiff.Builds) != 0 {
		t.Error(testedFuncName, "Empty build lists' merging does not result in empty list")
	}

	if len(dataDiff.Registries) != 0 {
		t.Error(testedFuncName, "Empty registry lists' merging does not result in empty list")
	}

	if len(dataDiff.Tests) != 0 {
		t.Error(testedFuncName, "Empty test lists' merging does not result in empty list")
	}

	if len(dataDiff.Results) != 0 {
		t.Error(testedFuncName, "Empty build result lists' merging does not result in empty list")
	}

	if !dataDiff.Day.IsZero() {
		t.Error(testedFuncName, "Null time value merging results in non-null value")
	}

	// Merge 2 non-empty objects

	initData = model.CollectedData{
		MAC:      "co:mp:ut:er",
		Username: "uname",
		Images: []model.Image{{
			Id:             "1",
			Name:           "Image1",
			ImageId:        "1",
			Description:    "A cool new image with all that fancy stuff",
			RegistryId:     "42",
			Tag:            "awesome",
			IlmTags:        []string{"yay", "hooray"},
			Location:       "public registry",
			SkipImageBuild: false,
		}, {
			Id:             "2",
			Name:           "Image2",
			ImageId:        "2",
			Description:    "A cool new image with all that fancy stuff",
			RegistryId:     "42",
			Tag:            "awesome",
			IlmTags:        []string{"yay", "hooray"},
			Location:       "public registry",
			SkipImageBuild: false,
		}},
		Projects: []model.Project{{
			Id:           "1",
			Name:         "Project1",
			CreationTime: "2002",
			Status:       "new",
			Images:       []model.Image{},
			Tests:        []model.Test{},
		},
			{
				Id:           "2",
				Name:         "Project2",
				CreationTime: "2002",
				Status:       "new",
				Images:       []model.Image{},
				Tests:        []model.Test{},
			}},
		Builds: []model.Build{{
			Id:        "1",
			ProjectId: "1",
			TestId:    "1",
			StartTime: "2002",
			Status:    model.Status{Status: "new"},
		},
			{
				Id:        "2",
				ProjectId: "2",
				TestId:    "2",
				StartTime: "2002",
				Status:    model.Status{Status: "new"},
			}},
		Registries: []model.Registry{{
			Id:   "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
			{
				Id:   "2",
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
			Id:            "1",
			BuildId:       "1",
			ResultEntries: []string{"1", "2"},
		},
			{
				Id:            "2",
				BuildId:       "2",
				ResultEntries: []string{"1", "2"},
			}},
		Day: time.Now(),
	}

	finalData = model.CollectedDataDiff{
		MAC:         "co:mp:ut:er",
		NewUserName: "uname",
		AddedImages: []model.Image{},
		DeletedImages: []model.Image{{
			Id:             "1",
			Name:           "Image1",
			ImageId:        "1",
			Description:    "A cool new image with all that fancy stuff",
			RegistryId:     "42",
			Tag:            "awesome",
			IlmTags:        []string{"yay", "hooray"},
			Location:       "public registry",
			SkipImageBuild: false,
		}, {
			Id:             "2",
			Name:           "Image2",
			ImageId:        "2",
			Description:    "A cool new image with all that fancy stuff",
			RegistryId:     "42",
			Tag:            "awesome",
			IlmTags:        []string{"yay", "hooray"},
			Location:       "public registry",
			SkipImageBuild: false,
		}},
		AddedProjects: []model.Project{},
		DeletedProjects: []model.Project{{
			Id:           "1",
			Name:         "Project1",
			CreationTime: "2002",
			Status:       "new",
			Images:       []model.Image{},
			Tests:        []model.Test{},
		},
			{
				Id:           "2",
				Name:         "Project2",
				CreationTime: "2002",
				Status:       "new",
				Images:       []model.Image{},
				Tests:        []model.Test{},
			}},
		AddedBuilds: []model.Build{},
		DeletedBuilds: []model.Build{{
			Id:        "1",
			ProjectId: "1",
			TestId:    "1",
			StartTime: "2002",
			Status:    model.Status{Status: "new"},
		},
			{
				Id:        "2",
				ProjectId: "2",
				TestId:    "2",
				StartTime: "2002",
				Status:    model.Status{Status: "new"},
			}},
		AddedRegistries: []model.Registry{},
		DeletedRegistries: []model.Registry{{
			Id:   "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
			{
				Id:   "2",
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
			Id:            "1",
			BuildId:       "1",
			ResultEntries: []string{"1", "2"},
		},
			{
				Id:            "2",
				BuildId:       "2",
				ResultEntries: []string{"1", "2"},
			}},
		NewDay: time.Now(),
	}

	dataDiff = MergeDiff(initData, finalData)

	if dataDiff.MAC != "co:mp:ut:er" {
		t.Error(testedFuncName, "MACs give an error when merging")
	}

	if len(dataDiff.Images) != 2 {
		t.Error(testedFuncName, "Image lists are not merging correctly")
	}

	if len(dataDiff.Projects) != 2 {
		t.Error(testedFuncName, "Project lists are not merging correctly")
	}

	if len(dataDiff.Builds) != 2 {
		t.Error(testedFuncName, "Build lists are not merging correctly")
	}

	if len(dataDiff.Registries) != 2 {
		t.Error(testedFuncName, "Registry lists are not merging  correctly")
	}

	if len(dataDiff.Tests) != 2 {
		t.Error(testedFuncName, "Test lists are not merging correctly")
	}

	if len(dataDiff.Results) != 2 {
		t.Error(testedFuncName, "Build result lists are not merging correctly")
	}

	finalData = model.CollectedDataDiff{}
	finalData.MAC = "co:mp:ut:er"

	dataDiff = MergeDiff(initData, finalData)

	if dataDiff.MAC != "co:mp:ut:er" {
		t.Error(testedFuncName, "MACs give an error when merging")
	}

	if len(dataDiff.Images) != 2 {
		t.Error(testedFuncName, "Image lists are not merging correctly")
	}

	if len(dataDiff.Projects) != 2 {
		t.Error(testedFuncName, "Project lists are not merging correctly")
	}

	if len(dataDiff.Builds) != 2 {
		t.Error(testedFuncName, "Build lists are not merging correctly")
	}

	if len(dataDiff.Registries) != 2 {
		t.Error(testedFuncName, "Registry lists are not merging  correctly")
	}

	if len(dataDiff.Tests) != 2 {
		t.Error(testedFuncName, "Test lists are not merging correctly")
	}

	if len(dataDiff.Results) != 2 {
		t.Error(testedFuncName, "Build result lists are not merging correctly")
	}

	finalData.MAC = "something different"

	dataDiff = MergeDiff(initData, finalData)

	if !CmpCollectedData(dataDiff, model.CollectedData{}) {
		t.Error(testedFuncName, "Two different MAC-ed datasets merged.")
	}

}
