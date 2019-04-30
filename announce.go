package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type message struct {
	Channel     string      `json:"channel"`
	Text        string      `json:"text"`
	Alias       string      `json:"alias"`
	Emoji       string      `json:"emoji"`
	Attachments attachments `json:"attachments"`
}

type attachments []attachment

type attachment struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
}

type messageResponse struct {
	Success bool `json:"success"`
}

func (m message) send(authToken string, userID string, URL string) bool {
	jsonMessage, _ := json.Marshal(m)
	fmt.Println("message json:", string(jsonMessage))

	req, err := http.NewRequest("POST", URL+"/api/v1/chat.postMessage", bytes.NewBuffer(jsonMessage))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("X-User-Id", userID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))

	var response messageResponse
	json.Unmarshal(body, &response)

	fmt.Println("response json success:", response.Success)

	if response.Success != true {
		log.Fatal("RocketChat announce failed")
	}
	return response.Success
}
