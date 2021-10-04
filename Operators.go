package main

import (
	"fmt"
)

func giveMeSpace() {
	fmt.Printf("\n################################ \n\n")
}

func main() {

	plus()
	giveMeSpace()
}

func plus() {
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)
}
