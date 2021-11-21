package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"

	Models "github.com/malcolm95/golang/go-quiz-api/models"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

var ServerURL string

// return API server address
func getServerAddress(cfg *Config) string {
	return "http://" + cfg.Server.Host + ":" + cfg.Server.Port
}

// load app config file
func loadConfig() *Config {
	var cfg Config
	readConfigFile(&cfg)

	return &cfg
}

// process and display error
func processError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err)
	os.Exit(2)
}

// read and decode app config file
func readConfigFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError("Failed to read config file", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError("Failed to decode config file", err)
	}
}

// retrieve quiz config including quiz questions
func GetQuizConfiguration() []Models.Question {
	// load config file
	config := loadConfig()

	// set API server URL
	ServerURL = getServerAddress(config)

	// construct request URL
	requestURL := ServerURL

	// set request route
	requestRoute := "/getAllQuestions"

	// construct request
	request, err := http.NewRequest(
		http.MethodGet,
		requestURL+requestRoute,
		nil,
	)

	if err != nil {
		processError("Creation of request for quiz questions failed", err)
	}

	// add accept header
	request.Header.Add("Accept", "application/json")

	// send request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		processError("Request for quiz questions failed - please make sure API server is running!", err)
	}

	// read response
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		processError("Failed to read response from quiz questions request", err)
	}

	quizQuestions := []Models.Question{}

	// parse response byte array into list of quiz questions
	if err := json.Unmarshal(responseBytes, &quizQuestions); err != nil {
		processError("Could not unmarshal response", err)
	}

	return quizQuestions
}
