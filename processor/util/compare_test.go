package util

import (
	"testing"
	"github.com/ilm-statistics/ilm-statistics/model"
	"time"
)

func TestCmpImages(t *testing.T){
	// Two equal images
	img1 := model.Image{
		ProjectId: "1",
		Id: "1",
		Name: "Image1",
		ImageId: "1",
		Description: "A cool new image with all that fancy stuff",
		Status: "new",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: "false",
	}
	img2 := model.Image{
		ProjectId: "1",
		Id: "1",
		Name: "Image1",
		ImageId: "1",
		Description: "A cool new image with all that fancy stuff",
		Status: "new",
		RegistryId: "42",
		Tag: "awesome",
		IlmTags: []string{"yay", "hooray"},
		Location: "public registry",
		SkipImageBuild: "false",
	}

	if !CmpImages(img1,img2) {
		t.Error("Two equal images evaluated as non-equal")
	}

	// One image is different from the other in one field (ProjectId)
	img2.ProjectId = "3"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different ProjectId)")
	}

	// One image is different from the other in one field (Name)
	img2.ProjectId = "1"
	img2.Name = "Image2"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different Name)")
	}

	// One image is different from the other in one field (ImageId)
	img2.Name = "Image1"
	img2.ImageId = "2"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different ImageId)")
	}

	// One image is different from the other in one field (Description)
	img2.ImageId = "1"
	img2.Description = "A cool new image with none of that fancy stuff"

	if CmpImages(img1, img2) {
		t.Error("Two non-equal images evaluated as equal (different Description)")
	}

	// One image is different from the other in one field (Status)
	img2.Description = "A cool new image with all that fancy stuff"
	img2.Status = "old"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different Status)")
	}

	// One image is different from the other in one field (RegistryId)
	img2.Status = "new"
	img2.RegistryId = "45"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different RegistryId)")
	}

	// One image is different from the other in one field (Tag)
	img2.RegistryId = "42"
	img2.Tag = "awful"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different Tag)")
	}

	// One image is different from the other in one field (IlmTags)
	img2.Tag = "awesome"
	img2.IlmTags = []string{"boo"}

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different IlmTags)")
	}

	// One image is different from the other in one field (Location)
	img2.IlmTags = []string{"yay", "hooray"}
	img2.Location = "private registry"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different Location)")
	}

	// One image is different from the other in one field (SkipImageBuild)
	img2.Location = "public registry"
	img2.SkipImageBuild = "true"

	if CmpImages(img1,img2) {
		t.Error("Two non-equal images evaluated as equal (different SkipImageBuild)")
	}
}

func TestCmpAccounts(t *testing.T) {

	// Two similar accounts
	acc1 := model.Account{
		Id: "1",
		FirstName: "Bruce",
		LastName: "Wayne",
		Username: "batman",
		Password: "batmobile",
		Roles: []string{"hero", "billionaire", "philanthropist"},
	}
	acc2 := model.Account{
		Id: "1",
		FirstName: "Bruce",
		LastName: "Wayne",
		Username: "batman",
		Password: "batmobile",
		Roles: []string{"hero", "billionaire", "philanthropist"},
	}
	if !CmpAccounts(acc1, acc2){
		t.Error("Two equal accounts evaluated as non-equal")
	}

	// One account is different from the other in one field (Id)
	acc2.Id = "2"

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different Id)")
	}

	// One account is different from the other in one field (FirstName)
	acc2.Id = "1"
	acc2.FirstName = "Brian"

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different FirstName)")
	}

	// One account is different from the other in one field (LastName)
	acc2.FirstName = "Bruce"
	acc2.LastName = "Lee"

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different LastName)")
	}

	// One account is different from the other in one field (Username)
	acc2.LastName = "Wayne"
	acc2.Username = "batsy"

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different Username)")
	}

	// One account is different from the other in one field (Password)
	acc2.Username = "batman"
	acc2.Password = "#batmobile"

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different Password)")
	}

	// One account is different from the other in one field (Roles)
	acc2.Password = "batmobile"
	acc2.Roles = []string{}

	if CmpAccounts(acc1, acc2){
		t.Error("Two non-equal accounts evaluated as equal (different Roles)")
	}

	// One account is different from the other in one field (different order in Roles) -> they should evaluate as equal
	acc2.Roles = []string{"billionaire", "philanthropist", "hero"}

	if !CmpAccounts(acc1, acc2){
		t.Error("Two equal accounts evaluated as non-equal (different order in Roles)")
	}

}

