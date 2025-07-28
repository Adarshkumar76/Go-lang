package main

func main() {
	// This is a simple program that demonstrates the use of variadic functions in Go.
	// Variadic functions can take a variable number of arguments.

	// Example 1: A variadic function that sums integers
	result := sum(1, 2, 3, 4, 5)
	println("Sum:", result)

	// Example 2: A variadic function with no arguments
	emptyResult := sum()
	println("Sum of no arguments:", emptyResult)

}

func sum(nums ...int) int {
	total := 0
	for a, num := range nums {
		total += num
		println("Index:", a, "Value:", num)
		println("Adding:", num, "Current Total:", total)
	}
	return total
}
