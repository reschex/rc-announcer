package main

import (
	"fmt"
	"log"
	"os"
)

type configuration struct {
	rcURL       string
	rcAuthToken string
	rcUserID    string
	rcUser      string
	rcUserPW    string
}

func main() {
	config := configuration{}
	config.loadConfig()

	if config.rcAuthToken == "" {
		config.rcAuthToken, config.rcUserID = getAuthToken(config.rcUser, config.rcUserPW, config.rcURL)
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

	shame := attachment{
		ImageURL: "http://static1.squarespace.com/static/573e24aae707eba722970e2c/573e26b47c65e44fad49d6bf/5900b692e4fcb56229197a17/1532044418725/untitled.png?format=1500w",
		Title:    "Shame",
	}
	Attachements := attachments{shame}

	announcement := message{
		Channel:     "", // general
		Text:        ``, // this is a test
		Alias:       "", // not_a robot
		Emoji:       "", // :troll:
		Attachments: Attachements,
	}

	announcement.send(config.rcAuthToken, config.rcUserID, config.rcURL)
}

func (c *configuration) loadConfig() {
	c.rcURL = os.Getenv("RC_URL")
	c.rcAuthToken = os.Getenv("RC_AUTH_TOKEN")
	c.rcUserID = os.Getenv("RC_USER_ID")
	c.rcUser = os.Getenv("RC_USER_NAME")
	c.rcUserPW = os.Getenv("RC_USER_PW")

	fmt.Printf("config on load: %+v", c)
}
