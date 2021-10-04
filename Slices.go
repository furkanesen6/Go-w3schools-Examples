package main

import (
	"fmt"
)

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	createSlices1()
	giveMeSpace()
	createSlices2()
	giveMeSpace()
	createSlices3()
	giveMeSpace()
	modify1()
	giveMeSpace()
	modify2()
	giveMeSpace()
	modify3()
	giveMeSpace()
	modify4()
	giveMeSpace()
	modify5()
	giveMeSpace()
	modify6()
	giveMeSpace()

}

func createSlices1() {
	mySlice := []int{}

	fmt.Println("Create a Slice With []datatype{values}")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice2 := []string{"Go", "Slices", "Are", "Powerful"}

	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
}

func createSlices2() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	mySlice := arr1[2:5]

	fmt.Println("Create a Slice From an Array")
	fmt.Printf("myslice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))
}

func createSlices3() {

	fmt.Println("Create a Slice With The make() Function")

	mySlice := make([]int, 5, 10)
	fmt.Printf("myslice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))

	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

func modify1() {
	prices := []int{10, 20, 30}

	fmt.Println("Access Elements of a Slice")
	fmt.Println(prices[0])
}

func modify2() {
	prices := []int{10, 20, 30}
	prices[0] = 50

	fmt.Println("Change Elements of a Slice")
	fmt.Println(prices[0])
}

func modify3() {
	prices := []int{10, 11, 12, 13, 14, 15}

	fmt.Println("Append Elements To a Slice")
	fmt.Printf("prices = %v\n", prices)
	fmt.Printf("length = %d\n", len(prices))
	fmt.Printf("capacity = %d\n", cap(prices))

	prices = append(prices, 20, 21, 22)

	fmt.Printf("prices = %v\n", prices)
	fmt.Printf("length = %d\n", len(prices))
	fmt.Printf("capacity = %d\n", cap(prices))

}

func modify4() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Println("Append One Slice To Another Slice")
	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}

func modify5() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Println("Change The Length of a Slice")
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

func modify6() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println("Memory Efficiency")
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))

	fmt.Printf("numbersCopy before copy = %v\n", numbersCopy)
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy after copy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
