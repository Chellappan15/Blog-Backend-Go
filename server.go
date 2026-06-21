package main

import (
	"fmt"
	"net/http"
	"blog-backend/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/posts/getposts", posts.GetPosts)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
