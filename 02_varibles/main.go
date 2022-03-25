package main

import "fmt"

func main() {
	// varibles can be declared in different ways a fews of those ways are displayed below.

	var favLinuxDistro string
	favLinuxDistro = "PopOS"

	println("My favorite linux distro is", favLinuxDistro)

	// Multiple Varibles can be declared at ounce as well with different datatypes.
	//%v is simply a placeholder that references a varible. This comes in handy.

	var firstName, lastName string = "Marcus", "Morris"

	/*var (
		age      = 31
		myHeight = "5'7"
		myWeight = 188
	)*/

	fmt.Printf("My name is %v %v \n", firstName, lastName)
	//fmt.Printf("I am %v years old, I stand %v and weigh %v pounds.\n", age, myHeight, myWeight)

	//Multiple Varibles can be declared in shorthand as well.
	friendName, friendLastName := "Latoya", "Griggs"
	fmt.Printf("I love my friend %v %v, she makes me smile!\n", friendName, friendLastName)

	var (
		height = 8
		width  = 11
		length = 13
	)

	fmt.Printf("The ideal room size is %v x %v x %v\n", length, width, height)

	const PI = 3.14159

	fmt.Println("Golang is cool it has constant varibles too, that cannot be modfied", PI)
}
