package controllers

import "net/http"

// method to confirm that server is up and running
func ServerRunning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running..."))
}
