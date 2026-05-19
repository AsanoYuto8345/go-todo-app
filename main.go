package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbHost     = "127.0.0.1"
	dbPort     = "5432"
	dbName     = "postgres"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Successfuly connected to the database")
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	handlers "go-todo-app/handlers"
// 	models "go-todo-app/models"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	todo1 := models.Todo{
// 		ID:        1,
// 		Title:     "todo1 Title",
// 		Content:   "todo1 content",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	jsonData, err := json.Marshal(todo1)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))

// 	e := echo.New()

// 	e.GET("/todos", handlers.GetAllTodosHandler)
// 	e.GET("/todos/:id", handlers.GetTodoByIdHandler)
// 	e.POST("/todos", handlers.CreateTodoHandler)
// 	e.PUT("/todos/:id", handlers.PutTodoByIdHandler)
// 	e.DELETE("/todos/:id", handlers.DeleteTodoByIdHandler)

// 	e.Logger.Fatal(e.Start(":8080"))

// }
