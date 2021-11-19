package mocks

import "github.com/malcolm95/golang/go-quiz-api/models"

// mock quiz object
var Quiz = models.Quiz{
	Questions: []models.Question{
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
			Id:          2,
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
		{
			Id:          3,
			Description: "What is Malta's national dish?",
			Category:    "Culture",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Pastizzi",
				},
				{
					Id:          2,
					Description: "Rabbit stew",
				},
				{
					Id:          3,
					Description: "Minestra",
				},
				{
					Id:          4,
					Description: "Timpana",
				},
			},
			CorrectAnswerId: 2,
		},
		{
			Id:          4,
			Description: "Which singer achieved Malta's highest rank in the Eurovision Song Contest?",
			Category:    "Entertainment",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Lynn Chircop",
				},
				{
					Id:          2,
					Description: "Mary Spiteri",
				},
				{
					Id:          3,
					Description: "Ira Losco",
				},
				{
					Id:          4,
					Description: "Destiny Chukunyere",
				},
			},
			CorrectAnswerId: 3,
		},
		{
			Id:          5,
			Description: "Which Maltese football club won last year's BOV Premier League (20/21)?",
			Category:    "Sports",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Hibernians",
				},
				{
					Id:          2,
					Description: "Birkirkara",
				},
				{
					Id:          3,
					Description: "Valletta",
				},
				{
					Id:          4,
					Description: "Hamrun Spartans",
				},
			},
			CorrectAnswerId: 4,
		},
	},
	Leaderboard: models.Leaderboard{
		Quizzers: []models.Quizzer{
			{
				Id:    5,
				Name:  "Tony Montana",
				Score: 2,
			},
			{
				Id:    4,
				Name:  "Barney Simpson",
				Score: 2,
			},
			{
				Id:    6,
				Name:  "Harry Potter",
				Score: 4,
			},
			{
				Id:    2,
				Name:  "Ed Sheeran",
				Score: 6,
			},
			{
				Id:    1,
				Name:  "Brad Pitt",
				Score: 10,
			},
			{
				Id:    3,
				Name:  "Tom Hardy",
				Score: 10,
			},
			{
				Id:    7,
				Name:  "Michael Phelps",
				Score: 10,
			},
		},
	},
}

// mock next quizzer Id
var NextQuizzerId = 8
