package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ilm-statistics/ilm-statistics/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CLIENTADDRESS = "CLIENTADDRESS"
	SERVERADDRESS = "SERVERADDRESS"
	PROTOCOL      = "http://"
	LOCALHOST     = "localhost"
	PORT_NAME     = "8082"
	AUTHPATH      = "/auth/login"
	ACCOUNTPATH   = "/api/accounts"
	PROJECTPATH   = "/api/projects"
	REGISTRYPATH  = "/api/registries"
	USERNAME      = "admin"
	PASSWORD      = "shipyard"
)

var credentials model.Credentials
var token = USERNAME + ":" + getAuthToken()

func setUrl() string {
	url := PROTOCOL + CLIENTADDRESS + ":" + PORT_NAME
	return url
}

func setCredentials(user model.Credentials) model.Credentials {
	user.Username = USERNAME
	user.Password = PASSWORD
	return user
}

func postAuthentication() []byte {
	path := setUrl() + AUTHPATH
	c := setCredentials(credentials)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	result, err := http.Post(path, "application/json; charset=utf-8", b)

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(result.Body)
	return body
}

func parseAuthResponse(body []byte) (string, error) {
	var auth model.Authentication

	err := json.Unmarshal(body, &auth)
	if err != nil {
		fmt.Println(err)
	}
	y := marshalObject(auth)
	split := strings.Split(y, ":")
	authToken := split[1]
	result := authToken[1 : len(authToken)-2]
	return result, err
}

func marshalObject(v interface{}) string {
	vBytes, _ := json.Marshal(v)
	return string(vBytes)
}
func getAuthToken() string {
	body := postAuthentication()
	s, _ := parseAuthResponse(body)
	x := string(s)
	return x
}
func getAccountsfromApi() []model.Account {
	var body []byte
	var accounts []model.Account
	url := setUrl() + ACCOUNTPATH

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.Unmarshal(body, &accounts)
	return accounts
}

func getProjectsfromApi() []model.Project {
	var body []byte
	var projects []model.Project
	url := setUrl() + PROJECTPATH

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}
	}
	json.Unmarshal(body, &projects)
	return projects
}

func getImagesfromApi() []model.Image {
	var images []model.Image
	var body []byte
	var projectImages []model.Image
	url := setUrl() + PROJECTPATH + "/"
	projects := getProjectsfromApi()

	for _, project := range projects {
		path := url + project.Id + "/images"

		client := &http.Client{}
		req, err := http.NewRequest("GET", path, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)

			json.Unmarshal(body, &projectImages)
			images = append(images, projectImages...)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return images
}

func getImagesFromAProject(project model.Project) []model.Image {
	var images []model.Image
	var projectImages []model.Image
	var body []byte

	url := setUrl() + PROJECTPATH + "/"
	path := url + project.Id + "/images"

	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)

		json.Unmarshal(body, &projectImages)
		images = append(images, projectImages...)

		if err != nil {
			log.Fatal(err)
		}
	}
	return images
}

func getTestsFromAProject(project model.Project) []model.Test {
	var tests []model.Test
	var projectTests []model.Test
	var body []byte

	url := setUrl() + PROJECTPATH + "/"
	path := url + project.Id + "/tests"

	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		json.Unmarshal(body, &projectTests)
		tests = append(tests, projectTests...)

		if err != nil {
			log.Fatal(err)
		}
	}
	return tests
}
func getTestsFromApi() []model.Test {
	var tests []model.Test
	var projectTests []model.Test
	var body []byte

	url := setUrl() + PROJECTPATH + "/"
	projects := getProjectsfromApi()

	for _, project := range projects {
		id := url + project.Id + "/tests"
		client := &http.Client{}
		req, err := http.NewRequest("GET", id, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &projectTests)
			tests = append(tests, projectTests...)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return tests
}

func getBuildsFromApi() []model.Build {
	var builds []model.Build
	var projectBuilds []model.Build
	var body []byte

	url := setUrl() + PROJECTPATH + "/"
	projects := getProjectsfromApi()

	for _, project := range projects {
		testIds := project.TestIds
		projId := project.Id

		for _, testId := range testIds {
			id := url + projId + "/tests/" + testId + "/builds"

			client := &http.Client{}
			req, err := http.NewRequest("GET", id, nil)
			req.Header.Add("X-Access-Token", token)
			response, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {
				defer response.Body.Close()
				body, err = ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				json.Unmarshal(body, &projectBuilds)
				builds = append(builds, projectBuilds...)
			}
		}
	}
	return builds
}

func getResultsFromApi() []model.BuildResult {
	var buildResults []model.BuildResult
	var projectBuildResults []model.BuildResult
	var body []byte

	url := setUrl() + PROJECTPATH + "/"
	builds := getBuildsFromApi()

	for _, build := range builds {
		testId := build.TestId
		projId := build.ProjectId
		buildId := build.Id
		id := url + projId + "/tests/" + testId + "/builds/" + buildId + "/results"

		client := &http.Client{}
		req, err := http.NewRequest("GET", id, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(body, &projectBuildResults)
		}
		buildResults = append(buildResults, projectBuildResults...)
	}

	return buildResults
}
func getRegistriesFromAPi() []model.Registry {
	var body []byte
	var registries []model.Registry

	url := setUrl() + REGISTRYPATH

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.Unmarshal(body, &registries)

	return registries
}

func getImagesFromRegistriesApi() []model.Repository {

	var images []model.Repository
	var repoImages []model.Repository
	var body []byte

	url := setUrl() + "/api/registries/"
	registries := getRegistriesFromAPi()

	for _, registry := range registries {
		id := url + registry.Id + "/repositories"
		client := &http.Client{}
		req, err := http.NewRequest("GET", id, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &repoImages)
			images = append(images, repoImages...)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return images
}

func setStatistics(stats model.CollectedData) model.CollectedData {
	uname := "admin"
	var result []model.Project
	projects := getProjectsfromApi()
	for _, proj := range projects {
		img := getImagesFromAProject(proj)
		tst := getTestsFromAProject(proj)
		proj.Images = img
		proj.Tests = tst
		result = append(result, proj)
	}
	stats.Projects = result

	acc := getAccountsfromApi()
	imgs := getImagesfromApi()
	tsts := getTestsFromApi()
	reg := getRegistriesFromAPi()
	bld := getBuildsFromApi()
	res := getResultsFromApi()
	repo := getImagesFromRegistriesApi()
	stats.Username = uname
	stats.Accounts = acc
	stats.Images = imgs
	stats.Tests = tsts
	stats.Registries = reg
	stats.Builds = bld
	stats.Results = res
	stats.Repositories = repo
	return stats
}

func postResponse() {
	var stats model.CollectedData
	s := setStatistics(stats)
	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(s)
	result, _ := http.Post(SERVERADDRESS, "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, result.Body)
}

/*func main() {
	postResponse()
}*/
