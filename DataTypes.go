package main

import (
	"fmt"
)

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// Boolean Data Type

	fmt.Println("Boolean Data Type")

	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	// Integer Data Type

	fmt.Println("Integer Data Type")

	var o int = 45000
	var p int = -45322

	fmt.Printf("Type: %T , Value: %v \n", o, o)
	fmt.Printf("Type: %T , Value: %v \n", p, p)

	// Float Data Types

	fmt.Println("Float Data Types")

	var r float32 = 123.78
	var t float64 = 1.7e+308

	fmt.Printf("Type: %T, value: %v\n", r, r)
	fmt.Printf("Type: %T, value: %v", t, t)

	// String Data Type

	fmt.Println("String Data Type")

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

}
