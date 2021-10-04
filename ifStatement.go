package main

import (
	"fmt"
)

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	giveMeSpace()
	condition1()
	giveMeSpace()
	condition2()
	giveMeSpace()
	condition3()
	giveMeSpace()
	condition4()
	giveMeSpace()
}

func condition1() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
}

func condition2() {
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

func condition3() {
	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}

func condition4() {
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}
