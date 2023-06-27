package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateTodoHandler godoc
// @Summary Create a new TODO
// @Description Create a new TODO with the input payload
// @Tags todo
// @Accept json
// @Produce json
// @Param data body TodoPayloadBody true "Create TODO"
// @Success 201 {object} Todo
// @Router /todos [post]
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

// GetTodoHandler godoc
// @Summary Get a single TODO by ID
// @Description Get specific TODO details by ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Router /todos/{id} [get]
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

// GetTodosHandler godoc
// @Summary Get all TODOs
// @Description Get all TODOs
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func (a *App) getTodosHandler(c echo.Context) error {
	todos, err := getTodos(a.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todos)
}

// UpdateTodoHandler godoc
// @Summary Update a TODO
// @Description Update a TODO with the input payload
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body TodoPayloadBody true "Update TODO"
// @Success 200 {object} Todo
// @Router /todos/{id} [put]
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

// DeleteTodoHandler godoc
// @Summary Delete a TODO
// @Description Delete a TODO by ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Router /todos/{id} [delete]
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

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (a *App) healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}