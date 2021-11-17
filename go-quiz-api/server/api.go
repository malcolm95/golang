package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malcolm95/golang/go-quiz-api/config"
	"github.com/malcolm95/golang/go-quiz-api/routes"
)

func Run() {
	config.Load()
	listen(config.PORT)
}

// Configure listen port.
func listen(port uint32) {
	r := mux.NewRouter().StrictSlash(true)

	// setup routes
	routes.SetupRoutes(r)

	log.Println("Server running on http://localhost:4001...")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
