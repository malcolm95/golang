package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/mocks"
	"github.com/malcolm95/golang/go-quiz-api/models"
)

// returns all mock questions in json format
func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Quiz.Questions)
}

// submit quizzer list of answers and retrieve score
func PostQuizzerSubmission(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse request body!", http.StatusBadRequest)
		return
	}

	quizzer := models.Quizzer{}
	err = json.Unmarshal(body, &quizzer)
	if err != nil {
		http.Error(w, "Failed to parse quizzer info!", http.StatusBadRequest)
		return
	}

	answers := []int{}
	err = json.Unmarshal(body, &answers)
	if err != nil {
		http.Error(w, "Failed to parse quizzer answers!", http.StatusBadRequest)
		return
	}

	// get next quizzer id
	quizzer.Id = getNextQuizzerId()

	// get submission score
	quizzer.Score = getQuizzerSubmissionScore(answers)

	// return response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quizzer)
}

func getNextQuizzerId() int {
	currentNextQuizzerId := mocks.NextQuizzerId
	mocks.NextQuizzerId++

	return currentNextQuizzerId
}

func getQuizzerSubmissionScore(answers []int) int {
	score := 0

	for index, question := range mocks.Quiz.Questions {
		if question.CorrectAnswerId == answers[index] {
			score++
		}
	}

	return score
}
