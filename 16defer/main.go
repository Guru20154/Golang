package main

import "fmt"

func main() {
	//defer line of code are executed at the last of the comiling 
	//however the order of defer is lifo(last in first out or executed in this case)
	defer fmt.Println("defer")
	defer fmt.Println("defer2")
	fmt.Println("About defer")
	mdefer()
}

func mdefer(){
	for i:=0;i<6;i++{
		defer fmt.Println(i)
	}
}
