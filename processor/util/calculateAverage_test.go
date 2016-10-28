package util

import (
	"testing"
	"github.com/ilm-statistics/ilm-statistics/model"
)

func TestSeparateByIp(t *testing.T) {
	t.Parallel()

	testedFuncName := "SeparateByIp: "

	// nil data
	result := SeparateByIp(nil)
	if len(result) != 0 {
		t.Error(testedFuncName + "On nil data returned map with length " + string(len(result)))
	}

	// Empty array
	var collData []model.CollectedData
	result = SeparateByIp(collData)
	if len(result) != 0 {
		t.Error(testedFuncName + "On empty array returned map with length " + string(len(result)))
	}

	// Array with one element
	collData = []model.CollectedData{}
	collData = append(collData, model.CollectedData{Ip: "1"})
	result = SeparateByIp(collData)
	if len(result) != 1 {
		t.Error(testedFuncName + "On one-element list returned map with len != 1")
	}

	if !CmpCollectedData(result["1"][0], model.CollectedData{Ip: "1"}){
		t.Error(testedFuncName + "On one-element list returned map with the element other than the input")
	}
	if len(result["1"]) != 1 {
		t.Error(testedFuncName+"On one-element list returned map having the length of the single value list != 1")
	}

	// Array with two identical elements
	collData = []model.CollectedData{}
	collData = append(collData, model.CollectedData{Ip: "1"})
	collData = append(collData, model.CollectedData{Ip: "1"})
	result = SeparateByIp(collData)
	if len(result) != 1 {
		t.Error(testedFuncName + "On two-element list (same elements) added more than one IP.")
	}
	if !CmpCollectedData(result["1"][0], model.CollectedData{Ip: "1"}){
		t.Error(testedFuncName + "On two-element list (same elements) the first element is not the same as the input")
	}
	if len(result["1"]) != 1{
		t.Error(testedFuncName+"On two-element list (same elements) did not add only one element")
	}

	// Array with two different elements and the same IP
	collData = []model.CollectedData{}
	collData = append(collData, model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"1"}}})
	collData = append(collData, model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"2"}}})
	result = SeparateByIp(collData)
	if len(result) != 1 {
		t.Error(testedFuncName + "On two-element list (different elements) added more than one IP.")
	}
	if !CmpCollectedData(result["1"][0], model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"1"}}}){
		t.Error(testedFuncName + "On two-element list (different elements) the first element is not the same as the first input")
	}
	if !CmpCollectedData(result["1"][1], model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"2"}}}){
		t.Error(testedFuncName + "On two-element list (different elements) the second element is not the same as the second input")
	}
	if len(result["1"]) != 2{
		t.Error(testedFuncName+"On two-element list (different elements) did not add two elements")
	}

	// Array with 2 same elements and 4 different
	collData = []model.CollectedData{}
	collData = append(collData, model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"1"}}})
	collData = append(collData, model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"1"}}})
	collData = append(collData, model.CollectedData{Ip: "2", Projects: []model.Project{{Id:"2"}}})
	collData = append(collData, model.CollectedData{Ip: "3", Projects: []model.Project{{Id:"3"}}})
	collData = append(collData, model.CollectedData{Ip: "4", Projects: []model.Project{{Id:"4"}}})
	collData = append(collData, model.CollectedData{Ip: "5", Projects: []model.Project{{Id:"5"}}})

	result = SeparateByIp(collData)
	if len(result) != 5 {
		t.Error(testedFuncName + "On list with 2 same elements and 4 different, added more or less than 5 IPs.")
	}
	if !CmpCollectedData(result["1"][0], model.CollectedData{Ip: "1", Projects: []model.Project{{Id:"1"}}}){
		t.Error(testedFuncName + "On list with 2 same elements and 4 different the first element is not the same as the first input")
	}
	if !CmpCollectedData(result["2"][0], model.CollectedData{Ip: "2", Projects: []model.Project{{Id:"2"}}}){
		t.Error(testedFuncName + "On list with 2 same elements and 4 different the second element is not the same as the third input")
	}
	if !CmpCollectedData(result["3"][0], model.CollectedData{Ip: "3", Projects: []model.Project{{Id:"3"}}}) {
		t.Error(testedFuncName + "On list with 2 same elements and 4 different the third element is not the same as the fourth input")
	}
	if !CmpCollectedData(result["4"][0], model.CollectedData{Ip: "4", Projects: []model.Project{{Id:"4"}}}){
		t.Error(testedFuncName + "On list with 2 same elements and 4 different the fourth element is not the same as the fifth input")
	}
	if !CmpCollectedData(result["5"][0], model.CollectedData{Ip: "5", Projects: []model.Project{{Id:"5"}}}){
		t.Error(testedFuncName + "On list with 2 same elements and 4 different the fifth element is not the same as the sixth input")
	}
	if len(result["1"]) != 1{
		t.Error(testedFuncName+"On list with 2 same elements and 4 different, did not add only one of the same input")
	}
	if len(result["2"]) != 1{
		t.Error(testedFuncName+"On list with 2 same elements and 4 different, did not add only one of the third input")
	}
	if len(result["3"]) != 1{
		t.Error(testedFuncName+"On list with 2 same elements and 4 different, did not add only one of the fourth input")
	}
	if len(result["4"]) != 1{
		t.Error(testedFuncName+"On list with 2 same elements and 4 different, did not add only one of the fifth input")
	}
	if len(result["5"]) != 1{
		t.Error(testedFuncName+"On list with 2 same elements and 4 different, did not add only one of the sixth input")
	}
}

