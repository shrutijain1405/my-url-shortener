package main

import (
    "fmt"
    "log"
    "my-url-shortener/connection"
    "my-url-shortener/service"
    "net/http"
)

func main() {
    fmt.Println("hello")
    connection.MakeDbConnection()
    service.InitServer()
    err := http.ListenAndServe(":8080", service.Router)
    if err != nil {
        log.Fatal(err)
    }
}
