package main

import (
	"fmt"
)

func main() {

	// Defer -last in first out
	// World, One, Two , Hello -reversed order will be print out
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Print(i)

	}
}
