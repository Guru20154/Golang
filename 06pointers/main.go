package main

import "fmt"

func main() {
	mynumber := 23

	var pointer = &mynumber

	fmt.Println(pointer)
	fmt.Println(*pointer)

	fmt.Println("*************")

	fmt.Println(mynumber)
	fmt.Println(&mynumber)
}
