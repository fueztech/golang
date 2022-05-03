package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a web page! Cool huh? %s", time.Now())
}

func main() {
	http.HandleFunc("/", helloworld)
	fmt.Println("Now Serving localhost:8080")
	http.ListenAndServe(":8080", nil)
}
