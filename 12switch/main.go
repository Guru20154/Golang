package main

import (
	"fmt"
	"math/rand"
	// "math/rand/v2"
	"time"
)

func main() {
	fmt.Println("Switch and Case")
	//rand.Seed is depricated
	//rand.Seed(time.Now().UnixNano())

	//one way to do this
	// s2 := rand.NewPCG(42, 1024)
	// r2 := rand.New(s2)
	// dicenumber := r2.IntN(6) + 1

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	dicenumber := rand.Intn(6) + 1
	fmt.Println("dice is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 you can move")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces and roll the dice again")
	default:
		fmt.Println("You broke the game")
	}
}
