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
	mux.HandleFunc("POST /posts/addpost", posts.AddPost)

	fmt.Println("Server started at :5000")
	db, error := database.Connect(os.Getenv("DATABASE_CONNECTION_STRING"))
	if error != nil {
		fmt.Println(error)
		return
	}
	// env := &ENV{DB: db}
	defer db.Close()
	fmt.Println("Connected to database")
	// mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
	// 	stats := env.DB.Stats()
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(stats)
	// })

	http.ListenAndServe(":5000", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
