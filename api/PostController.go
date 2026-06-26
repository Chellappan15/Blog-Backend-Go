package posts

import (
	"blog-backend/database"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var mockDB = []Post{
	{ID: 1, Title: "Introduction to Go", Description: "Go is fast and simple."},
	{ID: 2, Title: "Understanding Struct Tags", Description: "Struct tags bridge Go and JSON."},
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Function called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockDB)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a new post")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newPost.Title == "" || newPost.Description == "" {
		http.Error(w, "Fill out all the details", http.StatusBadRequest)
	}
	fmt.Println("Got the request")

	postAdded, dbError := database.DB.Db.Exec(database.PostAddQuery(), newPost.Title, newPost.Description)
	fmt.Println(postAdded)

	if dbError != nil {
		http.Error(w, dbError.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Post added successfully")
}

func editPost(w http.ResponseWriter, r *http.Request) {
	var updatePost Post
	if err := json.NewDecoder(r.Body).Decode(&updatePost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
