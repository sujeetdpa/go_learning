package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Taking User input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Name is :", name)
	fmt.Printf("Data Type of name is: %T \n", name)

	//Taking numeric value
	fmt.Println("Enter your rating value : ")
	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Occured while reading rating")
	}

	//Before Conversion
	fmt.Printf("Data Type of rating is %T \n", rating)

	//After Conversion
	convertedRating, converr := strconv.ParseFloat(rating, 64)

	if converr != nil {
		//The error is due to present of \n character in the end of numeric value
		fmt.Println("Error occured while conversion: ", converr)
	}

	//Correct conversion
	fmt.Println("Trimming the spaces and then converting")
	convertedRating, _ = strconv.ParseFloat(strings.TrimSpace(rating), 64)
	fmt.Println("Converted Rating Value: ", convertedRating)
}
