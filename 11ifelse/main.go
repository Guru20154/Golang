package main

import "fmt"

func main() {
	fmt.Println("If Else in go")
	login := 23
	var res string
	if login < 10 {
		res = "Not a regular user"
	} else{
		res = "Regular User"
	}
	fmt.Println("User is",res)

	if num := 35; num%5==0{
		fmt.Println("Number is multiple of 5")
	} else{
		fmt.Println("Number is not a multiple of 5")
	}
}
