package main

import "fmt"

func main() {
	var stu string = "sujeet"
	var stuid int = 34
	sturank := 2 //Can be done only inside of a function. declaration and initialisation cannot be separated
	fmt.Println("Student Name: ", stu)
	fmt.Println("Student Id: ", stuid)
	fmt.Println("Student Rank: ", sturank)

	var name string
	var boo bool
	name = "name"
	boo = false
	fmt.Println(name)
	fmt.Println(boo)

	///Multiple declaration
	var a, b, c int = 2, 3, 5
	fmt.Println(a, b, c)

	//Variable declaration in block
	var (
		aa int     = 4
		bb string  = "string"
		cc float32 = 5787.89
	)
	fmt.Println(aa, bb, cc)

}
