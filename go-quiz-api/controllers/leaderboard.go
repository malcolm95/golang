package controllers

import (
	"encoding/json"
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

// returns current leaderboard
func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	// return leaderboard
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Quiz.Leaderboard)
}

// get quizzer leaderboard percentile
func GetQuizzerLeaderboardPercentile(w http.ResponseWriter, r *http.Request) {
	// fetch quizzer score within request
	quizzerScoreParam := r.URL.Query().Get("quizzerScore")
	if quizzerScoreParam == "" {
		http.Error(w, "Quizzer score not found!", http.StatusNotFound)
		return
	}

	// parse quizzer score
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
				break
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

	// return calculated percentile
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%.2f", percentile)))
}

// sorts quizzers in leaderboard by score in ascending order to facilitate percentile calculations
func sortQuizzersByScore() {
	sort.SliceStable(mocks.Quiz.Leaderboard.Quizzers[:], func(i, j int) bool {
		return mocks.Quiz.Leaderboard.Quizzers[i].Score < mocks.Quiz.Leaderboard.Quizzers[j].Score
	})
}
