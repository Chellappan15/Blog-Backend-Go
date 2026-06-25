package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var mockDB = []Post{
	{ID: 1, Title: "Introduction to Go", Content: "Go is fast and simple."},
	{ID: 2, Title: "Understanding Struct Tags", Content: "Struct tags bridge Go and JSON."},
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

	mockDB = append(mockDB, newPost)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Post added successfully ", mockDB)
	json.NewEncoder(w).Encode("Post added successfully")
}

func editPost(w http.ResponseWriter, r *http.Request) {
	var updatePost Post
	if err := json.NewDecoder(r.Body).Decode(&updatePost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
