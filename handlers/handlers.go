package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TodoListHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get All Todos")
}

func GetTodoByIdHandler(c echo.Context) error {
	return c.String(http.StatusOK, "do hogehoge")
}

func CreateTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "create new todo!")
}

func PutTodoByIdHandler(c echo.Context) error {
	return c.String(http.StatusOK, "update todo with ID:2")
}

func DeleteTodoByIdHandler(c echo.Context) error {
	return c.String(http.StatusOK, "delete todo with ID:3")
}
