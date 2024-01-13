package main

import "fmt"

const LoginToken string = "jatan88"

func main() {
	var username string = "jatan"
	fmt.Println(username)
	fmt.Printf("Variable type: %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("Variable type: %T \n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type: %T \n", smallVal)

	var smallFloat float32 = 32.33333333333333
	fmt.Println(smallFloat)
	fmt.Printf("Variable type: %T \n", smallFloat)

	// implicit type
	var name = "Jatan"
	fmt.Println(name)
	fmt.Printf("Variable type: %T \n", name)

	// no var style
	num := 5
	fmt.Println(num)
	fmt.Printf("Variable type: %T \n", num)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type: %T \n", LoginToken)
}
