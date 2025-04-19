package main

import "fmt"

func main() {
	variable()
}

func variable() {
	var a int = 10
	var b = 20
	c := 30
	const pi = 3.14

	var f32 float32 = 54.2
	var f64 float64 = 547.3

	fmt.Println(a, b, c, f32, f64)
}
