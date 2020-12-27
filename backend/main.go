package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Chat app v0.0.1 started")

	setupRoutes()
	http.ListenAndServe(":8080", nil)

	fmt.Println("Chat app v0.0.1 finished")
}

func setupRoutes() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf("Simple server is open")
		})
}
