package main

import "fmt"

func main() {
	var arr = [3]int{3, 6, 9}
	var sarr = [3]string{}
	sarr[0] = "hello"
	sarr[2] = "world"
	fmt.Println(arr) //print whole array
	fmt.Println(sarr)
	fmt.Println(sarr[1]) // print empty string

	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//Initialize only 1st and 2nd
	arr4 := [5]int{1: 10, 2: 40}
	fmt.Println(arr4)

	var narr [4]int
	narr[3] = 6
	fmt.Println(narr)
	test()
}
func test() {
	fmt.Println("hello")
}
