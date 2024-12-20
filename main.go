package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Type int

const (
	UrgentImportant Type = iota
	NotUrgentImportant
	UrgentNotimportant
	NotUrgentNotImportant
)

type Todo struct {
	Id          int
	Description string
	Status      bool
	Type        Type
}

func isValidTodoType(t Type) bool {
	return t >= UrgentImportant && t <= NotUrgentNotImportant
}

func main() {
	fmt.Println("Listening on port 8181")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("GET all todos")
		w.Write([]byte(message))
	})

	mux.HandleFunc("GET /todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		message := fmt.Sprintf("GET single todo with id: %s", id)
		w.Write([]byte(message))
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		decorder := json.NewDecoder(r.Body)

		if err := decorder.Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !isValidTodoType(todo.Type) {
			http.Error(w, "Invallid 'Type' field value - must be between 1 and 4", http.StatusBadRequest)
			return
		}

		message := fmt.Sprintf("POST new todo: %+v", todo)
		w.Write([]byte(message))
	})

	mux.HandleFunc("DELETE /todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		message := fmt.Sprintf("Delete single todo with id: %s", id)
		w.Write([]byte(message))
	})

	if err := http.ListenAndServe("localhost:8181", mux); err != nil {
		fmt.Println(err.Error())
	}
}
