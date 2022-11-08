package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - https://devdocs.io/go/ "

	file, err := os.Create("./myGoTutFile.txt")

	// if err != nil {
	// 	panic(err)
	// 	}
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myGoTutFile.txt")

}
func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}
func checkNilErr(err error) { // We can create func() err and call the func error everywhere so it does not through an error
	if err != nil {
		panic(err)
	}

}
