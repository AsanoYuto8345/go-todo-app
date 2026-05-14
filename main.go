package main

import (
	"encoding/json"
	"fmt"
	handlers "go-todo-app/handlers"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func main() {
	todo1 := Todo{
		ID:        1,
		Title:     "todo1 Title",
		Content:   "todo1 content",
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	jsonData, err := json.Marshal(todo1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	e := echo.New()

	e.GET("/todos/list", handlers.TodoListHandler)
	e.GET("/todos/:id", handlers.GetTodoByIdHandler)
	e.POST("/todos", handlers.CreateTodoHandler)
	e.PUT("/todos/:id", handlers.PutTodoByIdHandler)
	e.DELETE("/todos/:id", handlers.DeleteTodoByIdHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
