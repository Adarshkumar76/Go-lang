package main

import "fmt"

// For -> only construct in go for loops there is no while or do-while loops
func main() {
	// for loop with range || classic for loop
	for j := 1; j <= 10; j++ {
		println("Value of j is:", j)
	}

	//while loop
	i := 1
	for i <= 10 {
		println("Value of i is:", i)
		i++ // increment i by 1
	}

	// for loop with range
	for q := range 3 {
		fmt.Printf("%d\n", q)
	}

	//	// for loop with range over a string
	//	s := "Hello"
	//	for index, char := range s {
	//		println("Index:", index, "Character:", string(char))
	//	}

	// nested for loop
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			println("x:", x, "y:", y)
		}
	}

	//	// break statement
	//	for k := 1; k <= 10; k++ {
	//		if k == 5 {
	//			println("Breaking the loop at k =", k)
	//			break // exit the loop when k is 5
	//		}
	//		println("k is:", k)
	//	}
	//	// continue statement
	//	for l := 1; l <= 10; l++ {
	//		if l%2 == 0 {
	//			println("Skipping even number:", l)
	//			continue // skip the rest of the loop for even numbers
	//		}
	//		println("Odd number:", l)
	//	}

	//	// infinite loop
	//	for {
	//		println("This is an infinite loop. Press Ctrl+C to stop.")
	//		// Uncomment the next line to break the infinite loop after 5 iterations
	//		// if i > 5 { break }
	//	}

	//	// labeled for loop
	//
outerLoop:

	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			if m == 2 && n == 2 {
				println("Breaking out of outer loop at m =", m, "n =", n)
				break outerLoop // break out of the outer loop
			}
			println("m:", m, "n:", n)
		}
	}
	//	// for loop with a condition
	//	for o := 1; o <= 10; o++ {
	//		if o%2 == 0 {
	//			println("Even number:", o)
	//		} else {
	//			println("Odd number:", o)
	//		}
	//	}
	//	// for loop with a post statement
	//	for p := 1; p <= 10; p++ {
	//		println("Value of p is:", p)
	//		if p == 5 {
	//			println("Reached halfway point at p =", p)
	//		}
	//	}
	//	// for loop with a pre statement
	//	for q := 10; q >= 1; q-- {
	//		println("Value of q is:", q)
	//		if q == 5 {
	//			println("Reached halfway point at q =", q)
	//		}
	//	}

}
