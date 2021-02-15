package main

import (
	"encoding/json"
	"fmt"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
)

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Desc   string `json:"desc"`
}

var Books []Book
var (
	username = "abc"
	password = "123"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Fprintf(w, validToken)
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
func basicAuthentication(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		u, p, ok := r.BasicAuth()
		if !ok {
			fmt.Println("Error parsing basic auth")
			w.WriteHeader(401)
			fmt.Fprintf(w, "Error parsing basic auth")
			return
		}
		if u != username {
			fmt.Printf("Username provided is incorrect: %s\n", u)
			w.WriteHeader(401)
			fmt.Fprintf(w, "Username provided is incorrect")
			return
		}
		if p != password {
			fmt.Printf("Password provided is incorrect: %s\n", p)
			w.WriteHeader(401)
			fmt.Fprintf(w, "Password provided is incorrect")
			return
		}
		fmt.Printf("Username: %s\n", u)
		fmt.Printf("Password: %s\n", p)
		w.WriteHeader(200)
		endpoint(w, r)
		return
	})
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.Handle("/books", basicAuthentication(createNewBook)).Methods("POST")
	myRouter.Handle("/books", basicAuthentication(returnAllBooks))
	myRouter.Handle("/books/{id}", basicAuthentication(updateBook)).Methods("PUT")

	myRouter.Handle("/books/{id}", basicAuthentication(deleteBook)).Methods("DELETE")
	myRouter.Handle("/books/{id}", basicAuthentication(returnSingleBook))
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Test title1", Author: "Shohag Rana", Genre: "Action", Desc: "Test Description1"},
		Book{Id: "2", Title: "Test title2", Author: "Sakib Al Amin", Genre: "Action", Desc: "Test Description2"},
	}
	handleRequests()
}
