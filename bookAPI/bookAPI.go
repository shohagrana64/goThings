package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Desc   string `json:"desc"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}
func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}
func createNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, book := range Books {
		if book.Id == id {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var book Book
			json.Unmarshal(reqBody, &book)
			Books[index] = book
			json.NewEncoder(w).Encode(book)
		}
	}
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/books", createNewBook).Methods("POST")
	myRouter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/books/{id}", returnSingleBook)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Test title1", Author: "Shohag Rana", Genre: "Action", Desc: "Test Description1"},
		Book{Id: "2", Title: "Test title2", Author: "Sakib Al Amin", Genre: "Action", Desc: "Test Description2"},
	}
	handleRequests()
}
