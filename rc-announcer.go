package main

import (
	"fmt"
	"log"
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

	fmt.Println("config on load:", config)

	if config.rcAuthToken == "" {
		config.rcAuthToken, config.rcUserID = getAuthToken(config.rcUser, config.rcUserPW, config.rcURL+"/api/v1/login")
		err := os.Setenv("RC_AUTH_TOKEN", config.rcAuthToken)
		if err != nil {
			log.Fatal("Unable to set AuthToken ENV", err)
		}
		err = os.Setenv("RC_USER_ID", config.rcAuthToken)
		if err != nil {
			log.Fatal("Unable to set UserID ENV", err)
		}
		fmt.Println("config after getAuthToken:", config)
	}

}
