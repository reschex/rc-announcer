package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(rocketchat configuration) *mux.Router {
	var routes = Routes{
		Route{
			"RequestEcho",
			"POST",
			"/",
			Index,
		},
		Route{
			"Grafana",
			"POST",
			"/grafana/{channel}",
			rocketchat.AnnounceGrafana,
		},
		Route{
			"test",
			"POST",
			"/announce/{channel}",
			rocketchat.AnnounceChannel,
		},
	}
	// To Do: Add route for channel list:
	// curl --insecure https://URL/api/v1/channels.list.joined -H "X-Auth-Token: ${RC_AUTH_TOKEN}" -H "X-User-Id: ${RC_USER_ID}"
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
