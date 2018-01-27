package main

import (
    "fmt"
    "net/http"
)

func HelloHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world.")
}

func main() {
    http.HandleFunc("/", HelloHandlers) // set router
    http.ListenAndServe(":8080", nil) // set listen port
    fmt.Println("server run at http://localhost:8080");
}
