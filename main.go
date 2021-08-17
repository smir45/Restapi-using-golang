package main //compulsary object

//importing packages in golang

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

//creating functions

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test Title", Description: "Test Description", Content: "Test content"},
	}

	fmt.Println("Endpoint Hit: All articles")
	json.NewEncoder(w).Encode(articles)
}

//creating post functions
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	beego.RegisterController("/", &controllers.IndexController{})
	
}

//creating add functions
func testAddArticles(w http.ResponseWriter, r *http.Request) {
	beego.RegisterController("/new", &controllers.NewController{})
}

//creating delete functions
func testDeleteArticles(w http.ResponseWriter, r *http.Request) {
	beego.RegisterController("/delete/: id([0-9]+)", &controllers.DeleteController{})
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	beego.RegisterController("/", &controllers.IndexController{})
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", indexPage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/articles", testAddArticles).Methods("PUT")
	myRouter.HandleFunc("/articles", testDeleteArticles).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequest()
}
