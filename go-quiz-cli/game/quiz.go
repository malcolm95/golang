package game

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	Models "github.com/malcolm95/golang/go-quiz-api/models"
	Config "github.com/malcolm95/golang/go-quiz-cli/config"
)

type Quiz struct {
	Questions []Models.Question
	Scores    []int
}

func Start() {
	// initialize new quiz
	quiz := Quiz{}

	// get questions
	quiz.Questions = Config.GetQuizQuestions()

	// play quiz rounds until player exits
	for {
		runPlayerRound(&quiz)

		var userAnswer string

		for {
			fmt.Printf("\nWould you like to exit the game? (Y/N): ")
			_, err := fmt.Scanln(&userAnswer)

			// validate user input
			if err == nil && (strings.ToUpper(userAnswer) == "Y" || strings.ToUpper(userAnswer) == "N") {
				if strings.ToUpper(userAnswer) == "Y" {
					return
				} else {
					break
				}
			}

			fmt.Println("Invalid answer! Please Y or N to exit the game or continue...")
		}
	}
}

// runs an iteration of the quiz for a player
func runPlayerRound(quiz *Quiz) {
	var playerScore int = 0

	fmt.Println("Welcome to a fun quiz about Malta!")
	fmt.Println("")

	for _, question := range quiz.Questions {
		displayQuestion(question)

		var userAnswerStr string
		var userAnswer int

		for {
			fmt.Println("")
			fmt.Printf("Enter an answer number: ")

			// read user input
			_, err := fmt.Scanln(&userAnswerStr)
			if err == nil {
				// convert user input to int
				userAnswer, err = strconv.Atoi(userAnswerStr)

				// validate user's answer
				if err == nil && userAnswer >= 1 && userAnswer <= 4 {
					break
				}
			}

			fmt.Println("Invalid answer! Please enter a number between 1 and 4...")
		}

		// check if correct answer
		if isCorrectAnswer(userAnswer, question) {
			playerScore++
		}

		fmt.Println("\n**************************")
		fmt.Println("")
	}

	// add score to quiz scores
	quiz.Scores = append(quiz.Scores, playerScore)
	sort.Ints(quiz.Scores)

	// calculate player ranking
	calculatePlayerRanking(quiz, playerScore)
}

// verifies whether the user answer is correct
func isCorrectAnswer(userAnswer int, question Models.Question) bool {
	return userAnswer == question.CorrectAnswerId
}

// calculates player's percentile rank based on other players' scores
func calculatePlayerRanking(quiz *Quiz, playerScore int) {
	var percentile float32

	if len(quiz.Scores) > 1 {
		pos := sort.SearchInts(quiz.Scores, playerScore)

		if pos > 0 {
			percentile = (float32(pos) / float32(len(quiz.Scores)-1)) * 100
		} else {
			percentile = 0
		}
	} else {
		percentile = 100
	}

	fmt.Println("You scored", playerScore, "out of", len(quiz.Questions))
	fmt.Println("You scored higher than", fmt.Sprintf("%.2f", percentile), "% of all other quizzers!")
}

// displays referenced question on screen
func displayQuestion(question Models.Question) {
	fmt.Println("Category:", question.Category)
	fmt.Println("Question:", question.Description)
	fmt.Println("--------------------------")
	fmt.Println("Answers")
	fmt.Println(question.Answers[0].Id, question.Answers[0].Description)
	fmt.Println(question.Answers[1].Id, question.Answers[1].Description)
	fmt.Println(question.Answers[2].Id, question.Answers[2].Description)
	fmt.Println(question.Answers[3].Id, question.Answers[3].Description)
}
