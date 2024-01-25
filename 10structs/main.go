package main

import "fmt"

func main() {
	fmt.Println("Structs")
	jatan := User{"Jatan", "jatan@go.dev", true, 23}
	fmt.Println(jatan)
	fmt.Printf("Details of Jatan are: %v\n", jatan)                    // Details of Jatan are: {Jatan jatan@go.dev true 23}
	fmt.Printf("Details of Jatan are: %+v\n", jatan)                   // Details of Jatan are: {Name:Jatan Email:jatan@go.dev Status:true Age:23}
	fmt.Printf("Name is %v and Email is %v.", jatan.Name, jatan.Email) // Name is Jatan and Email is jatan@go.dev.

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
