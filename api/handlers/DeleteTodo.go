package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Shivansh-yadva13/todo_react_go/api/mocks"
	"github.com/gorilla/mux"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatalf(err.Error())
	}

	for idx, todo := range mocks.TodoList {
		if todo.Id == id {
			mocks.TodoList = append(mocks.TodoList[:idx], mocks.TodoList[idx+1:]...) // mocks.TodoList[:idx] idx -> is not included
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Todo Deleted!")
			break
		}
	}
}
