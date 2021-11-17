package models

type Leaderboard struct {
	Quizzers []Quizzer `json:"quizzers"`
}
