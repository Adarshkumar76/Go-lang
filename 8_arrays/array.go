package main

func main() {
	// Declare and initialize an array of integers
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Print the array
	for i := 0; i < len(numbers); i++ {
		println(numbers[i])
	}

	// Modify an element in the array
	numbers[2] = 10

	// Print the modified array
	for i := 0; i < len(numbers); i++ {
		println(numbers[i])
	}
	// Declare and initialize an array of strings
	var fruits [3]string = [3]string{"apple", "banana", "cherry"}
	// Print the array of strings
	for i := 0; i < len(fruits); i++ {
		println(fruits[i])
	}

}
