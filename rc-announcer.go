package main

import (
	"fmt"
	"os"
)

func main() {
	var config struct {
		rcURL       string
		rcAuthToken string
		rcUserID    string
		rcUser      string
		rcUserPW    string
	}

	config.rcURL = os.Getenv("RC_URL")
	config.rcAuthToken = os.Getenv("RC_AUTH_TOKEN")
	config.rcUserID = os.Getenv("RC_USER_ID")
	config.rcUser = os.Getenv("RC_USER_NAME")
	config.rcUserPW = os.Getenv("RC_USER_PW")

	fmt.Println(config)

	if config.rcAuthToken == "" {
		getAuthToken(config.rcUser, config.rcUserPW, config.rcURL+"/api/v1/login")
	}

}
