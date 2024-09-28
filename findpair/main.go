package main

import (
	"fmt"
	"os"
	"strconv"
)

// CheckInvalid checks if the string starts with '[' and ends with ']'.
// It is used to validate the input format.
func CheckInvalid(s string) bool {
	return (s[0] == '[' && s[len(s)-1] == ']')
}

func main() {
	// Retrieve command-line arguments (input string and target sum)
	str := os.Args[1:]

	// Check if the target sum (str[1]) has more than 2 characters, indicating invalid input.
	if len(str[1]) > 2 {
		fmt.Println("Invalid target sum.")
		return
	}

	// Check if the first input string is in the correct format using CheckInvalid.
	if !CheckInvalid(str[0]) {
		fmt.Println("Invalid input")
		return
	}

	// Check for invalid characters (e.g., letters) in the input string.
	if res, found := CheckForCharacter(str[0]); found {
		fmt.Println(res)
		return
	}

	// Convert the target sum string to an integer.
	target, _ := strconv.Atoi(str[1])
	fmt.Println(target)

	// Split the input string into individual numbers.
	str1 := Split(str[0])

	// Find pairs of numbers that add up to the target sum.
	slice := findPairs(str1, target)

	// Print the pairs that sum to the target.
	fmt.Printf("Pairs with sum %d : %v\n", target, slice)
}

// Split takes a string of numbers and returns a slice of individual numbers as strings.
// It removes brackets, commas, spaces, etc., and only keeps the numbers.
func Split(s string) []string {
	var temp string
	var result []string
	for _, c := range s {
		// Concatenate characters to temp if they are not brackets, commas, or spaces.
		if c != '[' && c != ',' && c != ' ' && c != ']' {
			temp += string(c)
		}
		// When encountering separators, append temp to result if it contains a number.
		if (c == '[' || c == ',' || c == ' ' || c == ']') && temp != "" {
			result = append(result, temp)
			temp = ""
		}
	}
	return result
}

// CheckForCharacter checks if the input string contains any lowercase letters.
// If found, it returns an error message and true. Otherwise, it returns false.
func CheckForCharacter(s string) (string, bool) {
	for _, c := range s {
		// Return error message if any character is a lowercase letter.
		if c >= 'a' && c <= 'z' {
			res := fmt.Sprintf("Invalid number: %c", c)
			return res, true
		}
	}
	return "", false
}

// findPairsWithMap finds all pairs of numbers that sum to the target using a map for efficient lookup.
// It returns a slice of pairs, where each pair is represented as two strings (indices of the numbers).
func findPairsWithMap(s []string, target int) (res [][]string) {
	var temp []string
	mapCheck := make(map[int]int) // Map to store numbers and their indices.
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		nextpair := target - num // Calculate the required number to form the target sum.

		// If the required number exists in the map, add the pair (i, index) to the result.
		if index, ok := mapCheck[nextpair]; ok {
			pairone := strconv.Itoa(i)
			pairtwo := strconv.Itoa(index)
			temp = append(temp, pairone, pairtwo)
			res = append(res, temp)
			temp = []string{} // Reset temp after adding the pair.
		}
		mapCheck[num] = i // Store the current number and its index in the map.
	}
	return
}

// findPairs is a brute force method to find all pairs of numbers that sum to the target.
// It returns a slice of pairs, where each pair is represented as two strings (indices of the numbers).
func findPairs(s []string, target int) (res [][]string) {
	var temp []string
	// Loop through the array and check pairs of numbers.
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			num1, _ := strconv.Atoi(s[i])
			num2, _ := strconv.Atoi(s[j])

			// If the sum of the pair equals the target, add the pair of indices to the result.
			if num1+num2 == target {
				pairone := strconv.Itoa(i)
				pairtwo := strconv.Itoa(j)
				temp = append(temp, pairone, pairtwo)
				res = append(res, temp)
				temp = []string{} // Reset temp after adding the pair.
			}
		}
	}
	return
}
