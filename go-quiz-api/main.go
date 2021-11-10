package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malcolm95/golang/go-quiz-api/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/questions", handlers.GetAllQuestions)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
