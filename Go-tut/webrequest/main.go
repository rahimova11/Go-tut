package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "hhtps: //lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)

	}

	fmt.Printf("Responce is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)

	}
	content := string(databytes)
	fmt.Println(content)
}
