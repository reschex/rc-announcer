package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func newRouter(rocketchat configuration) *mux.Router {
	var routes = routes{
		route{
			"RequestEcho",
			"POST",
			"/",
			index,
		},
		route{
			"Grafana",
			"POST",
			"/grafana/{channel}",
			rocketchat.AnnounceGrafana,
		},
		route{
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
