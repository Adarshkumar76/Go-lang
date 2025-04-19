package main

import "fmt"

func main() {
	fmt.Println(myFunc(6, 4))
}

func myFunc(a, b int) int {
	return a + b
}

// OR
// func myFunc(a int, b int) (result int) {
//	result a+b
// 	return result
// }
