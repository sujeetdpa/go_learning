package main

import "fmt"

type Weekday int

const (
	SUNDAY Weekday = iota + 1
	MONDAY
	TUESDAY
)

func (w Weekday) String() string {
	return [...]string{"SUNDAY", "MONDAY", "TUESDAY"}[w-1]
}

type Student struct {
	Id   int
	Name string
	Day  Weekday
}

func main() {
	dday := SUNDAY
	fmt.Println(dday)
	fmt.Printf("Type: %T", dday)

}
