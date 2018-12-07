package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRoutes(path string, payload []byte) error {
	path = path + "/routes"
	res, err := http.Post(path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	code := res.StatusCode
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body := string(byteBody)
	fmt.Printf("[Status: %d] %s", code, body)
	return nil
}

func MakeServices(path string, payload []byte) error {
	path = path + "/services"
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	code := resp.StatusCode
	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body := string(byteBody)
	fmt.Printf("[Status: %d] %s", code, body)

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
