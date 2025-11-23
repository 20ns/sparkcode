package main
// Includes error handling and returning 400 status codes when JSON is missing.


import (
	"encoding/json"
	"log"
	"net/http"
)
// Match OpenAPI spec, added a data structure

type ToDo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos []ToDo

func main() {
	http.HandleFunc("/", ToDoListHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// Handle Cors, in production will specify the origins.

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(todos)
		return
	}

	if r.Method == "POST" {
		var todo ToDo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil || todo.Title == "" || todo.Description == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input"})
			return
		}

		todos = append(todos, todo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
