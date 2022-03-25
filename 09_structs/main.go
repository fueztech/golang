package main

import (
	"fmt"
)

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) getArea() int {
	return r.Length * r.Width
}

//Adding json allows us to encode json values. Note that name is lowercased to access json instead of the struct Name.
type Person struct {
	Name string `json:"name"`
	age  uint8
	Race string
}

// Structs embedded into each other declaring the Field and Adding the Defined Struct next to it.
type Employee struct {
	ID     uint64
	Person Person
	Salary uint
}

func main() {

	var rect1 Rectangle
	rect1.Length = 5
	rect1.Width = 4

	fmt.Printf("area of rect1: %d\n", rect1.getArea())
	rect2 := Rectangle{}
	rect2.Length = 6
	rect2.Width = 9

	/* The order of delcaring these matters 5 is Length, 4 will be Width.
	rect3 := Rectangle{5, 4}
	fmt.Printf("rect3 length: %v width: %v\n", rect3.Length, rect3.Width)

	//Best when a struct has alot of fields to choose between
	rect4 := Rectangle{
		Width:  5,
		Length: 4,
	}
	fmt.Printf("rect4 length: %v width: %v\n", rect4.Length, rect4.Width)

	//new(****) Returns a pointer to a zero value, if do not set the Fields they will be empty strings or ints etc...
	var rect5 = new(Rectangle)
	rect5.Length = 5

	var rect6 = &Rectangle{}
	rect6.Width = 3

	Marcus := Employee{
		ID:     621475,
		Salary: 40000,
		Person: Person{
			Name: "Marcus",
			age:  31,
		},
	}

	fmt.Printf("Marcus: %v\n", Marcus)

	aziah := Person{
		Name: "Aziah",
		age:  31,
	}

	data, err := json.Marshal(aziah)
	if err != nil {
		// error
		fmt.Printf("Error: %v\n", err.Error())
	}
	Print out encoded json data and strings
	fmt.Printf("JSON: %v\n", data)
	fmt.Printf("JSON String: %s\n", data)*/

}
