package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type configuration struct {
	rcURL       string
	rcAuthToken string
	rcUserID    string
	rcUser      string
	rcUserPW    string
}

type loginTokens struct {
	AuthToken string `json:"authToken"`
	UserID    string `json:"userId"`
}
type loginResponse struct {
	Status string      `json:"status"`
	Data   loginTokens `json:"data"`
}

type loginCredentials struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

type message struct {
	Channel     string      `json:"channel"`
	Text        string      `json:"text"`
	Alias       string      `json:"alias"`
	Emoji       string      `json:"emoji"`
	Avatar      string      `json:"avatar"`
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

func main() {
	rocketchat := configuration{}
	rocketchat.loadConfig()

	if rocketchat.rcAuthToken == "" {
		log.Println("Unable to find AuthToken")
		rocketchat.getAuthToken()
	}

	// testMessage(config)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/announce/{channel}", config.AnnounceChannel)
	// router.HandleFunc("/grafana/{channel}", config.AnnounceGrafana)
	router := newRouter(rocketchat)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func testMessage(config configuration, channel string) {
	shame := attachment{
		ImageURL: "http://static1.squarespace.com/static/573e24aae707eba722970e2c/573e26b47c65e44fad49d6bf/5900b692e4fcb56229197a17/1532044418725/untitled.png?format=1500w",
		Title:    "Shame",
	}
	Attachements := attachments{shame}

	announcement := message{
		Channel:     channel,        // general
		Text:        `testing 123`,  // this is a test
		Alias:       "rc-announcer", // not_a robot
		Emoji:       "",             // :troll:
		Avatar:      "",             // <baseURL>/avatar/<username>
		Attachments: Attachements,
	}

	log.Printf("Sending announcement to \"%+v\" as \"%+v\"\n", announcement.Channel, announcement.Alias)
	resp, body := rcPost(config, `/api/v1/chat.postMessage`, announcement)
	var response messageResponse
	json.Unmarshal(body, &response)

	if response.Success != true {
		log.Printf("Response headers: %+v\n", resp.Header)
		log.Printf("Response body: %+v\n", string(body))
		log.Println("RocketChat announce failed")
	}
}

func (c *configuration) loadConfig() {
	c.rcURL = os.Getenv("RC_URL")
	c.rcAuthToken = os.Getenv("RC_AUTH_TOKEN")
	c.rcUserID = os.Getenv("RC_USER_ID")
	c.rcUser = os.Getenv("RC_USER_NAME")
	c.rcUserPW = os.Getenv("RC_USER_PW")

	log.Printf("Config on load: %+v\n", c)
}

func (c *configuration) getAuthToken() {
	// https://rocket.chat/docs/developer-guides/rest-api/authentication/login/
	log.Println("Trying to Login")
	credentials := loginCredentials{
		Username: c.rcUser,
		Password: c.rcUserPW,
	}
	resp, body := rcPost(*c, "/api/v1/login", credentials)

	var response loginResponse
	json.Unmarshal(body, &response)

	log.Println("Login status:", response.Status)
	log.Println("Login authToken:", response.Data.AuthToken)
	log.Println("Login userId:", response.Data.UserID)
	if response.Status != "success" {
		log.Printf("Response headers: %+v\n", resp.Header)
		log.Fatal("Unable to acquire new authentication tokens - exiting.")
	}

	c.rcAuthToken = response.Data.AuthToken
	c.rcUserID = response.Data.UserID
	err := os.Setenv("RC_AUTH_TOKEN", response.Data.AuthToken)
	if err != nil {
		log.Fatal("Unable to set AuthToken ENV", err)
	}
	err = os.Setenv("RC_USER_ID", response.Data.UserID)
	if err != nil {
		log.Fatal("Unable to set UserID ENV", err)
	}
}
