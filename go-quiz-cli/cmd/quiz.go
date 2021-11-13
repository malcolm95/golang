/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/spf13/cobra"

	Models "github.com/malcolm95/golang/go-quiz-api/models"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var questions = setupQuiz()

		var scores []int

		for {
			var correctScore int = 0

			for _, question := range questions {
				fmt.Println("Category:", question.Category)
				fmt.Println("Question:", question.Description)
				fmt.Println("--------------------------")
				fmt.Println("Answers")
				fmt.Println(question.Answers[0].Id, question.Answers[0].Description)
				fmt.Println(question.Answers[1].Id, question.Answers[1].Description)
				fmt.Println(question.Answers[2].Id, question.Answers[2].Description)
				fmt.Println(question.Answers[3].Id, question.Answers[3].Description)

				fmt.Println("")
				fmt.Printf("Answer: ")

				var userAnswer int
				fmt.Scanln(&userAnswer)

				if userAnswer == question.CorrectAnswerId {
					correctScore++
				}

				fmt.Println("\n**************************")
				fmt.Println("")
			}

			scores = append(scores, correctScore)
			sort.Ints(scores)

			var percentile float32

			if len(scores) > 1 {
				pos := sort.SearchInts(scores, correctScore)

				fmt.Println("Current Position:", pos)

				if pos > 0 {
					percentile = (float32(pos) / float32(len(scores)-1)) * 100
				} else {
					percentile = 0
				}
			} else {
				percentile = 100
			}

			fmt.Println("Your score is", correctScore, "out of", len(questions))
			fmt.Println("You scored higher than", fmt.Sprintf("%.2f", percentile), "% of all quizzers")
		}
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}

func setupQuiz() []Models.Question {
	requestURL := "http://localhost:4000/questions"
	responseBytes := getQuizQuestions(requestURL)

	questions := []Models.Question{}

	if err := json.Unmarshal(responseBytes, &questions); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	return questions
}

func getQuizQuestions(requestURL string) []byte {
	fmt.Println("Getting questions...\n")

	request, err := http.NewRequest(
		http.MethodGet,
		requestURL,
		nil,
	)

	if err != nil {
		log.Printf("Request for quiz questions failed - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Go Quiz (github.com/malcolm95/golang/go-quiz-cli")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Request for quiz questions failed - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response from quiz questions request - %v", err)
	}

	return responseBytes
}
