package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Status: "ok"})
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
