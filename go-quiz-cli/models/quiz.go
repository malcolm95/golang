package models

import (
	Models "github.com/malcolm95/golang/go-quiz-api/models"
)

type Quiz struct {
	Questions []Models.Question
	Scores    []int
}

func start(quiz *Quiz) {
	// get questions
	quiz.Questions = Config.getQuizQuestions()
}
