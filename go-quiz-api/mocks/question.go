package mocks

import (
	"github.com/malcolm95/golang/go-quiz-api/models"
)

var Questions = []models.Question{
	{
		Id:          1,
		Description: "What is the capital city of Malta?",
		Category:    "Geography",
		Answers: []models.Answer{
			{
				Id:          1,
				Description: "Valletta",
				IsCorrect:   true,
			},
			{
				Id:          2,
				Description: "Imdina",
				IsCorrect:   false,
			},
			{
				Id:          3,
				Description: "Mosta",
				IsCorrect:   false,
			},
			{
				Id:          3,
				Description: "Birkirkara",
				IsCorrect:   false,
			},
		},
	},
	{
		Id:          1,
		Description: "In which year dId the Great Siege happen in Malta?",
		Category:    "History",
		Answers: []models.Answer{
			{
				Id:          1,
				Description: "1560",
				IsCorrect:   false,
			},
			{
				Id:          2,
				Description: "1635",
				IsCorrect:   false,
			},
			{
				Id:          3,
				Description: "1497",
				IsCorrect:   false,
			},
			{
				Id:          3,
				Description: "1565",
				IsCorrect:   true,
			},
		},
	},
}
