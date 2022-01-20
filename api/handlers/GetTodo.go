package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Shivansh-yadva13/todo_react_go/api/mocks"
	"github.com/gorilla/mux"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, todo := range mocks.TodoList {
		if todo.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
		}
	}
}
