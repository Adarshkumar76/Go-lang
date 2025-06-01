package main

import (
	"fmt"
	"slices"
)

// Sorted collects values from seq into a new slice, sorts the slice, and returns it.

func main() {
	seq := func(yield func(int) bool) {
		// 	seq is an anonymous function assigned to a variable.
		// It takes one parameter: yield, which is a function itself.
		// The yield function takes an int and returns a bool.

		flag := -1
		for i := 0; i < 10; i += 2 {
			flag = -flag
			if !yield(i * flag) {
				return
			}
			// The value i * flag is passed to the yield function.
			// If yield returns false, the loop stops early with return.
		}
	}

	s := slices.Sorted(seq)
	fmt.Println(s)
	fmt.Println(slices.IsSorted(s))
}
