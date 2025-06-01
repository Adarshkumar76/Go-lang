package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict

func main() {
	m2 := map[int]int{2: 4, 3: 5}
	fmt.Println(m2) //map[2: 4, 3: 5]
	m1 := make(map[string]int)
	m1["a"] = 10
	m1["b"] = 20

	fmt.Println(m1["a"]) //10
	i, isPresent := m1["c"]
	//OR _, isPresent := m1["c"]
	fmt.Println(i, isPresent) //0 false - because the kay of "c" is not present in map

	for key, val := range m1 {
		fmt.Println("Key", key, ": value", val) //Key a : value 10 '\n' Key b : value 20
	}

	fmt.Println(len(m1)) //2
	fmt.Println(m1)      //map[a:10 b:20]
	delete(m1, "a")
	fmt.Println(m1) //map[b:20]

	// equality check with maps package
	boole := maps.Equal(m1, m1)
	println(boole) //true

}
