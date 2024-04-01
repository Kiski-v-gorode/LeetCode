package main

import (
	"strconv" // The strconv package is used for converting numbers to strings
)

// The compress function takes a byte slice chars as input and returns an integer value
func compress(chars []byte) int {
	// Initialize a character counter to count repeating characters
	count := 1

	// Iterate from the second-to-last character to the beginning
	for i := len(chars) - 2; i >= 0; i-- {
		// If the current character is equal to the next character
		if chars[i] == chars[i+1] {
			// Increment the count of repeating characters
			count++
			// Remove the repeating character from the slice
			chars = append(chars[:i], chars[i+1:]...)
		} else {
			// If the count of repeating characters is greater than 1
			if count > 1 {
				// Convert the count to a string and append it to the chars slice
				for index, value := range strconv.Itoa(count) {
					chars = append(chars[:i+index+2], append([]byte{byte(value)}, chars[i+2+index:]...)...)
				}
			}
			// Reset the counter as a new character is encountered
			count = 1
		}
	}

	// After the loop, handle the last character
	if count > 1 {
		// Convert the count to a string and append it to the chars slice
		for index, value := range strconv.Itoa(count) {
			chars = append(chars[:index+1], append([]byte{byte(value)}, chars[1+index:]...)...)
		}
	}
	// Return the length of the chars slice
	return len(chars)
}
