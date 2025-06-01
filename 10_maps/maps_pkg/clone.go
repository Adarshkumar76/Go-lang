//Clone returns a copy of m. This is a shallow clone: the new keys and values are set using ordinary assignment.

package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"key": 1,
	}
	m2 := maps.Clone(m1)
	m2["key"] = 100
	fmt.Println(m1["key"]) //1
	fmt.Println(m2["key"]) //100

	m3 := map[string][]int{
		"key": {1, 2, 3},
	}
	m4 := maps.Clone(m3)

	fmt.Println(m4["key"][0]) //1

	m4["key"][0] = 100
	//maps.Clone() function creates a new map,
	fmt.Println(m3["key"][0]) //100
	fmt.Println(m4["key"][0]) //100

	// example of create a deep copy of the map
	m5 := map[string][]int{"key": {1, 2, 3}}
	m6 := make(map[string][]int)
	for k, v := range m5 {
		m6[k] = make([]int, len(v))
		copy(m6[k], v)
	}
	fmt.Println(m5)
	fmt.Println(m6)
}
