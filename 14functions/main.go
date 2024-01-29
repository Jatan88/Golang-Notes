package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greet()
	greetTwo()

	result := add(2, 2)
	fmt.Println("sum is:", result)

	proResult, msg := proAdd(1, 2, 3, 4, 5)
	fmt.Println("Pro sum is:", proResult)
	fmt.Println("Pro msg is:", msg)
}

func greet() {
	fmt.Println("Namaste")
}

func greetTwo() {
	fmt.Println("greet 2")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proAdd(num ...int) (int, string) {
	total := 0
	for _, val := range num {
		total += val
	}

	return total, "The result is from Pro Add Function "
}
