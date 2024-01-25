package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruits)

	fruits = append(fruits, "Mango", "Banana")
	//fmt.Println(fruits)

	fruits = append(fruits[1:3])
	//fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 345
	highScores[3] = 456
	//highScores[4] = 567 // This will throw an error

	highScores = append(highScores, 555, 666, 321)

	//fmt.Println(highScores)

	sort.Ints(highScores)
	//fmt.Println(highScores)
	//fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slice based on index

	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // [react javascript python ruby]
}
