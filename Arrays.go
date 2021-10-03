package main

import (
	"fmt"
)

func main() {

	basics()
	giveMeSpace()
	basicsUndefinedLength()
	giveMeSpace()
	stringArray()
	giveMeSpace()
	elementOfArray()
	giveMeSpace()
	changeElement()
	giveMeSpace()
	arrayInitialization()
	giveMeSpace()
	initializeOnlySpecificElements()
	giveMeSpace()
	lengthOfArray()
	giveMeSpace()

}

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func basics() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Printf("arr1, Type: %T , value: %v \n", arr1, arr1)
	fmt.Printf("arr2, Type: %T , value: %v \n", arr2, arr2)
}

func basicsUndefinedLength() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Printf("arr1, Type: %T , value: %v \n", arr1, arr1)
	fmt.Printf("arr2, Type: %T , value: %v \n", arr2, arr2)
}

func stringArray() {
	var cars = [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Printf("cars, Type: %T , value: %v \n", cars, cars)
}

func elementOfArray() {
	prices := [...]int{10, 20, 30}

	fmt.Printf("prices, Tyep: %T , value: %v \n", prices, prices)
	fmt.Printf("element[1]: %v \n", prices[1])

}

func changeElement() {
	prices := [...]int{12, 13, 14, 15}
	prices[2] = 50
	fmt.Printf("element[2] changed from 13 to 50 : %v \n", prices[2])
}

func arrayInitialization() {
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func initializeOnlySpecificElements() {
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)

}

func lengthOfArray() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
