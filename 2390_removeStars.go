package main

import "strings"

// Function removeStars takes a string s as input and returns a string without "*" characters.
func removeStars(s string) string {
	// Create an empty stack to store characters of the string without "*" characters.
	stack := []string{}

	// Iterate through each character in the input string s.
	for i := 0; i < len(s); i++ {
		// Get the character of string s at index i.
		character := string(s[i])

		// Check if the current character is "*".
		if character != "*" {
			// If the character is not "*", append it to the stack.
			stack = append(stack, character)
		} else {
			// If the character is "*", remove the last character from the stack.
			stack = stack[:len(stack)-1]
		}
	}

	// Return a string obtained from the elements of the stack joined into a single string without a separator.
	return strings.Join(stack, "")
}
