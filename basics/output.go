package main

import "fmt"

func main() {
	var i int = 785
	var s string = "hello"
	var b bool = true

	//%v for value
	//%T for type
	fmt.Printf("i value= %v and type=%T\n", i, i)
	fmt.Printf("s value= %v and type=%T\n", s, s)
	fmt.Printf("b value= %v and type=%T\n", b, b)

	i = 15

	fmt.Printf("%b\n", i)   // Base 2
	fmt.Printf("%d\n", i)   // Base 10 -> Decimal
	fmt.Printf("%+d\n", i)  // Base 10 print sign also
	fmt.Printf("%o\n", i)   //Base 8
	fmt.Printf("%x\n", i)   //Base 16 lower case
	fmt.Printf("%X\n", i)   //Base 16 upper case
	fmt.Printf("%#x\n", i)  //Base 16, with leading 0x
	fmt.Printf("%4d\n", i)  //Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", i) //Pad with spaces (width 4, left justified)
	fmt.Printf("%05d\n", i) //Pad with zeroes (width 5
}
