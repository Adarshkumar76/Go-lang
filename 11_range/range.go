package main

func main() {
	// This is a simple program that demonstrates the use of the range keyword in Go.
	// The range keyword is used to iterate over elements in various data structures like arrays, slices, maps, and strings.

	// Example 1: Iterating over a slice
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		println("Index:", index, "Value:", value)
	}

	// Example 2: Iterating over a map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		println("Key:", key, "Value:", value)
	}

	// Example 3: Iterating over a string
	str := "hello"
	for index, char := range str {
		println("Index:", index, "Character:", string(char))
	}

	// Example 4: Using range with only the value
	for _, value := range slice {
		println("Value:", value)
	}
	// Example 5: Using range with only the index
	for index := range slice {
		println("Index:", index)
	}

	// Example 6: Using range with a multi-dimensional slice
	multiSlice := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for i, innerSlice := range multiSlice {
		for j, value := range innerSlice {
			println("Outer Index:", i, "Inner Index:", j, "Value:", value)
		}
	}
}
