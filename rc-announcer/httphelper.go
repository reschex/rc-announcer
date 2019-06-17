package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func rcPost(c configuration, endpoint string, p interface{}) (*http.Response, []byte) {
	payload, _ := json.Marshal(p)
	req, err := http.NewRequest("POST", c.rcURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.rcAuthToken)
	req.Header.Set("X-User-Id", c.rcUserID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	log.Println("Response Status:	 ", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// fmt.Println("response Body:", string(body))

	return resp, body

}
