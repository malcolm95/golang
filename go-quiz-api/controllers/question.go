package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/malcolm95/golang/go-quiz-api/mocks"
)

// returns all mock questions in json format
func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Questions)
}
