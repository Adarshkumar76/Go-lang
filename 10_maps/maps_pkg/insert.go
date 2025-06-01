package main

import (
	"fmt"
	"maps"
	"slices"
)

// Insert adds the key-value pairs from seq to m. If a key in seq already exists in m, its value will be overwritten.
func main() {
	m1 := map[int]string{
		1000: "THOUSAND",
	}
	s1 := []string{"zero", "one", "two", "three"}
	maps.Insert(m1, slices.All(s1))
	fmt.Println("m1 is:", m1)
}

// m1 := map[int]string{1000: "THOUSAND"} creates a map m1 with a single key-value pair, where the key is an integer 1000 and the value is the string "THOUSAND".

// s1 := []string{"zero", "one", "two", "three"} creates a slice s1 with four string elements.

// maps.Insert(m1, slices.All(s1)) is the key part of the code. Here's what's happening:

// slices.All(s1) creates a new map that contains all the key-value pairs from the s1 slice. The keys are the indices of the slice elements, and the values are the corresponding slice elements.

// In this case, the resulting map would be {0: "zero", 1: "one", 2: "two", 3: "three"}.

// maps.Insert(m1, slices.All(s1)) takes the m1 map and inserts all the key-value pairs from the map returned by slices.All(s1) into it.
