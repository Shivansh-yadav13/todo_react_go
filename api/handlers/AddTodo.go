package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Shivansh-yadva13/todo_react_go/api/mocks"
	"github.com/Shivansh-yadva13/todo_react_go/api/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var todo models.Todos
	json.Unmarshal(body, &todo)
	todo.Id = len(mocks.TodoList) + 1

	mocks.TodoList = append(mocks.TodoList, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Todo Added!")
}
