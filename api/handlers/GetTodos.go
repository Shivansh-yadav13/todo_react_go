package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Shivansh-yadva13/todo_react_go/api/mocks"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.TodoList)
}
