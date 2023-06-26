package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create a new TODO
func (a *App) createTodoHandler(c echo.Context) error {
	var t Todo
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := t.createTodo(a.DB); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, t)
}

// Get a single TODO
func (a *App) getTodoHandler(c echo.Context) error {
	id, err := getTodoID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	t, err := getTodoByID(a.DB, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, t)
}

// Get all TODOs
func (a *App) getTodosHandler(c echo.Context) error {
	todos, err := getTodos(a.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todos)
}

// Update a TODO
func (a *App) updateTodoHandler(c echo.Context) error {
	id, err := getTodoID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	var t Todo
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	t.ID = id
	if err := t.updateTodo(a.DB); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, t)
}

// Delete a TODO
func (a *App) deleteTodoHandler(c echo.Context) error {
	id, err := getTodoID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	t := Todo{ID: id}
	if err := t.deleteTodo(a.DB); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted"})
}

// Helper function to extract todo ID from request context
func getTodoID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Helper function to get a todo by ID from the database
func getTodoByID(db *sql.DB, id int) (Todo, error) {
	t := Todo{ID: id}
	err := t.getTodo(db)
	return t, err
}

func (a *App) healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}