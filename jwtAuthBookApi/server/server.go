package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
		},
	})
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

// set environment variable first. This can be done by:
// export MY_JWT_TOKEN="icecreamKhabo"
// then use the following command. This will ensure security of the signing key.
// var mySigningKey = os.Get("MY_JWT_TOKEN")
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
	fmt.Fprintf(w, "\nWelcome to the HomePage!\n The secret is secret.\n")

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
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		processedToken := r.Header["Authorization"][0]
		processedTokens := strings.Split(processedToken, " ")
		//fmt.Fprintf(w, processedToken)
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(processedTokens[1], func(token *jwt.Token) (interface{}, error) {
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
	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, basicAuthentication(homePage))).Methods("GET")
	myRouter.Handle("/books", promhttp.InstrumentHandlerCounter(httpRequestsTotal, isAuthorized(createNewBook))).Methods("POST")
	myRouter.Handle("/books", promhttp.InstrumentHandlerCounter(httpRequestsTotal, isAuthorized(returnAllBooks))).Methods("GET")
	myRouter.Handle("/books/{id}", promhttp.InstrumentHandlerCounter(httpRequestsTotal, isAuthorized(updateBook))).Methods("PUT")
	myRouter.Handle("/books/{id}", promhttp.InstrumentHandlerCounter(httpRequestsTotal, isAuthorized(deleteBook))).Methods("DELETE")
	myRouter.Handle("/books/{id}", promhttp.InstrumentHandlerCounter(httpRequestsTotal, isAuthorized(returnSingleBook))).Methods("GET")
	myRouter.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{})).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Test title1", Author: "Shohag Rana", Genre: "Mystery", Desc: "Test Description1"},
		Book{Id: "2", Title: "Test title2", Author: "Sakib Al Amin", Genre: "Action", Desc: "Test Description2"},
	}
	handleRequests()
}
