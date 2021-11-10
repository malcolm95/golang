package models

type Question struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Answers     []Answer `json:"answers"`
}

type Answer struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCorrect   bool   `json:"isCorrect"`
}
