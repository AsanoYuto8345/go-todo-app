package main

import (
	handlers "go-todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/todos/list", handlers.TodoListHandler)
	e.GET("/todos/1", handlers.GetTodoByIdHandler)
	e.POST("/todos", handlers.CreateTodoHandler)
	e.PUT("/todos/2", handlers.PutTodoByIdHandler)
	e.DELETE("/todos/3", handlers.DeleteTodoByIdHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
