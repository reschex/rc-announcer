package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

type evalMatches []evalMatch

type evalMatch struct {
	Value  float64 `json:"value"`
	Metric string  `json:"metric"`
}

type grafanaAlert struct {
	EvalMatches evalMatches `json:"evalMatches"`
	ImageURL    string      `json:"imageUrl"`
	Message     string      `json:"message"`
	RuleName    string      `json:"ruleName"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func (config configuration) AnnounceGrafana(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel := vars["channel"]

	var alert grafanaAlert

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&alert); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // if we can't decode json, return http 422 - unprocessable entity error
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println(err)
		}
	}

	log.Printf("received alert: %+v\n", alert)
	announcement := alert.ConvertToMessage(channel)

	log.Printf("Sending announcement to \"%+v\" as \"%+v\"\n", announcement.Channel, announcement.Alias)
	resp, body := rcPost(config, `/api/v1/chat.postMessage`, announcement)
	CheckResponse(resp, body)
}

func (config configuration) AnnounceChannel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel := vars["channel"]

	message := message{
		Channel: channel,
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // if we can't decode json, return http 422 - unprocessable entity error
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println(err)
		}
	}
	if message.Text != "" {
		log.Printf("Received message: %+v\n", message)
		resp, body := rcPost(config, `/api/v1/chat.postMessage`, message)
		CheckResponse(resp, body)
	} else {
		log.Println("Received empty or unparsable message.")
	}
}

func (alert grafanaAlert) ConvertToMessage(channel string) message {
	var emoji string

	Attachments := attachments{}

	if alert.ImageURL != "" {
		image := attachment{
			ImageURL: alert.ImageURL,
			Title:    alert.Title,
		}
		Attachments = attachments{image}
	}

	text := "*" + alert.Title + "*" + "\n" + alert.Message

	for _, em := range alert.EvalMatches {
		text = text + "\nMetric: " + em.Metric + "\nValue: " + fmt.Sprintf("%.3f", em.Value)
	}

	text = text + "\nStatus: " + alert.State

	if alert.State == "ok" {
		emoji = ":ballot_box_with_check:"
	} else {
		emoji = ":prometheus:"
	}
	announcement := message{
		Channel:     channel,         // general
		Text:        text,            // this is a test
		Alias:       "Grafana-Alert", // not_a robot
		Emoji:       emoji,           // :troll:
		Avatar:      "",              // <baseURL>/avatar/<username>
		Attachments: Attachments,
	}
	return announcement
}

func CheckResponse(r *http.Response, b []byte) {
	var respBody messageResponse
	json.Unmarshal(b, &respBody)
	if r.StatusCode != 200 || respBody.Success != true {
		log.Printf("Response Headers: 	%+v\n", r.Header)
		log.Printf("Response Body: 		%+v\n", string(b))

	}
}
