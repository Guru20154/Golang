package main

import "fmt"

func greeter() {
	fmt.Println("Hello lost friend")
}
func main() {
	fmt.Println("Functions in golang")
	greeter()

	res, str := adder(3,5)
	fmt.Printf("Reuslt is %v and message is %v",res,str)
	
	res2 := dynamadder(3,5,6,4,3,2,5,)
	fmt.Println("Reuslt is ",res2)
}

func dynamadder(values ...int) int {
	total := 0;
	for _,value := range values{
		total+=value
	} 
	return total
}

func adder(valone int, valtwo int) (int,string) {
	return valone+valtwo, "Hello there"
}

