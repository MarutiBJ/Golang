package main

import "fmt"

// Global Variables

var Name string = "Maruti"

// package level variable

var age int = 52

// Constant variable

const mobileNumber string = "904xxxx203"

func main() {
	fmt.Println("Variables")

	// := -> Only used for local variable declaration
	address := "Maruti Krupa, Adarsh Nagar, Pathri"

	fmt.Println(Name, age, mobileNumber, address)
}
