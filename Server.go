package main

import (
	posts "blog-backend/api"
	"blog-backend/database"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/posts/getposts", posts.GetPosts)

	fmt.Println("Server started at :8080")
	db, error := database.Connect(os.Getenv("DATABASE_CONNECTION_STRING"))
	if error != nil {
		fmt.Println(error)
		return
	}
	defer db.Close()
	fmt.Println("Connected to database")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
