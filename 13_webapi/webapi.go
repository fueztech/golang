package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API! Check things out while you are here!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Marcus Morris"

	fmt.Fprintf(response, "A little bit about Marcus Morris...")
	fmt.Println("Endpoint Hit: ", who)
}

func request25() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8040", nil))

}

func main() {
	request25()
}
