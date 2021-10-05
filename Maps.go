package main

import "fmt"

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	giveMeSpace()
	createMapsVar()
	giveMeSpace()
	createMapsMake()
	giveMeSpace()
	manipulatingMapElements()
	giveMeSpace()
	checkingForSpecificElement()
	giveMeSpace()
	referenceMaps()
	giveMeSpace()
	iteratingMaps()
	giveMeSpace()
}

func createMapsVar() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a --> Type: %T\tValue %v\n", a, a)
	fmt.Printf("b --> Type: %T\tValue %v\n", b, b)
}

func createMapsMake() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

func manipulatingMapElements() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf(a["brand"])

	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element

	fmt.Println(a)

	delete(a, "year") // deleting "year" element

	fmt.Println(a)
}

func checkingForSpecificElement() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value
	val4, _ := a["model"]   // Only checking for existing value and not its key
	val5 := a["model"]      // Only checking for existing value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
	fmt.Println(val4)
	fmt.Println(val5)
}

func referenceMaps() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Printf("a --> Type: %T\tValue %v\n", a, a)
	fmt.Printf("b --> Type: %T\tValue %v\n", b, b)

	b["Year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Printf("a --> Type: %T\tValue %v\n", a, a)
	fmt.Printf("b --> Type: %T\tValue %v\n", b, b)
}

func iteratingMaps() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}
