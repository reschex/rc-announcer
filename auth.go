package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getAuthToken(u string, p string, URL string) {

	authValues := map[string]string{
		"username": u,
		"password": p,
	}

	jsonAuthValues, _ := json.Marshal(authValues)
	//fmt.Println(string(jsonAuthValues))

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonAuthValues))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var response map[string]interface{}

	json.Unmarshal(body, &response)
	fmt.Println("message:", response["message"])

	// {
	// 	"status": "success",
	// 	"data": {
	// 	  "authToken": "gEQV-hPRIVATEgArGPFn_bTHISkISxAiFAKEmKEYi",
	// 	  "userId": "8ALSO9FAKE3adXz6x"
	// 	}
	//   }
}
