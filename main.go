package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func JsonResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func onlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func ping(w http.ResponseWriter, r *http.Request) {
	if !onlyGet(w, r) {
		return
	}
	response := Response{Message: "pong"}
	JsonResponse(w, response)
}

func health(w http.ResponseWriter, r *http.Request) {
	if !onlyGet(w, r) {
		return
	}
	response := Response{Status: "ok"}
	JsonResponse(w, response)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if !onlyGet(w, r) {
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := "Hello " + name + "!"
	response := Response{Message: message}
	JsonResponse(w, response)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if !onlyGet(w, r) {
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := parts[2]
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message := "User " + name
	response := Response{Message: message}
	JsonResponse(w, response)
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/health", health)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/users/", usersHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
