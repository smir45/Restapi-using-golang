package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint is working")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Indexpage Endpoint Hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", indexPage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequest()
}
