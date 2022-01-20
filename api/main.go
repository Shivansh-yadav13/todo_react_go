package main

import (
	"log"
	"net/http"

	"github.com/Shivansh-yadva13/todo_react_go/api/route"
	"github.com/gorilla/handlers"
)

func main() {
	log.Println("Api running at Port 3001 http://localhost:3000")
	http.ListenAndServe(":3001", handlers.CORS()(route.Router()))
}
