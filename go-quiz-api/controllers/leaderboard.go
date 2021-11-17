package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/malcolm95/golang/go-quiz-api/mocks"
	"github.com/malcolm95/golang/go-quiz-api/models"
)

// submit quizzer to leaderboard
func SubmitQuizzer(quizzer models.Quizzer) {
	// add quizzer to leaderboard
	mocks.Quiz.Leaderboard.Quizzers = append(mocks.Quiz.Leaderboard.Quizzers, quizzer)

	// sort leaderboard
	sortQuizzersByScore()
}

// get quizzer leaderboard percentile
func GetQuizzerLeaderboardPercentile(w http.ResponseWriter, r *http.Request) {
	quizzerScoreParam := r.URL.Query().Get("quizzerScore")
	if quizzerScoreParam == "" {
		http.Error(w, "Quizzer score not found!", http.StatusNotFound)
		return
	}

	quizzerScore, err := strconv.Atoi(quizzerScoreParam)
	if err != nil {
		http.Error(w, "Invalid quizzer score!", http.StatusNotFound)
		return
	}

	var percentile float32

	if len(mocks.Quiz.Leaderboard.Quizzers) > 1 {
		sameQuizzerScoreIndex := 0

		// get index of first quizzer with equal score
		for index, q := range mocks.Quiz.Leaderboard.Quizzers {
			// store index of matching score
			if quizzerScore == q.Score {
				sameQuizzerScoreIndex = index
			}
		}

		if sameQuizzerScoreIndex > 0 {
			// calculate percentile of quizzers score when compared to all OTHER quizzers' scores
			percentile = (float32(sameQuizzerScoreIndex) / float32(len(mocks.Quiz.Leaderboard.Quizzers)-1)) * 100
		} else {
			// return 0% if last position
			percentile = 0
		}
	} else {
		// return 100% if first entry
		percentile = 100
	}

	// return response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%.2f", percentile)))
}

func sortQuizzersByScore() {
	// sorts leaderboard quizzers by score in ascending order to facilitate percentile calculations
	sort.SliceStable(mocks.Quiz.Leaderboard.Quizzers, func(i, j int) bool {
		return mocks.Quiz.Leaderboard.Quizzers[i].Score < mocks.Quiz.Leaderboard.Quizzers[j].Score
	})
}
