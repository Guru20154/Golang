package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number:")

	//comma ok syntax || error ok syntax 
	input, _ := reader.ReadString('\n')
	fmt.Println("Your number is: ",input)
	fmt.Printf("Input type %T",input)
}
