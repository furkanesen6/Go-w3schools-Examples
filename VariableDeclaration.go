package main

import (
	"fmt"
)

var e int
var f int = 2
var g = 3

func main() {

	// Variable Declaration With Initial Value

	var student string = "okay"
	var student2 = "erdin"
	x := 10

	fmt.Println("Variable Declaration With Initial Value")
	fmt.Println(student)
	fmt.Println(student2)
	fmt.Println(x)

	// Variable Declaration Without Initial Value

	var a string
	var b int
	var c bool
	var d float32

	fmt.Println("Variable Declaration Without Initial Value")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Value Assignment After Declaration

	var student3 string
	student3 = "ömer"

	fmt.Println("Value Assignment After Declaration")
	fmt.Println(student3)

	// Difference Between var and :=

	e = 1

	fmt.Println("Difference Between var and :=")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Go Multiple Variable Declaration

	var h, j, k, l int = 1, 3, 5, 7

	fmt.Println("Go Multiple Variable Declaration")
	fmt.Println(h)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	var m, n = 6, "Hello"
	o, p := 7, "World!"

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)

	// Go Variable Declaration in a Block

	var (
		r int
		s int    = 1
		t string = "hello"
	)

	fmt.Println("Go Variable Declaration in a Block")
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

}
