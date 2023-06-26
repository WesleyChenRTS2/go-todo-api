package main

func main() {
    a := App{}
    a.Initialize("postgres", "mysecretpassword", "localhost", "testdb")
    a.Run(":8080")
}
