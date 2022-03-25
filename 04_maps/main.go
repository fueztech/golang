package main

import "fmt"

func main() {

	ages := make(map[string]int)
	ages["Marcus"] = 31
	fmt.Printf("Marcus is %d years old\n", ages["Marcus"])

	// Maps can be incremented over
	ages["Latoya"] += 31 // Love of my life
	fmt.Printf("Latoya is %d years old\n", ages["Latoya"])

	// you cannot get the address of values though
	//addr := &ages["Aziah"] // cannot take addres of a map value

	// Below is a map literal vs a make map.
	gpas := map[string]float32{
		"Marcus": 3.4,
		"Latoya": 3.9,
	}
	fmt.Printf("Marcus GPA is %.2f\n", gpas["Marcus"])
	fmt.Printf("Latoya GPS is %.2f\n", gpas["Latoya"])

	// Zero type of an unintialized map is nil.
	var visited map[string]bool
	visited = make(map[string]bool)
	visited["A"] = true
	fmt.Printf("A has been visited? %v\n", visited["A"])

	fruits := []string{
		"bananas",
		"kiwis",
		"apples",
		"strawberries",
		"tomatoes",
		"bananas",
	}
	fmt.Printf("Duplicate Fruits? %s\n", findDuplicates(fruits))
}

func findDuplicates(words []string) string {
	dupMap := make(map[string]bool)

	for i := 0; i < len(words); i++ {

		if !dupMap[words[i]] {
			dupMap[words[i]] = true
		} else {
			return words[i]
		}
	}
	return ("")
}
