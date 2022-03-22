package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// Headers handle cookies
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: steelblue'>I Love Aziah</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//w.Write([]byte("Content-Type",))
	book := Book{Title: "Fuez Art", Author: "Marcus Morris", Pages: 150}

	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/Hello", Hello) /*func(rw http.ResponseWriter, r *http.Request)*/
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
