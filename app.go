package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type App struct {
    Router *echo.Echo
    DB     *sql.DB
}

func (a *App) Initialize(user, password, host, dbname string) {
    connectionString := a.DBConnectionString(user, password, host, dbname)
    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }
	

    a.Router = echo.New()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    a.Router.Logger.Fatal(a.Router.Start(addr))
}

func (a *App) initializeRoutes() {
    a.Router.GET("/todos", a.getTodosHandler)
    a.Router.POST("/todo", a.createTodoHandler)
    a.Router.GET("/todo/:id", a.getTodoHandler)
    a.Router.PUT("/todo/:id", a.updateTodoHandler)
    a.Router.DELETE("/todo/:id", a.deleteTodoHandler)
	a.Router.GET("/health", a.healthCheckHandler)
}

func (a *App) DBConnectionString (user, password, host, dbname string) string {
	
	return "postgres://" + user + ":" + password + "@" + host + "/" + dbname + "?sslmode=disable"

}

