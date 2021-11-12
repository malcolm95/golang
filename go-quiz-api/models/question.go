package models

type Question struct {
	Id          		int      	`json:"id"`
	Description 		string   	`json:"description"`
	Category    		string   	`json:"category"`
	Answers     		[]Answer 	`json:"answers"`
	CorrectAnswerId 	int			`json:"correctAnswerId"`			
}

type Answer struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}