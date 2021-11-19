package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// load list of api routes
func Load() []Route {
	routes := [][]Route{
		serverRoutes,
		quizRoutes,
		leaderboardRoutes,
	}

	var routesList []Route

	for _, r := range routes {
		routesList = append(routesList, r...)
	}

	return routesList
}

// setup routes to be handled
func SetupRoutes(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler)
	}

	return r
}
