package handlers

import (
	"fmt"
	models "go-todo-app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTodosHandler(c echo.Context) error {
	var allTodos []models.Todo = []models.Todo{models.Todo1, models.Todo2}

	return c.JSON(http.StatusOK, allTodos)
}

func GetTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	log.Println(id)

	return c.JSON(http.StatusOK, models.Todo1)
}

func CreateTodoHandler(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid JSON format"})
	}

	return c.JSON(http.StatusOK, todo)
}

func PutTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid JSON format"})
	}

	todo.ID = id

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Todo with ID %d has been deleted", id),
	})
}
