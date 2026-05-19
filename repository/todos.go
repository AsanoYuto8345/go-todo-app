package repository

import (
	"go-todo-app/models"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetAllTodos(db *sqlx.DB) ([]models.Todo, error) {
	dbTodos := []struct {
		ID        int       `db:"id"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}{}

	query := "SELECT id, title, content, created_at, updated_at FROM todos"

	err := db.Select(&dbTodos, query)
	if err != nil {
		return nil, err
	}

	todos := make([]models.Todo, 0)
	for _, dbTodo := range dbTodos {
		todos = append(todos, models.Todo{
			ID:        dbTodo.ID,
			Title:     dbTodo.Title,
			Content:   dbTodo.Content,
			CreatedAt: dbTodo.CreatedAt,
			UpdatedAt: dbTodo.UpdatedAt,
		})
	}
	return todos, nil
}
