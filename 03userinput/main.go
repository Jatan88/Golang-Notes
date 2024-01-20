package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type: %T", input)
}
