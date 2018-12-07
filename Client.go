package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRoutes(path string) error {
	res, err := http.Get("http://localhost:8001/routes/523ceb34-d8a5-4e19-81d7-7a6dd3e60c2a")
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	service := Routes{}
	jsonErr := json.Unmarshal(body, &service)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Println(service)
	return nil
}

func MakeServices(path string, payload []byte) error {
	path = path + "/services"
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	temp := Service{}
	json.NewDecoder(resp.Body).Decode(&temp)
	log.Println(temp)
	return nil
}

func getRoutes(path string, routes *[]Routes) error {
	path = path + "/routes"
	res, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return readErr
	}

	tempStruct := struct {
		Data []Routes `json:"data"`
	}{}

	jsonErr := json.Unmarshal(body, &tempStruct)
	if jsonErr != nil {
		return jsonErr
	}

	*routes = tempStruct.Data

	return nil
}

func getServices(path string, services *[]Service) error {
	path = path + "/services"
	res, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		return err

	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return readErr
	}

	tempStruct := struct {
		Data []Service `json:"data"`
	}{}

	jsonErr := json.Unmarshal(body, &tempStruct)
	if jsonErr != nil {
		return jsonErr
	}
	*services = tempStruct.Data
	return nil
}
