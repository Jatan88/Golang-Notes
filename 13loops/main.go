package main

import "fmt"

func main() {

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index %v and value %v\n", index, day)
	// }

	num := 1

	for num < 10 {

		if num == 2 {
			goto lco
		}

		if num == 5 {
			//break
			num++
			continue
		}

		fmt.Println(num)
		num++
	}

lco:
	fmt.Println("welcome to lco")
}
