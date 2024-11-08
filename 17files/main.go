package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in go")
	con := "First file creation"
	file, err := os.Create("./firstFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, con)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)

	readFile("./firstFile.txt")
	defer file.Close()
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
