package routes

import (
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/controllers"
)

var questionRoutes = []Route{
	{
		URI:     "/getAllQuestions",
		Method:  http.MethodGet,
		Handler: controllers.GetAllQuestions,
	},
}
