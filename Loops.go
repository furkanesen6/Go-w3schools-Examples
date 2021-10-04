package main

import (
	"fmt"
)

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	giveMeSpace()
	forLoop1()
	giveMeSpace()
	forLoop2()
	giveMeSpace()
	forLoop3()
	giveMeSpace()
	nestedLoops()
	giveMeSpace()
	rangeKeyword()
	giveMeSpace()
}

func forLoop1() {
	fmt.Println("for Loop Examples")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forLoop2() {
	fmt.Println("The continue Statement")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func forLoop3() {
	fmt.Println("The break statement")

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

func nestedLoops() {
	fmt.Println("Nested loops")

	adj := [2][3]string{{"big", "tasty", "salty"}, {"big3", "tasty2", "salty1"}}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(adj[0]); j++ {
			fmt.Println(adj[i][j])
		}
	}
}

func rangeKeyword() {
	fmt.Println("The Range Keyword")

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
