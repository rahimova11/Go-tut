package main

import (
	"fmt"
	"sort"
	"time"
)

// declare constance - they cannot be changed - are declared outside of functions
const LoginToken string = "ghabbhhjd" // Public

func main() {
	var username string = "Gulchehra"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	// Go - time
	fmt.Println()
	fmt.Println("Welcome to time study of golang!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	// Reversed time
	createdDate := time.Date(2020, time.October, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	fmt.Println("")

	presentTime.Format("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	// & - referencing
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

	// Array - is very less used data structure in GO

	// var fruitList [4]string
	// fruitList[0] = "Apple"
	// fruitList[1] = "Tomato"
	// fruitList[3] = "Peach"

	// fmt.Println("Fruit list is: ", fruitList)
	// fmt.Println("Fruit list is: ", len(fruitList))

	// var vegList = [5]string{"potato", "beans", "mushroom"}
	// fmt.Println("Veggie list is: ", len(vegList))

	fmt.Println("")

	// Slices - most used data stucture
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana") // append() method takes the list and adds Mango and Banana to fruitList
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore)) // IntsAreSorted - gives True or False
	sort.Ints(highScore)                       // Inst sorts a lice of integer in increasing order
	fmt.Println(highScore)

	fmt.Println("")

	// How to remove a value from Slice based on Index?  Append() re-alocates the memory
	var courses = []string{"reactjs", "java", "swift", "ruby", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	fmt.Println("")

	// Maps - key & value pairs
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB") // delete() method removes value from map
	fmt.Println("List of all languages: ", languages)

	//Loops are interesting in Golang
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)

		fmt.Println("")
	}
	fmt.Println("Structs in Golang")
	// there is No inheritance in Golang; No super or parent

	gulchehra := User{"Gulchehra", "gulchehra@go.dev", true, 16}
	fmt.Println(gulchehra)
	fmt.Printf("gulchehra details are: %+v\n", gulchehra)
	fmt.Printf("Name is %v and email is %v.", gulchehra.Name, gulchehra.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
