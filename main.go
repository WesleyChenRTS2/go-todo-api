package main

// @title Echo Swagger TODO API
// @version 1.0
// @description This is a sample server Echo Swagger server for TODO API.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
    a := App{}
    a.Initialize("postgres", "mysecretpassword", "localhost", "testdb")
    a.Run(":8080")
}
