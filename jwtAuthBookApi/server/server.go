package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

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

//var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("icecreamKhabo")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Shohag Rana"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	fmt.Fprintf(w, validToken)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "Welcome to the HomePage!\n The secret is secret.\n The Token:")

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
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", isAuthorized(homePage))
	myRouter.Handle("/books", isAuthorized(createNewBook)).Methods("POST")
	myRouter.Handle("/books", isAuthorized(returnAllBooks))
	myRouter.Handle("/books/{id}", isAuthorized(updateBook)).Methods("PUT")

	myRouter.Handle("/books/{id}", isAuthorized(deleteBook)).Methods("DELETE")
	myRouter.Handle("/books/{id}", isAuthorized(returnSingleBook))
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Test title1", Author: "Shohag Rana", Genre: "Action", Desc: "Test Description1"},
		Book{Id: "2", Title: "Test title2", Author: "Sakib Al Amin", Genre: "Action", Desc: "Test Description2"},
	}
	handleRequests()
}
