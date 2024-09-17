package main

import "fmt"

func main() {
	// Define an array of integers and a target sum
	s := []int{2, 7, 11, 15}
	target := 9

	// Call the twoSum function and print the result
	fmt.Println(twoSum(s, target))
}

// twoSum function takes an array of integers (nums) and a target sum (target),
// and returns the indices of the two numbers that add up to the target
func twoSum(nums []int, target int) []int {
	// Create an empty slice to store the result (indices of the two numbers)
	var result []int

	// Create a map (hash table) to store the value and its corresponding index
	MapValue := make(map[int]int)

	// Iterate through the nums slice
	for i, num := range nums {
		// Calculate the complement (the number needed to reach the target sum)
		nextnum := target - num

		// Check if the complement exists in the map
		if index, ok := MapValue[nextnum]; ok {
			// If found, append the indices of the complement and the current number to the result
			result = append(result, index, i)
			break // Break the loop as we have found the answer
		} else {
			// If not found, add the current number and its index to the map
			MapValue[num] = i
		}
	}

	// Return the result containing the indices of the two numbers
	return result
}
