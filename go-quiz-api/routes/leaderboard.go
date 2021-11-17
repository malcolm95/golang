package routes

import (
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/controllers"
)

var leaderboardRoutes = []Route{
	{
		URI:     "/getQuizzerLeaderboardPercentile",
		Method:  http.MethodGet,
		Handler: controllers.GetQuizzerLeaderboardPercentile,
	},
}
