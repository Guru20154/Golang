package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is time module")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("15:04"))

	createdDate := time.Date(2020, time.January, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
