package main

import "fmt"

func main() {
	var slice = []string{"Apple", "Tomato"}
	fmt.Printf("Type of slice is %T\n", slice)

	slice = append(slice, "mango", "bannanan")
	fmt.Println(slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Println(slice)
}
