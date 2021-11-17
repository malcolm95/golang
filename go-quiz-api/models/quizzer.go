package models

type Quizzer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Score   int    `json:"score"`
}
