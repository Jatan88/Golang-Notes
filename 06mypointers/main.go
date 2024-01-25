package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("value of ptr is, ", ptr) // <nil>

	num := 23
	var ptr1 = &num
	fmt.Println("the value of ptr1 is: ", ptr1)  // 0xc00000a0c8
	fmt.Println("the value of ptr1 is: ", *ptr1) // 23

	*ptr1 = *ptr1 + 2
	fmt.Println("the value of ptr1 is: ", num) // 25

}
