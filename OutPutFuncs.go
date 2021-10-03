package main

import (
	"fmt"
)

func main() {
	var g, j string = "deneme", "deneme2"

	fmt.Print(g, "\n", j)

	var a string = "Hello"
	var b int = 15

	fmt.Print("\n")

	fmt.Printf("i has value: %v and type: %T\n", a, a)
	fmt.Printf("j has value: %v and type: %T", b, b)

	// Formatting Verbs for Printf()

	var c = 15.5
	var txt = "Hello World!"

	fmt.Println("Formatting Verbs for Printf()")

	fmt.Printf("%v\n", c)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%v%%\n", c)
	fmt.Printf("%T\n", c)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var i = 15

	fmt.Println("Integer Formatting Verbs")
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	// String Formatting Verbs

	fmt.Println("String Formatting Verbs")
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// Boolean Formatting Verbs

	var x = true
	var y = false

	fmt.Println("Boolean Formatting Verbs")

	fmt.Printf("%t\n", x)
	fmt.Printf("%t\n", y)

	// Float Formatting Verbs

	fmt.Println("Float Formatting Verbs")

	var PI = 3.14
	fmt.Printf("%e\n", PI)
	fmt.Printf("%f\n", PI)
	fmt.Printf("%.2f\n", PI)
	fmt.Printf("%6.2f\n", PI)
	fmt.Printf("%g\n", PI)

}
