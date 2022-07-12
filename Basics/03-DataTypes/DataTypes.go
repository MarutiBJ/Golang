package main

import "fmt"

func main() {
	fmt.Println("Data Types in Golang")

	// Numbers

	var num1 int = 10 // int8, int16
	var num2 float64 = 20.21
	var complex complex128 = 3 + 5i

	var positiveNumber uint = 10 // Only Positive numbers are allowed

	fmt.Println("Addition: ", num1+int(num2))
	fmt.Println("Complex Number: ", real(complex), imag(complex), complex)
	fmt.Println(positiveNumber)

	// Strings

	var Name string = "Maruti Jogdand"
	fmt.Println("Name : ", Name)

	// Boolean

	var isValid bool = false
	fmt.Println(isValid)

	// Rune -> stores a single character ascii code for alphabets

	var space rune = '$'
	fmt.Println(space)

	// Array :- Fixed Size

	// Declaration
	var arr [3]int // [0,0,0]
	// Initialize to default values [0,0,0] even if we add single element to array still len of array is 3 as we mentioned while delcaration
	fmt.Println(arr)

	// Declaration and Initialization
	var Names [5]string = [5]string{"Maruti", "Atul"}
	fmt.Println(Names, len(Names), Names[1])

	// Slice -> Go Specific Data Type prefered over Array, It is faster and Dynamic

	// Declaration -> use make function to initialization
	var slice []int
	fmt.Println("Slice: ", slice) // Initialization to empty list -> []

	slice = make([]int, 0)

	var Numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println(Numbers, len(Numbers), Numbers[3])

	// Map -> Key Value Pair i.e Dictionary, Map always needs to initialize use make function to initialize

	// Declaration
	var Employees map[int]string

	// Default initialization
	Employees = make(map[int]string, 0) // always give size as zero

	Employees[1] = "Maruti Jogdand"
	Employees[2] = "Atul Jogdand"

	fmt.Println(Employees)

}
