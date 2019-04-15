package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getAuthToken(u string, p string, url string) {
	authValues := map[string]string{
		"username": u,
		"password": p,
	}

	jsonAuthValues, _ := json.Marshal(authValues)
	//fmt.Println(string(jsonAuthValues))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonAuthValues))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
