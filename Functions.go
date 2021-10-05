package main

import "fmt"

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	giveMeSpace()
	parametersAndArguments("Furkan")
	giveMeSpace()
	multipleParameters("Furkan", 24)
	giveMeSpace()
	fmt.Println(add(5, 6))
	giveMeSpace()
	fmt.Println(add2(3, 7))
	giveMeSpace()
	total := add3(8, 9)
	fmt.Println(total)
	giveMeSpace()
	fmt.Println(multipleReturn(7, "Surprise"))
	giveMeSpace()
	a, b := multipleReturn(3, "Hi")
	fmt.Printf("Type: %T , Value: %v \n", a, a)
	fmt.Printf("Type: %T , Value: %v \n", b, b)
	giveMeSpace()
	testcount(1)
	giveMeSpace()
	fmt.Println(factorial_recursion(4))
	giveMeSpace()
}

func parametersAndArguments(fname string) {

	fmt.Println(fname)
}

func multipleParameters(fname string, age int) {

	fmt.Printf("%v : %v\n", fname, age)

}

func add(x int, y int) int {
	return x + y
}

func add2(x int, y int) (result int) {
	result = x + y
	return
}

func add3(x int, y int) (result int) {
	result = x + y
	return
}

func multipleReturn(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}
