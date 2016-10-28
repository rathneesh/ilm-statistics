package util

import (
	"testing"
	"github.com/ilm-statistics/ilm-statistics/model"
	"time"
)

func TestDiffImageList(t *testing.T) {
	t.Parallel()

	// Test if two lists of images are equal
	imgs1 := []model.Image{{
		Id: "1",
		Name: "Image1",
		ImageId: "1",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}, {
		Id: "2",
		Name: "Image2",
		ImageId: "2",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}}

	imgs2 := []model.Image{{
		Id: "1",
		Name: "Image1",
		ImageId: "1",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}, {
		Id: "2",
		Name: "Image2",
		ImageId: "2",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}}

	addedImgs, deletedImgs := DiffImageList(imgs1, imgs2)
	if len(addedImgs) != 0 || len(deletedImgs) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding image
	img := model.Image{
		Id: "3",
		Name: "Image3",
		ImageId: "3",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}
	imgs2 = append(imgs2, img)

	addedImgs, deletedImgs = DiffImageList(imgs1, imgs2)

	if len(addedImgs) != 1 || len(deletedImgs) != 0 || !CmpImages(addedImgs[0], img){
		t.Error("Added image not recognized")
	}

	// Test for deleting image
	imgs1 = append(imgs1, img)

	img = model.Image{
		Id: "4",
		Name: "Image4",
		ImageId: "4",
		Description: "A cool new image with all that fancy stuff",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: false,
	}
	imgs1 = append(imgs1, img)

	addedImgs, deletedImgs = DiffImageList(imgs1, imgs2)

	if len(addedImgs) != 0 || len(deletedImgs) != 1 || !CmpImages(deletedImgs[0], img){
		t.Error("Deleted image not recognized")
	}

}

func TestDiffProjectList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists
	projs1 := []model.Project {
		{
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
		},
	}
	projs2 := []model.Project {
		{
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
		},
	}

	addedProjects, deletedProjects := DiffProjectList(projs1, projs2)

	if len(addedProjects) != 0 || len(deletedProjects) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding project

	proj := model.Project{Id: "3",
		Name: "Project3",
		CreationTime: "2002",
		Status: "new",
		Images: []model.Image{},
		Tests: []model.Test{},
	}

	projs2 = append(projs2, proj)

	addedProjects, deletedProjects = DiffProjectList(projs1, projs2)

	if len(addedProjects) != 1 || len(deletedProjects) != 0 || !CmpProjects(addedProjects[0], proj) {
		t.Error("Added account not recognized")
	}

	// Test for deleting project

	projs1 = append(projs1, proj)

	proj = model.Project{Id: "4",
		Name: "Project4",
		CreationTime: "2002",
		Status: "new",
		Images: []model.Image{},
		Tests: []model.Test{},
	}

	projs1 = append(projs1, proj)

	addedProjects, deletedProjects = DiffProjectList(projs1, projs2)

	if len(addedProjects) != 0 || len(deletedProjects) != 1 || !CmpProjects(deletedProjects[0], proj) {
		t.Error("Deleted account not recognized")
	}
}

func TestDiffBuildList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists

	builds1 := []model.Build{{
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
		},
	}

	builds2 := []model.Build{{
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
		},
	}

	addedBuilds, deletedBuilds := DiffBuildList(builds1, builds2)

	if len(addedBuilds) != 0 || len(deletedBuilds) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding build

	build := model.Build{
		Id: "3",
		ProjectId: "3",
		TestId: "3",
		StartTime: "2002",
		Status: model.Status{Status:"new"},
	}

	builds2 = append(builds2, build)
	addedBuilds, deletedBuilds = DiffBuildList(builds1, builds2)

	if len(addedBuilds) != 1 || len(deletedBuilds) != 0 || !CmpBuilds(addedBuilds[0], build) {
		t.Error("Added build not recognized")
	}

	// Test for deleting build

	builds1 = append(builds1, build)

	build = model.Build{
		Id: "4",
		ProjectId: "4",
		TestId: "4",
		StartTime: "2002",
		Status: model.Status{Status:"new"},
	}

	builds1 = append(builds1, build)
	addedBuilds, deletedBuilds = DiffBuildList(builds1, builds2)

	if len(addedBuilds) != 0 || len(deletedBuilds) != 1 || !CmpBuilds(deletedBuilds[0], build) {
		t.Error("Deleted build not recognized")
	}

}

