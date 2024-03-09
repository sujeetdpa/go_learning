package main

import "fmt"

// Multiple constant declaration
const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	const PI float32 = 3.14
	const g float32 = 9.8

	fmt.Println("Pi: ", PI)
	fmt.Println("g value: ", g)

	fmt.Println("Multiple contants: ", A, B, C)

	var t int
	println(t)
}