func TestCmpProviders(t *testing.T) {

	// Two equal providers
	prov1 := model.Provider{
		ProviderType: "clair",
	}
	prov2 := model.Provider {
		ProviderType: "clair",
	}

	if !CmpProviders(prov1, prov2) {
		t.Error("Two equal providers evaluated as non-equal")
	}

	// One provider is different from the other in one field (providerType)
	prov2.ProviderType = "clair2"

	if CmpProviders(prov1, prov2) {
		t.Error("Two non-equal providers evaluated as equal")
	}
}

func TestCmpTests(t *testing.T) {

	//Two equal tests
	test1 := model.Test{
		Id: "1",
		ProjectId: "1",
		Provider: model.Provider{
			ProviderType: "clair",
		},
	}
	test2 := model.Test{
		Id: "1",
		ProjectId: "1",
		Provider: model.Provider{
			ProviderType: "clair",
		},
	}

	if !CmpTests(test1, test2){
		t.Error("Two equal tests evaluated as non-equal")
	}

	// One test is different from the other in one field (Id)
	test2.Id = "2"

	if CmpTests(test1,test2){
		t.Error("Two non-equal tests evaluated as equal (different Id)")
	}

	// One test is different from the other in one field (ProjectId)
	test2.Id = "1"
	test2.ProjectId = "2"

	if CmpTests(test1,test2){
		t.Error("Two non-equal tests evaluated as equal (different ProjectId)")
	}

	// One test is different from the other in one field (Provider)
	test2.ProjectId = "1"
	test2.Provider.ProviderType = "not clair"

	if CmpTests(test1,test2){
		t.Error("Two non-equal tests evaluated as equal (different Provider)")
	}
}

func TestCmpProjects(t *testing.T) {

	//Two equal projects
	proj1 := model.Project{
		Id: "1",
		Name: "Project1",
		Author: "Author1",
		CreationTime: "2002",
		LastRunTime: "2002",
		Status: "new",
		Images: []model.Image{},
		Tests: []model.Test{},
	}
	proj2 := model.Project{
		Id: "1",
		Name: "Project1",
		Author: "Author1",
		CreationTime: "2002",
		LastRunTime: "2002",
		Status: "new",
		Images: []model.Image{},
		Tests: []model.Test{},
	}

	if !CmpProjects(proj1, proj2) {
		t.Error("Two equal projects evaluated as non-equal")
	}

	// One project is different from the other in one field (Id)
	proj2.Id = "2"

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different Id)")
	}

	// One project is different from the other in one field (Name)
	proj2.Id = "1"
	proj2.Name = "Project2"

	if CmpProjects(proj1,proj2){
		t.Error("Two non-equal projects evaluated as equal (different Name)")
	}

	// One project is different from the other in one field (Author)
	proj2.Name = "Project1"
	proj2.Author = "Author2"

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different Author)")
	}

	// One project is different from the other in one field (CreationTime)
	proj2.Author = "Author1"
	proj2.CreationTime = "2003"

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different CreationTime)")
	}

	// One project is different from the other in one field (LastRunTime)
	proj2.CreationTime = "2002"
	proj2.LastRunTime = "2016"

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different LastRunTime)")
	}

	// One project is different from the other in one field (Status)
	proj2.LastRunTime = "2002"
	proj2.Status = "old"

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different Status)")
	}

	// One project is different from the other in one field (Images)
	proj2.Status = "new"
	proj2.Images = []model.Image{
		{
			ProjectId: "1",
			Id: "1",
			Name: "Image1",
			ImageId: "1",
			Description: "A cool new image with all that fancy stuff",
			Status: "new",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		},
		{
			ProjectId: "1",
			Id: "12",
			Name: "Image1",
			ImageId: "12",
			Description: "A cool new image with all that fancy stuff",
			Status: "new",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		},
	}

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different Images)")
	}

	// One project is different from the other in one field (different order in Images) -> they should evaluate as equal
	proj1.Images = []model.Image{
		{
			ProjectId: "1",
			Id: "12",
			Name: "Image1",
			ImageId: "12",
			Description: "A cool new image with all that fancy stuff",
			Status: "new",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		},
		{
			ProjectId: "1",
			Id: "1",
			Name: "Image1",
			ImageId: "1",
			Description: "A cool new image with all that fancy stuff",
			Status: "new",
			RegistryId: "42",
			Tag: "awesome",
			IlmTags: []string{"yay", "hooray"},
			Location: "public registry",
			SkipImageBuild: "false",
		},
	}

	if !CmpProjects(proj1, proj2){
		t.Error("Two equal projects evaluated as non-equal (different order in Images)")
	}

	// One project is different from the other in one field (Tests)
	proj2.Tests = []model.Test{
		{
			Id: "1",
			ProjectId: "1",
			Provider: model.Provider{
				ProviderType: "clair",
			},
		},
		{
			Id: "2",
			ProjectId: "1",
			Provider: model.Provider{
				ProviderType: "clair",
			},
		},
	}

	if CmpProjects(proj1, proj2){
		t.Error("Two non-equal projects evaluated as equal (different Tests)")
	}
}

