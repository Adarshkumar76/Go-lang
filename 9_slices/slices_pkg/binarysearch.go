package main

import (
	"fmt"
	"slices"
)

// BinarySearch searches for target in a sorted slice and returns the earliest position where target is found, or the position where target would appear in the sort order; it also returns a bool saying whether the target is really found in the slice. The slice must be sorted in increasing order.

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}
