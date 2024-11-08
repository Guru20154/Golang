package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	fmt.Println("No inheritance no super or parent")

	guru := User{"Gurkanwal", "gurkanwal@gmail.com", true, 23}
	fmt.Println(guru)
	fmt.Printf("User Detail: %+v\n",guru)
	fmt.Printf("Name is %v and Email is %v",guru.Name,guru.Email)
}

//capital letter in User and variable is important if variable name is in small it is not exportable
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
