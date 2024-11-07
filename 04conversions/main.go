package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our app from 1-5: ")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("Thank for rating ",input)

	stringrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(stringrating+1)
	}
}
