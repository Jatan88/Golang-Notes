package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println(diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice number is 1")
		fallthrough
	case 2:
		fmt.Println("Dice number is 2")
		fallthrough
	case 3:
		fmt.Println("Dice number is 3")
		fallthrough
	case 4:
		fmt.Println("Dice number is 4")
		fallthrough
	case 5:
		fmt.Println("Dice number is 5")
		fallthrough
	case 6:
		fmt.Println("Dice number is 6")
		fallthrough
	default:
		fmt.Println("Roll once again")
	}
}
