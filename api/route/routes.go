package route

import (
	"net/http"

	"github.com/Shivansh-yadva13/todo_react_go/api/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/get_todos", handlers.GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/api/get_todo/{id}", handlers.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/api/add_todo", handlers.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/api/del/{id}", handlers.DeleteTodo).Methods(http.MethodDelete)
	router.HandleFunc("/api/complete_todo/{id}", handlers.CompleteTodo).Methods(http.MethodPut)

	return router
}
