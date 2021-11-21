package game

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	Models "github.com/malcolm95/golang/go-quiz-api/models"
	Config "github.com/malcolm95/golang/go-quiz-cli/config"
)

func Start() {
	// get quiz questions
	quizQuestions := Config.GetQuizConfiguration()

	fmt.Println("*******************************************************")
	fmt.Println("  Welcome to Maltese Trivia - A fun quiz about Malta!  ")
	fmt.Println("*******************************************************")

	// loop through main menu until user exits
	for {
		// display main menu
		menuOption := displayMainMenu()

		fmt.Println("")

		switch menuOption {
		case 1:
			runQuizzerRound(quizQuestions)
		case 2:
			showLeaderboard()
		case 3:
			exitQuiz()
		}
	}
}

// display main menu and store quizzer's menu option choice
func displayMainMenu() int {
	fmt.Println("\n\nMain Menu")
	fmt.Println("--------------------")
	fmt.Println("1 Play quiz")
	fmt.Println("2 View leaderboard")
	fmt.Println("3 Exit")

	var menuOptionStr string
	var menuOption int

	for {
		fmt.Printf("\nEnter an option: ")

		// read user input
		_, err := fmt.Scanln(&menuOptionStr)
		if err == nil {
			// convert user input to int
			menuOption, err = strconv.Atoi(menuOptionStr)

			// validate user's answer
			if err == nil && menuOption >= 1 && menuOption <= 3 {
				break
			}

			fmt.Println("Invalid option! Please enter a valid menu option between 1 and 3...")
		}
	}

	return menuOption
}

// run one quiz round and return quizzer's performance
func runQuizzerRound(quizQuestions []Models.Question) {
	// play a quiz round
	quizzer, quizzerAnswers := playQuizRound(quizQuestions)

	// submit quizzer answers
	quizzer = submitAnswers(quizzer, quizzerAnswers)

	// get quizzer result
	getQuizzerLeaderboardPercentile(quizzer)
}

// runs an iteration of the quiz for a player and return quizzer + list of answers
func playQuizRound(quizQuestions []Models.Question) (Models.Quizzer, []int) {
	var quizzerName string
	var err error

	for {
		fmt.Printf("Enter your quizzer name: ")

		// read user input
		in := bufio.NewReader(os.Stdin)
		quizzerName, err = in.ReadString('\n')

		if err == nil && len(quizzerName) > 0 {
			break
		}

		fmt.Println("Invalid quizzer name! Please enter a suitable name...")
	}

	// initialize new quizzer
	quizzer := Models.Quizzer{
		Id:    0,
		Name:  strings.TrimSpace(quizzerName),
		Score: 0,
	}

	fmt.Println("")

	userAnswers := []int{}

	for _, question := range quizQuestions {
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

		// add user answer to list of answers
		userAnswers = append(userAnswers, userAnswer)

		fmt.Println("\n**************************")
		fmt.Println("")
	}

	return quizzer, userAnswers
}

// displays referenced question on screen
func displayQuestion(question Models.Question) {
	fmt.Println("Category:", question.Category)
	fmt.Println("Question:", question.Description)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Answers")
	fmt.Println(question.Answers[0].Id, question.Answers[0].Description)
	fmt.Println(question.Answers[1].Id, question.Answers[1].Description)
	fmt.Println(question.Answers[2].Id, question.Answers[2].Description)
	fmt.Println(question.Answers[3].Id, question.Answers[3].Description)
}

func submitAnswers(quizzer Models.Quizzer, quizzerAnswers []int) Models.Quizzer {
	// construct quiz submission
	quizSubmission := Models.QuizSubmission{
		Quizzer: quizzer,
		Answers: quizzerAnswers,
	}

	// encode post body
	data, err := json.Marshal(quizSubmission)

	if err != nil {
		processError("Encoding of quiz submission failed", err)
	}

	requestRoute := "/postQuizzerSubmission"

	// construct post request
	response, err := http.Post(
		Config.ServerURL+requestRoute,
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		processError("Request to post quizzer submission failed", err)
	}

	// read response
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		processError("Failed to read response from quizzer submission request", err)
	}

	if err := json.Unmarshal(responseBytes, &quizzer); err != nil {
		processError("Could not unmarshal response", err)
	}

	return quizzer
}

func getQuizzerLeaderboardPercentile(quizzer Models.Quizzer) {
	requestRoute := "/getQuizzerLeaderboardPercentile"

	// construct request
	request, err := http.NewRequest(
		http.MethodGet,
		Config.ServerURL+requestRoute,
		nil,
	)

	if err != nil {
		processError("Creation of request for quizzer ranking failed", err)
	}

	// add accept header
	request.Header.Add("Accept", "application/json")

	// add request parameters
	queryValues := request.URL.Query()
	queryValues.Add("quizzerScore", strconv.Itoa(quizzer.Score))
	request.URL.RawQuery = queryValues.Encode()

	// send request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		processError("Request for quizzer ranking failed", err)
	}

	// read response
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		processError("Failed to read response from quizzer ranking request", err)
	}

	// print quizzer results
	fmt.Println("You scored", quizzer.Score, "out of 10 and did better than", string(responseBytes), "% of the other quizzers!")
}

// retrieves current quiz leaderboard and displays contents on screen
func showLeaderboard() {
	requestRoute := "/getLeaderboard"

	// construct request
	request, err := http.NewRequest(
		http.MethodGet,
		Config.ServerURL+requestRoute,
		nil,
	)

	if err != nil {
		processError("Creation of request for leaderboard failed", err)
	}

	// add accept header
	request.Header.Add("Accept", "application/json")

	// send request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		processError("Request for leaderboard failed", err)
	}

	// read response
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		processError("Failed to read response from quiz questions request", err)
	}

	leaderboard := Models.Leaderboard{}

	if err := json.Unmarshal(responseBytes, &leaderboard); err != nil {
		processError("Could not unmarshal leaderboard response", err)
	}

	// display leaderboard contents
	displayLeaderboard(leaderboard)
}

// prints quiz leaderboard on screen
func displayLeaderboard(leaderboard Models.Leaderboard) {
	fmt.Println("LEADERBOARD")
	fmt.Println("-----------")
	fmt.Println("")

	for i := len(leaderboard.Quizzers) - 1; i >= 0; i-- {
		fmt.Println(leaderboard.Quizzers[i].Name, "-", leaderboard.Quizzers[i].Score, "points")
	}
}

// process and display error
func processError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err)
	os.Exit(2)
}

func exitQuiz() {
	fmt.Println("Hope you enjoyed the quiz! Goodbye :)")
	os.Exit(0)
}
