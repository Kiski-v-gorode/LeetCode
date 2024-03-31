package main

import (
	"slices"
)

func reverseVowels(s string) string {
	// Convert the string to a byte slice for mutability
	result := []byte(s)

	// Initialize two pointers, one pointing to the beginning of the string and the other to the end
	i, j := 0, len(s)-1

	// Define the vowels
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	// Iterate through the string until the pointers meet
	for i < j {
		// If both the i-th and j-th characters are vowels, swap them and advance the pointers
		if slices.Contains(vowels, string(result[i])) && slices.Contains(vowels, string(result[j])) {
			result[i], result[j] = result[j], result[i]
			i++
			j--
			// If the i-th character is a vowel but the j-th is not, only decrement j
		} else if slices.Contains(vowels, string(result[i])) && !slices.Contains(vowels, string(result[j])) {
			j--
			// If the i-th character is not a vowel but the j-th is, only increment i
		} else if !slices.Contains(vowels, string(result[i])) && slices.Contains(vowels, string(result[j])) {
			i++
			// If neither the i-th nor the j-th characters are vowels, decrement both pointers
		} else {
			i++
			j--
		}
	}

	// Convert the byte slice back to a string and return the result
	return string(result)
}