func TestCmpBuilds(t *testing.T) {

	// Two equal builds
	build1 := model.Build{
		Id: "1",
		ProjectId: "1",
		TestId: "1",
		StartTime: "2002",
		Status: model.Status{Status:"new"},
	}
	build2 := model.Build{
		Id: "1",
		ProjectId: "1",
		TestId: "1",
		StartTime: "2002",
		Status: model.Status{Status:"new",},
	}

	if !CmpBuilds(build1, build2){
		t.Error("Two equal projects evaluated as non-equal")
	}

	// One of the builds is different from the other in one field (Id)
	build2.Id = "2"

	if CmpBuilds(build1, build2) {
		t.Error("Two non-equal projects evaluated as equal (different Id)")
	}

	//One of the builds is different from the other in one field (ProjectsId)
	build2.Id = "1"
	build2.ProjectId = "2"

	if CmpBuilds(build1, build2) {
		t.Error("Two non-equal projects evaluated as equal (different ProjectId)")
	}

	//One of the builds is different from the other in one field (TestId)
	build2.ProjectId = "1"
	build2.TestId = "2"

	if CmpBuilds(build1, build2) {
		t.Error("Two non-equal projects evaluated as equal (different TestId)")
	}

	//One of the builds is different from the other in one field (StartTime)
	build2.TestId = "1"
	build2.StartTime = "2016"

	if CmpBuilds(build1, build2) {
		t.Error("Two non-equal projects evaluated as equal (different StartTime)")
	}

	//One of the builds is different from the other in one field (Status)
	build2.StartTime = "2002"
	build2.Status.Status = "old"

	if CmpBuilds(build1, build2) {
		t.Error("Two non-equal projects evaluated as equal (different Status)")
	}

}

func TestCmpRegistries(t *testing.T) {

	// Two equal registries
	reg1 := model.Registry{
		Id: "1",
		Name: "Private1",
		Addr: "localhost:5000",
	}
	reg2 := model.Registry{
		Id: "1",
		Name: "Private1",
		Addr: "localhost:5000",
	}

	if !CmpRegistries(reg1,reg2){
		t.Error("Two equal registries evaluated as non-equal")
	}

	// One of the registries is different from the other in one field (Id)
	reg2.Id = "2"

	if CmpRegistries(reg1,reg2){
		t.Error("Two non-equal registries evaluated as equal (different Id)")
	}

	// One of the registries is different from the other in one field (Name)
	reg2.Id = "1"
	reg2.Name = "Private2"

	if CmpRegistries(reg1,reg2){
		t.Error("Two non-equal registries evaluated as equal (different Name)")
	}

	// One of the registries is different from the other in one field (Addr)
	reg2.Name = "Private1"
	reg2.Addr = "localhost:5001"

	if CmpRegistries(reg1,reg2){
		t.Error("Two non-equal registries evaluated as equal (different Addr)")
	}

}

func TestCmpBuildResults(t *testing.T) {

	// Two equal buildResults
	res1 := model.BuildResult{
		ID: "1",
		BuildId: "1",
		ResultEntries: []string{"1","2"},
	}

	res2 := model.BuildResult{
		ID: "1",
		BuildId: "1",
		ResultEntries: []string{"1","2"},
	}

	if !CmpBuildResults(res1,res2){
		t.Error("Two equal build results evaluated as non-equal")
	}

	// One of the build results is different from the other in one field (ID)
	res2.ID = "2"

	if CmpBuildResults(res1,res2){
		t.Error("Two non-equal build results evaluated as equal (different ID)")
	}

	// One of the build results is different from the other in one field (BuildId)
	res2.ID = "1"
	res2.BuildId = "2"

	if CmpBuildResults(res1,res2){
		t.Error("Two non-equal build results evaluated as equal (different BuildId)")
	}

	// One of the build results is different from the other in one field (ResultEntries)
	res2.BuildId = "1"
	res2.ResultEntries = []string{}

	if CmpBuildResults(res1,res2){
		t.Error("Two non-equal build results evaluated as equal (different ResultEntries)")
	}

	// One build result is different from the other in one field (different order in ResultEntries) -> they should evaluate as equal
	res2.ResultEntries = []string{"2","1"}

	if !CmpBuildResults(res1,res2){
		t.Error("Two equal build results evaluated as non-equal (different order in ResultEntries)")
	}
}