func TestDiffRegistryList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists

	regs1 := []model.Registry{
		{
			Id: "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
		{
			Id: "2",
			Name: "Private2",
			Addr: "localhost:5002",
		},
	}

	regs2 := []model.Registry{
		{
			Id: "1",
			Name: "Private1",
			Addr: "localhost:5000",
		},
		{
			Id: "2",
			Name: "Private2",
			Addr: "localhost:5002",
		},
	}

	addedRegistries, deletedRegistries := DiffRegistryList(regs1,regs2)

	if len(addedRegistries) != 0 || len(deletedRegistries) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding registry

	reg := model.Registry{
		Id : "3",
		Name: "Private3",
		Addr: "localhost:5000",
	}

	regs2 = append(regs2, reg)

	addedRegistries, deletedRegistries = DiffRegistryList(regs1, regs2)

	if len(addedRegistries) != 1 || len(deletedRegistries) != 0 || !CmpRegistries(addedRegistries[0], reg){
		t.Error("Added registry not recognized")
	}

	// Test for deleting registry -
	regs1 = append(regs1, reg)

	reg = model.Registry{
		Id: "4",
		Name: "Private4",
		Addr: "localhost:5000",
	}
	regs1 = append(regs1, reg)

	addedRegistries, deletedRegistries = DiffRegistryList(regs1, regs2)

	if len(addedRegistries) != 0 || len(deletedRegistries) != 1 || !CmpRegistries(deletedRegistries[0], reg) {
		t.Error("Deleted registry not recognized")
	}
}

func TestDiffTestList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists

	tests1 := []model.Test{
		{
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
		},
	}

	tests2 := []model.Test{
		{
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
		},
	}

	addedTests, deletedTests := DiffTestList(tests1, tests2)

	if len(addedTests) != 0 || len(deletedTests) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding test

	test := model.Test{
		Id: "3",
		Provider: model.Provider{
			ProviderType: "clair",
		},
	}

	tests2 = append(tests2, test)

	addedTests, deletedTests = DiffTestList(tests1, tests2)

	if len(addedTests) != 1 || len(deletedTests) != 0 || !CmpTests(addedTests[0], test) {
		t.Error("Added test not recognized")
	}

	// Test for deleting test
	tests1 = append(tests1, test)

	test = model.Test{
		Id: "4",
		Provider: model.Provider{
			ProviderType: "clair",
		},
	}

	tests1 = append(tests1, test)

	addedTests, deletedTests = DiffTestList(tests1, tests2)

	if len(addedTests) != 0 || len(deletedTests) != 1 || !CmpTests(deletedTests[0], test) {
		t.Error("Deleted test not recognized")
	}
}

func TestDiffResultList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists

	results1 := []model.BuildResult{
		{
			Id: "1",
			BuildId: "1",
			ResultEntries: []string{"1","2"},
		},
		{
			Id: "2",
			BuildId: "2",
			ResultEntries: []string{"1","2"},
		},
	}

	results2 := []model.BuildResult{
		{
			Id: "1",
			BuildId: "1",
			ResultEntries: []string{"1","2"},
		},
		{
			Id: "2",
			BuildId: "2",
			ResultEntries: []string{"1","2"},
		},
	}

	addedResults, deletedResults := DiffResultList(results1, results2)

	if len(addedResults) != 0 || len(deletedResults) != 0 {
		t.Error("Two similar results evaluated as different")
	}

	// Test for adding result

	result := model.BuildResult{
		Id: "3",
		BuildId: "3",
		ResultEntries: []string{"1","2"},
	}

	results2 = append(results2, result)

	addedResults, deletedResults = DiffResultList(results1, results2)

	if len(addedResults) != 1 || len(deletedResults) != 0  || !CmpBuildResults(addedResults[0], result){
		t.Error("Added result not recognized")
	}

	// Test for deleting result

	results1 = append(results1, result)

	result = model.BuildResult{
		Id: "4",
		BuildId: "4",
		ResultEntries: []string{"1", "2"},
	}

	results1 = append(results1, result)

	addedResults, deletedResults = DiffResultList(results1, results2)

	if len(addedResults) != 0 || len(deletedResults) != 1 || !CmpBuildResults(deletedResults[0], result) {
		t.Error("Deleted result not recognized")
	}
}

