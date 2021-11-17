package routes

import (
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/controllers"
)

var quizRoutes = []Route{
	{
		URI:     "/getAllQuestions",
		Method:  http.MethodGet,
		Handler: controllers.GetAllQuestions,
	},
	{
		URI:     "/postQuizzerSubmission",
		Method:  http.MethodPost,
		Handler: controllers.PostQuizzerSubmission,
	},
}
