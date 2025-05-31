package main

func main() {
	const a int = 10
	const b int = 20
	const c int = a + b
	println("Sum of a and b is:", c)

	const d string = "Hello, World!"
	println("String value is:", d)

	// complex numbers
	const h complex128 = 1 + 2i
	println("Complex value is:", h)

	const (
		e int = 30
		f int = 40
	)
	println("Sum of e and f is:", e+f)
}