func TestDiffRepositoryList(t *testing.T) {
	t.Parallel()

	// Test for two similar lists

	repos1 := []model.Repository{
		{
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
		},
	}

	repos2 := []model.Repository{
		{
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
		},
	}

	addedRepositories, deletedRepositories := DiffRepositoryList(repos1, repos2)

	if len(addedRepositories) != 0 || len(deletedRepositories) != 0 {
		t.Error("Two similar lists evaluated as different")
	}

	// Test for adding repository

	repo := model.Repository{
		Name: "Repo3",
		Tag: "private-ish",
		FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
		Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
		HasProblems: false,
		Message: "Another message",
		RegistryUrl: "localhost:5000",
		RegistryName: "Randy",
	}

	repos2 = append(repos2, repo)

	addedRepositories, deletedRepositories = DiffRepositoryList(repos1, repos2)

	if len(addedRepositories) != 1 || len(deletedRepositories) != 0 || !CmpRepositories(addedRepositories[0], repo){
		t.Error("Added repository not recognized")
	}

	// Test for deleting repository

	repos1 = append(repos1, repo)

	repo = model.Repository{
		Name: "Repo4",
		Tag: "public-ish",
		FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
		Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
		HasProblems: false,
		Message: "... and another message",
		RegistryUrl: "localhost:5000",
		RegistryName: "Robert",
	}

	repos1 = append(repos1, repo)

	addedRepositories, deletedRepositories = DiffRepositoryList(repos1, repos2)

	if len(addedRepositories) != 0 || len(deletedRepositories) != 1 || !CmpRepositories(deletedRepositories[0], repo){
		t.Error("Deleted repository not recognized")
	}
}

func TestDiffCollectedData(t *testing.T) {
	t.Parallel()

	oldData := model.CollectedData{}
	newData := model.CollectedData{}
	diff := DiffCollectedData(oldData, newData)

	if diff.NewUserName != "" {
		t.Error("Difference of username between two empty data structs was found")
	}

	if len(diff.AddedImages) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Images)")
	}

	if len(diff.DeletedImages) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Images)")
	}

	if len(diff.AddedProjects) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Projects)")
	}

	if len(diff.DeletedProjects) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Projects)")
	}

	if len(diff.AddedBuilds) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Builds)")
	}

	if len(diff.DeletedBuilds) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Builds)")
	}

	if len(diff.AddedRegistries) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Registries)")
	}

	if len(diff.DeletedRegistries) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Registries)")
	}

	if len(diff.AddedTests) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Tests)")
	}

	if len(diff.DeletedTests) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Tests)")
	}

	if len(diff.AddedResults) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Added Results)")
	}

	if len(diff.DeletedResults) != 0 {
		t.Error("Non-empty list resulted from differentiating two empty lists (Deleted Results)")
	}

	if !diff.NewDay.IsZero() {
		t.Error("Time difference found between to zero entities (New Day)")
	}


	newData.Username = "uname"
	newData.Images = append(newData.Images, model.Image{Id: "1", Name: "image1"})
	oldData.Images = append(oldData.Images, model.Image{Id: "2", Name: "image2"})
	newData.Projects = append(newData.Projects, model.Project{Id: "1"})
	oldData.Projects = append(oldData.Projects, model.Project{Id: "2"})
	newData.Builds = append(newData.Builds, model.Build{Id: "1"})
	oldData.Builds = append(oldData.Builds, model.Build{Id: "2"})
	newData.Registries = append(newData.Registries, model.Registry{Id: "1"})
	oldData.Registries = append(oldData.Registries, model.Registry{Id: "2"})
	newData.Tests = append(newData.Tests, model.Test{Id: "1"})
	oldData.Tests = append(oldData.Tests, model.Test{Id: "2"})
	newData.Results = append(newData.Results, model.BuildResult{Id: "1"})
	oldData.Results = append(oldData.Results, model.BuildResult{Id: "2"})
	newData.Day = time.Now()

	diff = DiffCollectedData(oldData, newData)

	if diff.NewUserName != "uname" {
		t.Error("Username differentiation does not work")
	}

	if len(diff.AddedImages) != 1 {
		t.Error("Image list differentiation does not work (add)")
	}

	if len(diff.DeletedImages) != 1 {
		t.Error("Image list differentiation does not work (delete)")
	}

	if len(diff.AddedProjects) != 1 {
		t.Error("Project list differentiation does not work (add)")
	}

	if len(diff.DeletedProjects) != 1 {
		t.Error("Project list differentiation does not work (delete)")
	}

	if len(diff.AddedBuilds) != 1 {
		t.Error("Build list differentiation does not work (add)")
	}

	if len(diff.DeletedBuilds) != 1 {
		t.Error("Build list differentiation does not work (delete)")
	}

	if len(diff.AddedRegistries) != 1 {
		t.Error("Registry list differentiation does not work (add)")
	}

	if len(diff.DeletedRegistries) != 1 {
		t.Error("Registry list differentiation does not work (delete)")
	}

	if len(diff.AddedTests) != 1 {
		t.Error("Test list differentiation does not work (add)")
	}

	if len(diff.DeletedTests) != 1 {
		t.Error("Test list differentiation does not work (delete)")
	}

	if len(diff.AddedResults) != 1 {
		t.Error("Result list differentiation does not work (add)")
	}

	if len(diff.DeletedResults) != 1 {
		t.Error("Result list differentiation does not work (delete)")
	}

	if diff.NewDay != newData.Day {
		t.Error("Time differentiation does not work")
	}
}