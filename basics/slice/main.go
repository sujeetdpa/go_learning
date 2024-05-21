package main

import (
	"fmt"
)

func main() {
	var sl = []int{}
	fmt.Println(len(sl))

	//slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4] //from index 2 to 3 with last one excluded
	fmt.Println(myslice)

	//create slice with make()
	makeSlice := make([]int, 7, 10)
	makeSlice[1] = 4
	makeSlice[2] = 57
	fmt.Println(makeSlice)
	fmt.Printf("type: %T", makeSlice)
}
