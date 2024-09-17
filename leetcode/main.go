package main

import (
	"fmt"
	"strings"
)

// uncommonFromSentences takes two sentences as input, splits them into words,
// and returns a slice of words that appear exactly once across both sentences.
func uncommonFromSentences(s1 string, s2 string) []string {
	// Split the sentences into slices of words.
	slice1 := strings.Split(s1, " ")
	slice2 := strings.Split(s2, " ")

	// Populate a map with word counts from both sentences.
	Mapaccumulate := PopulateMap(s1, s2)

	// Initialize the result slice to store uncommon words.
	var res []string

	// Check words from the first sentence. If they appear only once, add to result.
	for _, word := range slice1 {
		if Mapaccumulate[word] == 1 {
			res = append(res, word)
		}
	}

	// Check words from the second sentence. If they appear only once, add to result.
	for _, word := range slice2 {
		if Mapaccumulate[word] == 1 {
			res = append(res, word)
		}
	}

	return res
}

// PopulateMap takes two sentences and returns a map with the count of each word across both sentences.
func PopulateMap(s1 string, s2 string) map[string]int {
	// Initialize an empty map to store word counts.
	Mapaccumulate := make(map[string]int)

	// Split both sentences into words.
	slice1 := strings.Split(s1, " ")
	slice2 := strings.Split(s2, " ")

	// Count occurrences of each word in the first sentence.
	for _, word := range slice1 {
		Mapaccumulate[word]++
	}

	// Count occurrences of each word in the second sentence.
	for _, word := range slice2 {
		Mapaccumulate[word]++
	}

	// Return the populated map of word counts.
	return Mapaccumulate
}

func main() {
	// Example input sentences
	s1 := "this apple is sweet"
	s2 := "this apple is sour"

	// Print the list of uncommon words from both sentences.
	fmt.Println(uncommonFromSentences(s1, s2)) // Output: [sweet sour]
}
