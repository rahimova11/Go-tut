package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// there is No inheritance in Golang; No super or parent

	gulchehra := User{"Gulchehra", "gulchehra@go.dev", true, 16}
	fmt.Println(gulchehra)
	fmt.Printf("gulchehra details are: %+v\n", gulchehra)

	gulchehra.GetStatus()
	gulchehra.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", gulchehra.Name, gulchehra.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // Method
	fmt.Println("Is uer active: ", u.Status)

}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)

}
