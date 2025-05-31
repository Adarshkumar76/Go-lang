package main

import "fmt"

func main() {
	a := 10
	var b int = 20
	var c int = a + b
	fmt.Printf("Sum of a and b is: %d\n", c)

	var d string = "Hello, World!"
	fmt.Printf("String value is: %s\n", d)
	//complex numbers
	var h complex128 = 1 + 2i
	fmt.Printf("Complex value is: %v\n", h)
}
