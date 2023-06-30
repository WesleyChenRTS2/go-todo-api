package main

// @title Echo Swagger TODO API
// @version 1.0
// @description This is a sample server Echo Swagger server for TODO API.

// @host localhost:8080
// @BasePath /
// @schemes http

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
      }


    a := App{}
    a.Initialize(os.Getenv("USER"), os.Getenv("PASSWORD"),os.Getenv("HOST"), os.Getenv("DBNAME"))
    a.Run(":8080")
}
