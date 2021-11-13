package routes

import (
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/controllers"
)

var serverRoutes = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: controllers.ServerRunning,
	},
}
