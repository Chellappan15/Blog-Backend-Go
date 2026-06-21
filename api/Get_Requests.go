package posts

import (
	"encoding/json"
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
	if mockDB == nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // 2. Set the content type to JSON
    w.Header().Set("Content-Type", "application/json")
    
    // 3. Encode and write the data directly to the response writer
    json.NewEncoder(w).Encode(mockDB)
}
