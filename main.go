package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test Title", Description: "Test Description", Content: "Test content"},
	}

	fmt.Println("Endpoint Hit: All articles")
	json.NewEncoder(w).Encode(articles)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Indexpage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequest()
}
