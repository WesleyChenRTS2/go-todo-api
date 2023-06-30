package main

import (
	"database/sql"
	"log"

	_ "github.com/WesleyChenRTS2/go-todo-api/docs/echoapi"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	
    // Echo Instance
    a.Router = echo.New()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    a.Router.Logger.Fatal(a.Router.Start(addr))
}


func (a *App) initializeRoutes() {
    // Routes
    a.Router.GET("/todos", a.getTodosHandler)
    a.Router.POST("/todo", a.createTodoHandler)
    a.Router.GET("/todo/:id", a.getTodoHandler)
    a.Router.PUT("/todo/:id", a.updateTodoHandler)
    a.Router.DELETE("/todo/:id", a.deleteTodoHandler)
	a.Router.GET("/health", a.healthCheckHandler)
    a.Router.GET("/swagger/*", echoSwagger.WrapHandler)

   
        // todo := models.Todos{ID: 1, Title: "My first todo"}
        // fmt.Println(todo.Title)
   
    
}

func (a *App) DBConnectionString (user, password, host, dbname string) string {
	
	return "postgres://" + user + ":" + password + "@" + host + "/" + dbname + "?sslmode=disable"

}

