package main

import (
	"strings"
)

func reverseWords(s string) string {
	// Split the string s into words using strings.Fields, which returns a slice of words
	result := strings.Fields(s)

	// Initialize two pointers, one pointing to the beginning of the slice and the other to the end
	i, j := 0, len(result)-1

	// Iterate through the slice until the pointers meet
	for i < j {
		// Swap the words at positions i and j in the slice
		result[i], result[j] = result[j], result[i]

		// Move the pointers towards each other
		i++
		j--
	}

	// Join the words in the reversed slice into a single string separated by spaces and return the result
	return strings.Join(result, " ")
}