func TestCmpRepositories(t *testing.T) {

	// Two equal repositories
	repo1 := model.Repository{
		Name: "Repo1",
		Tag: "private",
		FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
		Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
		HasProblems: false,
		Message: "msg",
		RegistryUrl: "localhost:5000",
		RegistryName: "Reggy",
	}
	repo2 := model.Repository{
		Name: "Repo1",
		Tag: "private",
		FsLayers: []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}},
		Signatures: []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, },
		HasProblems: false,
		Message: "msg",
		RegistryUrl: "localhost:5000",
		RegistryName: "Reggy",
	}

	if !CmpRepositories(repo1,repo2){
		t.Error("Two equal repositories evaluates as non-equal")
	}

	// One repository is different from the other in one field (Name)
	repo2.Name = "Repo2"

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different Name)")
	}

	// One repository is different from the other in one field (Tag)
	repo2.Name = "Repo1"
	repo2.Tag = "public"

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different Tag)")
	}

	// One repository is different from the other in one field (FsLayers)
	repo2.Tag = "private"
	repo2.FsLayers = []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "500"}}
	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different FsLayers)")
	}

	// One repository is different from the other in one field (Signatures)
	repo2.FsLayers = []model.FsLayer{ {BlobSum: "50",}, {BlobSum: "100"}}
	repo2.Signatures = []model.Signature{ {Header: model.Header{Algorithm:"algs12"}, Signature: "signed", Protected: "yes" }, }

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different Signatures)")
	}

	// One repository is different from the other in one field (HasProblems)
	repo2.Signatures = []model.Signature{ {Header: model.Header{Algorithm:"algs"}, Signature: "signed", Protected: "yes" }, }
	repo2.HasProblems = true

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different HasProblems)")
	}

	// One repository is different from the other in one field (Message)
	repo2.HasProblems = false
	repo2.Message = "Hello World!"

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different Message)")
	}

	// One repository is different from the other in one field (RegistryUrl)
	repo2.Message = "msg"
	repo2.RegistryUrl = "localhost:8080"

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different RegistryUrl)")
	}

	// One repository is different from the other in one field (RegistryName)
	repo2.RegistryUrl = "localhost:5000"
	repo2.RegistryName = "Registry"

	if CmpRepositories(repo1,repo2){
		t.Error("Two non-equal repositories evaluated as equal (different RegistryName)")
	}
}

func TestCmpCollectedData(t *testing.T) {
	data1 := model.CollectedData{}
	data2 := model.CollectedData{}

	// Test two empty data struncts
	if !CmpCollectedData(data1, data2){
		t.Error("Two empty structs evaluated as non-equal")
	}

	data1.MAC = "co:mp:ut:er"
	data2.MAC = "co:mp:ut:er"

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (MAC set)")
	}

	data1.Username = "uname"
	data2.Username = "uname"

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (username set)")
	}

	data1.Images = append(data1.Images, model.Image{ProjectId: "1", Id: "1", Name: "image1"})
	data2.Images = append(data2.Images, model.Image{ProjectId: "1", Id: "1", Name: "image1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (images set)")
	}

	data1.Accounts = append(data1.Accounts, model.Account{Id: "1"})
	data2.Accounts = append(data2.Accounts, model.Account{Id: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (accounts set)")
	}

	data1.Projects = append(data1.Projects, model.Project{Id: "1"})
	data2.Projects = append(data2.Projects, model.Project{Id: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (projects set)")
	}

	data1.Builds = append(data1.Builds, model.Build{Id: "1"})
	data2.Builds = append(data2.Builds, model.Build{Id: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (builds set)")
	}

	data1.Registries = append(data1.Registries, model.Registry{Id: "1"})
	data2.Registries = append(data2.Registries, model.Registry{Id: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (registries set)")
	}

	data1.Tests = append(data1.Tests, model.Test{Id: "1"})
	data2.Tests = append(data2.Tests, model.Test{Id: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (tests set)")
	}

	data1.Results = append(data1.Results, model.BuildResult{ID: "1"})
	data2.Results = append(data2.Results, model.BuildResult{ID: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (build results set)")
	}

	data1.Repositories = append(data1.Repositories, model.Repository{Name: "1"})
	data2.Repositories = append(data2.Repositories, model.Repository{Name: "1"})

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (repositories set)")
	}

	data1.Day = time.Now()
	data2.Day = data1.Day

	if !CmpCollectedData(data1, data2) {
		t.Error("Two equal structs evaluated as non-equal (days set)")
	}
}