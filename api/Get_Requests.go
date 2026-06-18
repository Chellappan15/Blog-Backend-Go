package posts

import (
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

func GetPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	return mockDB, nil
}
