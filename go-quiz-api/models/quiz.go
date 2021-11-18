package models

type Quiz struct {
	Questions   []Question  `json:"questions"`
	Leaderboard Leaderboard `json:"leaderboard"`
}

type QuizSubmission struct {
	Quizzer Quizzer `json:"quizzer"`
	Answers []int   `json:"answers"`
}
