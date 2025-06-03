package main

func add(a, b int) int {
	return a + b
}

func getLang() (string, string, int) {
	return "Go", "java", 82
}

// passing a function as an argument
func processIt(fn func(a int) int) {
	result := fn(5)
	println("Processed integer:", result)
}

// returning a function
func returnFunc() func(a int) int {
	return func(a int) int {
		return a + 10
	}
}

func main() {
	result := add(3, 4)
	println("The result is:", result)
	a, b, c := getLang()
	println("Languages:", a, b, "and number:", c)

	processIt(func(a int) int {
		return a * 2
	})
	funcResult := returnFunc()
	println("Returned function result:", funcResult(5))
}
