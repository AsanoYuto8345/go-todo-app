package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TodoListHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get All Todos")
}

func GetTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "todo with ID: "+strconv.Itoa(id))
}

func CreateTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "create new todo!")
}

func PutTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "update todo with ID: "+strconv.Itoa(id))
}

func DeleteTodoByIdHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "delete todo with ID: "+strconv.Itoa(id))
}
