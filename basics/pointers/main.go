package main

import "fmt"

func main() {
	var intPointer *int
	x := 5
	intPointer = &x
	fmt.Println("Address: ", intPointer)
	fmt.Println("Value: ", *intPointer)

	str := "hello"
	strptr := &str
	fmt.Println("String pointer: ", strptr)
	fmt.Println("String Value: ", *strptr)
}
