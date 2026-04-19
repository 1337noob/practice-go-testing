package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Tasks = map[int]Task{
	1: {ID: 1, Name: "Task 1"},
	2: {ID: 2, Name: "Task 2"},
	3: {ID: 3, Name: "Task 3"},
	4: {ID: 4, Name: "Task 4"},
	5: {ID: 5, Name: "Task 5"},
}

func findTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	task, ok := Tasks[id]
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func NewServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks/{id}", findTaskByIDHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
