package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//https://rocket.chat/docs/developer-guides/rest-api/authentication/login/
type loginTokens struct {
	AuthToken string `json:"authToken"`
	UserID    string `json:"userId"`
}
type loginResponse struct {
	Status string      `json:"status"`
	Data   loginTokens `json:"data"`
}

func getAuthToken(u string, p string, URL string) (string, string) {

	authValues := map[string]string{
		"user":     u,
		"password": p,
	}

	jsonAuthValues, _ := json.Marshal(authValues)
	// fmt.Println("using map:", string(jsonAuthValues))

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonAuthValues))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// fmt.Println("response Body:", string(body))

	var authResponse loginResponse
	json.Unmarshal(body, &authResponse)

	fmt.Println("API status:", authResponse.Status)
	fmt.Println("authToken:", authResponse.Data.AuthToken)
	fmt.Println("userId:", authResponse.Data.UserID)
	if authResponse.Status != "success" {
		log.Fatal("Unable to aquire new authentication tokens - exiting.")
	}
	return authResponse.Data.AuthToken, authResponse.Data.UserID
}
