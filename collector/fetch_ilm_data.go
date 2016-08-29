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

var token string
var credentials model.Credentials

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
	y := marshalOb(auth)
	split := strings.Split(y, ":")
	authToken := split[1]
	result := authToken[1 : len(authToken)-2]
	return result, err
}

func marshalOb(v interface{}) string {
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
	token = USERNAME + ":" + getAuthToken()
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
	projects := []model.Project{}
	token = USERNAME + ":" + getAuthToken()
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
	var result []model.Image
	var body2 []byte
	token = USERNAME + ":" + getAuthToken()
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
			myResult := []model.Image{}
			defer response.Body.Close()
			body2, err = ioutil.ReadAll(response.Body)

			json.Unmarshal(body2, &myResult)
			result = append(result, myResult...)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return result
}

func getImagesFromAProject(project model.Project) []model.Image {
	var result []model.Image
	var body2 []byte
	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + PROJECTPATH + "/"

	path := url + project.Id + "/images"

	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		myResult := []model.Image{}
		defer response.Body.Close()
		body2, err = ioutil.ReadAll(response.Body)

		json.Unmarshal(body2, &myResult)
		result = append(result, myResult...)

		if err != nil {
			log.Fatal(err)
		}
	}
	return result
}

func getTestsFromAProject(project model.Project) []model.Test {
	var result []model.Test
	var body2 []byte
	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + PROJECTPATH + "/"
	path := url + project.Id + "/tests"

	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		myResult := []model.Test{}
		defer response.Body.Close()
		body2, err = ioutil.ReadAll(response.Body)

		json.Unmarshal(body2, &myResult)
		result = append(result, myResult...)

		if err != nil {
			log.Fatal(err)
		}
	}
	return result
}
func getTestsFromApi() []model.Test {
	var body2 []byte
	var result []model.Test

	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + PROJECTPATH + "/"
	body := getProjectsfromApi()

	for _, data := range body {
		projId := url + data.Id + "/tests"
		client := &http.Client{}
		req, err := http.NewRequest("GET", projId, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			myResult := []model.Test{}

			defer response.Body.Close()
			body2, err = ioutil.ReadAll(response.Body)
			json.Unmarshal(body2, &myResult)
			result = append(result, myResult...)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return result
}

func getBuildsFromApi() []model.Build {
	var body2 []byte
	var result []model.Build
	myResult := []model.Build{}

	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + PROJECTPATH + "/"
	testsBody := getTestsFromApi()

	for _, data := range testsBody {
		testId := data.Id
		projId := data.ProjectId
		id := url + projId + "/tests/" + testId + "/builds"

		client := &http.Client{}
		req, err := http.NewRequest("GET", id, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			body2, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(body2, &myResult)
			result = append(result, myResult...)
		}
	}
	return result
}

func getResultsFromApi() []model.BuildResult {
	var body2 []byte
	var result []model.BuildResult
	myResult := []model.BuildResult{}

	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + PROJECTPATH + "/"
	buildsBody := getBuildsFromApi()
	for _, data := range buildsBody {
		testId := data.TestId
		projId := data.ProjectId
		buildId := data.Id
		id := url + projId + "/tests/" + testId + "/builds/" + buildId + "/results"

		client := &http.Client{}
		req, err := http.NewRequest("GET", id, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {

			defer response.Body.Close()
			body2, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(body2, &myResult)
		}
			result = append(result, myResult...)
	}
	return result

}
func getRegistriesFromAPi() []model.Registry {
	var body2 []byte

	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + REGISTRYPATH

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Access-Token", token)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body2, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	myResult := []model.Registry{}
	json.Unmarshal(body2, &myResult)

	return myResult
}

func getImagesFromRegistriesApi() []model.Repository {

	var result []model.Repository
	var body2 []byte
	token = USERNAME + ":" + getAuthToken()
	url := setUrl() + "/api/registries/"
	body := getRegistriesFromAPi()

	for _, data := range body {
		projId := url + data.Id + "/repositories"
		client := &http.Client{}
		req, err := http.NewRequest("GET", projId, nil)
		req.Header.Add("X-Access-Token", token)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			myResult := []model.Repository{}

			defer response.Body.Close()
			body2, err = ioutil.ReadAll(response.Body)
			json.Unmarshal(body2, &myResult)
			result = append(result, myResult...)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return result
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
