package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		//
		println("Positive number")
	} else if num < 0 {
		println("Negative number")
	} else {
		// if not true
		println("You did not enter the number 5")
	}

	println("This will keep running!")
}
