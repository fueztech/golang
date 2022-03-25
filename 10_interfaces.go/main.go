package main

import "fmt"

// Book Store!
// Interfaces are cool they can extend any type.

type printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

type Mug struct {
	Name  string
	Price float32
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f", d.Name, d.Price)
}

func (b *Book) printInfo() {
	fmt.Printf("Book: %s Price: %.2f", b.Title, b.Price)
}

func (m Mug) printInfo() {
	fmt.Printf("Mug: %s Price %.2f\n", m.Name, m.Price)
}

/*func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I'm a string: %s\n", i)
	case int:
		fmt.Printf("I'm a int %d\n", i)
	case Book:
		fmt.Printf("I'm a book: %s\n", i.(Book).Title)
}*/

func main() {

	seeingTrees := &Book{
		Title: "Seeing Trees",
		Price: 9.75,
	}

	tea := Drink{
		Name:  "Hibiscus Latte",
		Price: 3.49,
	}

	teaMug := Mug{
		Name:  "Ember Mug",
		Price: 3.00,
	}

	seeingTrees.printInfo()
	tea.printInfo()
	teaMug.printInfo()
	fmt.Printf("This is the address for Book: %p\n", &seeingTrees)

	// array, bookinfo and teainfo are now of type printer
	info := []printer{seeingTrees, tea, teaMug}

	info[0].printInfo()
	info[1].printInfo()
	info[2].printInfo()

	/*empty("Marcus")
	empty(23)
	empty(seeingTrees)*/

}
