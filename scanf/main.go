package main

import "fmt"

func main() {
	// Declare two string variables to hold the user's name and age
	var Name, Age string

	// Prompt the user to enter their name
	fmt.Printf("Enter Your Name: ")
	// Read the user's name using fmt.Scanf, but it will only capture the first word
	fmt.Scanf("%s", &Name)

	// Prompt the user to enter their age
	fmt.Printf("Enter Your Age: ")
	// Read the user's age using fmt.Scanf
	fmt.Scanf("%s", &Age)

	// Print the name and age entered by the user
	fmt.Printf("Your Name is %s and your age is %s\n", Name, Age)
}
