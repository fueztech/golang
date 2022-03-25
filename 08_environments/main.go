package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("HOME:", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Printf("the env var SHELL is not set!")
	} else {
		fmt.Println("SHELL:", shell)
	}

	err := os.Setenv("NAME", "Marcus")
	if err != nil {
		fmt.Printf("Could not set the env var NAME!")
	}

	fmt.Printf("NAME: %s\n", os.Getenv("NAME"))
}
