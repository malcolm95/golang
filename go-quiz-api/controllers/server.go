package controllers

import "net/http"

// generic method to confirm that server is running
func ServerRunning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running..."))
}
