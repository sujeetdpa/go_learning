package main

import (
	"fmt"
)

type DivideError struct {
	Dividend int
}

// DivideError implementing  error interface
func (d DivideError) Error() string {
	return fmt.Sprintf("Error: Divided by %v", d.Dividend)
}

func divide(a int, b int) (int, error) {
	//Custom error
	// if b == 0 {
	// 	return -1, DivideError{Dividend: b}

	// } else {
	// 	return a / b, nil
	// }

	if b == 0 {
		return -1, fmt.Errorf("Error: Divided by %v", b)

	} else {
		return a / b, nil
	}
}

func main() {
	d, err := divide(6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division: ", d)
	}

}
