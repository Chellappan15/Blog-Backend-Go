package main

import (
	"fmt"
	"net/http"

	"github.com/Chellappan15/Blog-Backend-Go/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/posts/getposts", api.GetPosts)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
