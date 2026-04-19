package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

var Users = map[int]*User{
	1: {ID: 1, Name: "John"},
	2: {ID: 2, Name: "Bob"},
	3: {ID: 3, Name: "Elena"},
}

func FindUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, ok := Users[id]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
