package main

import (
	handlers "go-todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/todos/list", handlers.TodoListHandler)
	e.GET("/todos/:id", handlers.GetTodoByIdHandler)
	e.POST("/todos", handlers.CreateTodoHandler)
	e.PUT("/todos/:id", handlers.PutTodoByIdHandler)
	e.DELETE("/todos/:id", handlers.DeleteTodoByIdHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
