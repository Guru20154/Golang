package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://hiteshchoudhary.com/"

func main() {
	fmt.Println("Web request")
	response,err := http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Type of req %T\n", response)

	//caller is responsible to close the response
	defer response.Body.Close()

	databytes,err := io.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
