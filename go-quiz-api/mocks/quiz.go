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
					Description: "Fabrizio Faniello",
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
		{
			Id:          6,
			Description: "Which Robin Williams movie was filmed in Malta?",
			Category:    "Entertainment",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Moscow On The Hudson",
				},
				{
					Id:          2,
					Description: "Popeye",
				},
				{
					Id:          3,
					Description: "The World According To Garp",
				},
				{
					Id:          4,
					Description: "Dead Poet's Society",
				},
			},
			CorrectAnswerId: 2,
		},
		{
			Id:          7,
			Description: "In which year did Malta become an indepedent country?",
			Category:    "Entertainment",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "1974",
				},
				{
					Id:          2,
					Description: "1959",
				},
				{
					Id:          3,
					Description: "1964",
				},
				{
					Id:          4,
					Description: "1982",
				},
			},
			CorrectAnswerId: 3,
		},
		{
			Id:          8,
			Description: "Which other language is Malta's other official language besides Maltese?",
			Category:    "Culture",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Italian",
				},
				{
					Id:          2,
					Description: "French",
				},
				{
					Id:          3,
					Description: "Arabic",
				},
				{
					Id:          4,
					Description: "English",
				},
			},
			CorrectAnswerId: 4,
		},
		{
			Id:          9,
			Description: "Which symbol is found on the Maltese flag?",
			Category:    "Culture",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Star",
				},
				{
					Id:          2,
					Description: "Cross",
				},
				{
					Id:          3,
					Description: "Heart",
				},
				{
					Id:          4,
					Description: "Leaf",
				},
			},
			CorrectAnswerId: 2,
		},
		{
			Id:          10,
			Description: "Between which two continents is Malta located?",
			Category:    "Geography",
			Answers: []models.Answer{
				{
					Id:          1,
					Description: "Europe and Africa",
				},
				{
					Id:          2,
					Description: "Africa and Asia",
				},
				{
					Id:          3,
					Description: "Europe and Asia",
				},
				{
					Id:          4,
					Description: "Asia and Australia",
				},
			},
			CorrectAnswerId: 1,
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
				Score: 7,
			},
			{
				Id:    3,
				Name:  "Tom Hardy",
				Score: 9,
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
