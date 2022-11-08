package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok // err ok ",_"
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Type fo this rating is %T", input)

}
