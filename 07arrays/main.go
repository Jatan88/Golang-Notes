package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Grapes"

	fmt.Println("Fruit List: ", fruits) // Fruit List:  [Apple Orange  Grapes]

	fruits[2] = "Cherry"
	fmt.Println("Fruit List: ", fruits) // Fruit List:  [Apple Orange Cherry Grapes]

	fmt.Println("Fruit List: ", len(fruits)) // Fruit List:  4

	var veggies = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veggies List: ", veggies)
}
