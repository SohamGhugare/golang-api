package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article structure
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

// Handling the articles endpoint
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello world",
		},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// Handling the root endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

// Registering all handlers
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
