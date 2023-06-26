package main

import (
	"database/sql"
)

type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func getTodos(db *sql.DB) ([]Todo, error) {
    rows, err := db.Query("SELECT * FROM todos")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    todos := []Todo{}

    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
            return nil, err
        }
        todos = append(todos, t)
    }

    return todos, nil
}

func (t *Todo) createTodo(db *sql.DB) error {
    return db.QueryRow(
        "INSERT INTO todos(title, completed) VALUES($1, $2) RETURNING id",
        t.Title, t.Completed).Scan(&t.ID)
}

func (t *Todo) getTodo(db *sql.DB) error {
    return db.QueryRow("SELECT title, completed FROM todos WHERE id=$1",
        t.ID).Scan(&t.Title, &t.Completed)
}

func (t *Todo) updateTodo(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE todos SET title=$1, completed=$2 WHERE id=$3",
            t.Title, t.Completed, t.ID)

    return err
}

func (t *Todo) deleteTodo(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM todos WHERE id=$1", t.ID)

    return err
}