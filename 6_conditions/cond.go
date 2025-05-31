package main

func main() {
	age := 18
	if age >= 18 {
		println("You are an adult.")
	} else {
		println("You are a minor.")
	}

	// Nested if-else
	if age < 13 {
		println("You are a child.")
	} else if age >= 13 && age < 18 {
		println("You are a teenager.")
	} else if age >= 18 && age < 60 {
		println("You are an adult.")
	} else {
		println("You are a senior citizen.")
	}

	if num := 10; num%2 == 0 {
		println("The number is even (using if with initialization).")
	} else {
		println("The number is odd (using if with initialization).")
	}

	// Short variable declaration
	if isAdult := age >= 18; isAdult {
		println("You are an adult (using short variable declaration).")
	} else {
		println("You are not an adult (using short variable declaration).")
	}

}
