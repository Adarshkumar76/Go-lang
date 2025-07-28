package main

// By value
// func change(num int) {
// 	num = 10
// 	fmt.Println("Inside change function:", num)

// }

// By reference
func change(num *int) {
	*num = 10
	println("Inside change function:", *num)
}

func main() {
	num := 5
	// fmt.Println("Before change function:", num)
	// change(num)
	// fmt.Println("After change function:", num) // Output: 5

	println(&num)
	change(&num)
	println("After change function:", num) // Output: 10
}
