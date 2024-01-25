package main

import "fmt"

func main() {
	count := 2
	var msg string

	if count > 10 {
		msg = "count is greater than 10"
	} else {
		msg = "count is less than 10"
	}

	fmt.Println(msg)

	num := 10
	var msg1 string

	if num%2 == 0 {
		msg1 = "even"
	} else {
		msg1 = "odd"
	}

	fmt.Println(msg1)

	var msg2 string
	if num1 := 20; num1 > 10 {
		msg2 = "num1 is greater than 10"
	} else {
		msg2 = "num1 is less than 10"
	}

	fmt.Println(msg2)

}