func TestShowImagesInRegistries(t *testing.T) {
	t.Parallel()

	// nil data
	testedFuncName := "ShowImagesInRegistries: "
	result := ShowImagesInRegistries(nil, nil)
	if len(result) != 0 {
		t.Error(testedFuncName + "On empty input returned non-empty value")
	}

	// nil image, 1 registry
	idToRegistries := map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nil, idToRegistries)
	if len(result) != 0 {
		t.Error(testedFuncName + "On empty idToImages returned non-empty value")
	}

	// one image from public registry, nil registry
	nameToImages := map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", Location: "Public Registry"}
	result = ShowImagesInRegistries(nameToImages, nil)
	if len(result) != 1 {
		t.Error(testedFuncName + "one image from public registry was not added to the map")
	}
	if len(result["Public Registry"]) != 1 {
		t.Error(testedFuncName + "more than one image added")
	}
	if result["Public Registry"][0] != "alpine:latest" {
		t.Error(testedFuncName + "added wrong image")
	}

	// one image from private registry, nil registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest"}
	result = ShowImagesInRegistries(nameToImages, nil)
	if len(result) != 1 {
		t.Error(testedFuncName + "not one image from unidentifiable registry was added to the map")
	}
	if len(result["Unidentifiable Registry"]) != 1 {
		t.Error(testedFuncName + "not one image from unidentifiable registry was added to its list")
	}
	if result["Unidentifiable Registry"][0] != "alpine:latest" {
		t.Error(testedFuncName + "wrong image was added")
	}

	// one image from private registry, one different private registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", RegistryId: "2"}
	idToRegistries = map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nameToImages, idToRegistries)
	if len(result) != 1 {
		t.Error(testedFuncName + "not one image from unidentifiable private registry was added to the map")
	}
	if len(result["Unidentifiable Registry"]) != 1 {
		t.Error(testedFuncName + "not one image from unidentifiable registry was added to its list")
	}
	if result["Unidentifiable Registry"][0] != "alpine:latest" {
		t.Error(testedFuncName + "wrong image was added")
	}


	// one image from private registry, the same private registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", RegistryId: "1"}
	idToRegistries = map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nameToImages, idToRegistries)
	if len(result) != 1 {
		t.Error(testedFuncName + "not one image from the private registry was added to the map")
	}
	if len(result["name1(addr1)"]) != 1 {
		t.Error(testedFuncName + "not one image from the private registry was added to its list")
	}
	if result["name1(addr1)"][0] != "alpine:latest" {
		t.Error(testedFuncName + "wrong image was added")
	}


	// 3 images from public registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", Location: "Public Registry"}
	nameToImages["busybox:latest"] = model.Image{Id: "2", Name: "busybox", Tag:"latest", Location: "Public Registry"}
	nameToImages["hello-world:latest"] = model.Image{Id: "3", Name: "hello-world", Tag:"latest", Location: "Public Registry"}
	result = ShowImagesInRegistries(nameToImages, nil)
	if len(result) != 1 {
		t.Error(testedFuncName + "images were not added to the same (public) registry")
	}
	if len(result["Public Registry"]) != 3 {
		t.Error(testedFuncName + "not 3 images were added")
	}
	if !(result["Public Registry"][0] == "alpine:latest" || result["Public Registry"][1] == "alpine:latest" || result["Public Registry"][2] == "alpine:latest"){
		t.Error(testedFuncName + "alpine:latest wasn't added")
	}
	if !(result["Public Registry"][0] == "busybox:latest" || result["Public Registry"][1] == "busybox:latest" || result["Public Registry"][2] == "busybox:latest"){
		t.Error(testedFuncName + "busybox:latest wasn't added")
	}
	if !(result["Public Registry"][0] == "hello-world:latest" || result["Public Registry"][1] == "hello-world:latest" || result["Public Registry"][2] == "hello-world:latest"){
		t.Error(testedFuncName + "hello-world:latest wasn't added")
	}

	// 3 images from public registry, one irrelevant registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", Location: "Public Registry"}
	nameToImages["busybox:latest"] = model.Image{Id: "2", Name: "busybox", Tag:"latest", Location: "Public Registry"}
	nameToImages["hello-world:latest"] = model.Image{Id: "3", Name: "hello-world", Tag:"latest", Location: "Public Registry"}
	idToRegistries = map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nameToImages, idToRegistries)
	if len(result) != 1 {
		t.Error(testedFuncName + "images were not added to the same (public) registry")
	}
	if len(result["Public Registry"]) != 3 {
		t.Error(testedFuncName + "not 3 images were added")
	}
	if !(result["Public Registry"][0] == "alpine:latest" || result["Public Registry"][1] == "alpine:latest" || result["Public Registry"][2] == "alpine:latest"){
		t.Error(testedFuncName + "alpine:latest wasn't added")
	}
	if !(result["Public Registry"][0] == "busybox:latest" || result["Public Registry"][1] == "busybox:latest" || result["Public Registry"][2] == "busybox:latest"){
		t.Error(testedFuncName + "busybox:latest wasn't added")
	}
	if !(result["Public Registry"][0] == "hello-world:latest" || result["Public Registry"][1] == "hello-world:latest" || result["Public Registry"][2] == "hello-world:latest"){
		t.Error(testedFuncName + "hello-world:latest wasn't added")
	}

	// 3 images from the same private registry, one registry
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest", RegistryId: "1"}
	nameToImages["busybox:latest"] = model.Image{Id: "2", Name: "busybox", Tag:"latest", RegistryId: "1"}
	nameToImages["hello-world:latest"] = model.Image{Id: "3", Name: "hello-world", Tag:"latest", RegistryId: "1"}
	idToRegistries = map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nameToImages, idToRegistries)
	if len(result) != 1 {
		t.Error(testedFuncName + "images were not added to the same (public) registry")
	}
	if len(result["name1(addr1)"]) != 3 {
		t.Error(testedFuncName + "not 3 images were added")
	}
	if !(result["name1(addr1)"][0] == "alpine:latest" || result["name1(addr1)"][1] == "alpine:latest" || result["name1(addr1)"][2] == "alpine:latest"){
		t.Error(testedFuncName + "alpine:latest wasn't added")
	}
	if !(result["name1(addr1)"][0] == "busybox:latest" || result["name1(addr1)"][1] == "busybox:latest" || result["name1(addr1)"][2] == "busybox:latest"){
		t.Error(testedFuncName + "busybox:latest wasn't added")
	}
	if !(result["name1(addr1)"][0] == "hello-world:latest" || result["name1(addr1)"][1] == "hello-world:latest" || result["name1(addr1)"][2] == "hello-world:latest"){
		t.Error(testedFuncName + "hello-world:latest wasn't added")
	}

	// 1 unidentifiable registry image, 1 public registry image, 1 private registry image, 1 private registry (same as mentioned before)
	nameToImages = map[string]model.Image{}
	nameToImages["alpine:latest"] = model.Image{Id: "1", Name: "alpine", Tag:"latest"}
	nameToImages["busybox:latest"] = model.Image{Id: "2", Name: "busybox", Tag:"latest", Location: "Public Registry"}
	nameToImages["hello-world:latest"] = model.Image{Id: "3", Name: "hello-world", Tag:"latest", RegistryId: "1"}
	idToRegistries = map[string]model.Registry{}
	idToRegistries["1"] = model.Registry{Id:"1", Name: "name1", Addr: "addr1"}
	result = ShowImagesInRegistries(nameToImages, idToRegistries)
	if len(result) != 3 {
		t.Error(testedFuncName + "images were not added to 3 different registry maps")
	}
	if len(result["name1(addr1)"]) != 1 {
		t.Error(testedFuncName + "not 1 images was added to the private registry")
	}
	if len(result["Unidentifiable Registry"]) != 1 {
		t.Error(testedFuncName + "not 1 image was added to the unidentifiable registry")
	}
	if len(result["Public Registry"]) != 1 {
		t.Error(testedFuncName + "not 1 image was added to the public registry")
	}
	if result["Unidentifiable Registry"][0] != "alpine:latest" {
		t.Error(testedFuncName + "alpine:latest wasn't added to the unidentifiable registry")
	}
	if result["Public Registry"][0] != "busybox:latest" {
		t.Error(testedFuncName + "busybox:latest wasn't added to the public registry")
	}
	if result["name1(addr1)"][0] != "hello-world:latest" {
		t.Error(testedFuncName + "hello-world:latest wasn't added to the private registry")
	}
}

func TestCalculateNoOfVulnerabilitiesFound(t *testing.T) {
	t.Parallel()

	testedFuncName := "CalculateNoOfVulnerabilitesFound: "

	//nil data
	result := CalculateNoOfVulnerabilitiesFound(nil, nil)
	if len(result) != 0 {
		t.Error(testedFuncName + "returned something on nil data")
	}

	// 1 element list, nil
	dataList := []model.CollectedData{{Ip:"1"}}
	result = CalculateNoOfVulnerabilitiesFound(dataList, nil)
	if len(result) != 0 {
		t.Error(testedFuncName + "returned something on nil idToBuild")
	}

	//nil list, 1 element map
	idToBuild := map[string]model.Build{}
	idToBuild["1"] = model.Build{Id:"1"}
	result = CalculateNoOfVulnerabilitiesFound(nil, idToBuild)
	if len(result) != 0 {
		t.Error(testedFuncName + "returned data on empty dataList")
	}
}