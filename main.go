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

func onlyGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware start")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("Middleware before next")
		next.ServeHTTP(w, r)
		fmt.Println("Middleware after next")
	})
}

func logPathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func logAppName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-App-Name", "go-http-playground")
		next.ServeHTTP(w, r)
	})
}

func JsonResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ping(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "pong"}
	JsonResponse(w, response)
}

func health(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: "ok"}
	JsonResponse(w, response)
}

func hello(w http.ResponseWriter, r *http.Request) {
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
	http.Handle("/ping", logAppName(onlyGetMiddleware(logPathMiddleware(http.HandlerFunc(ping)))))
	http.Handle("/health", onlyGetMiddleware(http.HandlerFunc(health)))
	http.Handle("/hello", onlyGetMiddleware(http.HandlerFunc(hello)))
	http.Handle("/users/", onlyGetMiddleware(http.HandlerFunc(usersHandler)))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
