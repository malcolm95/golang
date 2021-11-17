package models

type Quiz struct {
	Questions   []Question  `json:"questions"`
	Leaderboard Leaderboard `json:"leaderboard"`
}
