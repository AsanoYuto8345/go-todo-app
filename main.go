package main

import (
	handlers "go-todo-app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos/list", handlers.TodoListHandler)
	http.HandleFunc("/todos/1", handlers.GetTodoByIdHandler)
	http.HandleFunc("/todos", handlers.CreateTodoHandler)
	http.HandleFunc("/todos/2", handlers.PutTodoByIdHandler)
	http.HandleFunc("/todos/3", handlers.DeleteTodoByIdHandler)
	log.Println("server start at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
