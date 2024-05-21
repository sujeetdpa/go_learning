package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning About Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//formatting time
	// "01 -02-2006(For date) 15:04:05(For time) Monday(for date name)" The format must use this date value its the standard mention in documentation
	//01---> for month,  02---> for day
	fmt.Println(presentTime.Format("02/01/2006 15:04:05 Monday"))

	//Creating Date
	createDate := time.Date(2021, time.July, 23, 10, 34, 45, 4546, time.UTC)
	fmt.Println("Created Date: ", createDate)
}
