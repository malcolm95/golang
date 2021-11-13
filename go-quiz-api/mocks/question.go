package mocks

import (
	"github.com/malcolm95/golang/go-quiz-api/models"
)

var quizQuestions = []models.Question{
	{
		Id:          1,
		Description: "What is the capital city of Malta?",
		Category:    "Geography",
		Answers: []models.Answer{
			{
				Id:          1,
				Description: "Valletta",
			},
			{
				Id:          2,
				Description: "Imdina",
			},
			{
				Id:          3,
				Description: "Mosta",
			},
			{
				Id:          4,
				Description: "Birkirkara",
			},
		},
		CorrectAnswerId: 1,
	},
	{
		Id:          1,
		Description: "In which year did the Great Siege happen in Malta?",
		Category:    "History",
		Answers: []models.Answer{
			{
				Id:          1,
				Description: "1560",
			},
			{
				Id:          2,
				Description: "1635",
			},
			{
				Id:          3,
				Description: "1497",
			},
			{
				Id:          4,
				Description: "1565",
			},
		},
		CorrectAnswerId: 4,
	},
}
