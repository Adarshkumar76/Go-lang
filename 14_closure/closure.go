package main

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	println(c()) // Output: 1
	println(c()) // Output: 2
}
