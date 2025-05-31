package main

import "time"

func main() {
	// Switch statement
	// day := "Monday"
	switch time.Now().Weekday().String() {
	case "Monday":
		println("It's the start of the week.")
	case "Tuesday":
		println("It's the second day of the week.")
	case "Wednesday":
		println("It's the middle of the week.")
	case "Thursday":
		println("It's almost the weekend.")
	case "Friday":
		println("It's the last working day of the week.")
	case "Saturday", "Sunday":
		println("It's the weekend!")
	default:
		println("Invalid day.")
	}

	// Type switch
	var value interface{} = "42" // interface{} can hold any type of value (it means the variable can hold any type of value)
	switch v := value.(type) {
	case int:
		println("Value is an integer:", v)
	case string:
		println("Value is a string:", v)
	case bool:
		println("Value is a boolean:", v)
	default:
		println("Unknown type")
	}

	// Fallthrough example
	switch num := 2; num {
	case 1:
		println("Number is 1")
	case 2:
		println("Number is 2")
		fallthrough // This will execute the next case as well
	case 3:
		println("Number is 3")
	default:
		println("Number is something else")
	}
}
