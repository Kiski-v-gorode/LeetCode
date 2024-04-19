package main

import (
	"strconv"
	"strings"
	"unicode"
)

// The decodeString function takes a string s and decodes the encoded string
// using the following rules:
// - Any number followed by '[' indicates that the characters inside the square
// brackets should be repeated as many times as specified by the number.
// - Square brackets can be nested.
// Returns the decoded string.
func decodeString(s string) string {
	// Create a stack to store intermediate results
	stack := []string{""}
	// Variable to store the number indicating repetitions
	number := 0
	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		// If the character is a digit, add it to the number
		if unicode.IsDigit(rune(s[i])) {
			num, _ := strconv.Atoi(string(s[i]))
			number = number*10 + num
			// If the character is '[', push the current number onto the stack
			// and create a new element in the stack for a new string
		} else if string(s[i]) == "[" {
			stack = append(stack, strconv.Itoa(number))
			stack = append(stack, "")
			number = 0
			// If the character is ']', it marks the end of the inner string,
			// so repeat it and append it to the previous string in the stack
		} else if string(s[i]) == "]" {
			// Extract the last element from the stack, which is the inner string
			str1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Extract the previous element from the stack, which is the repetition number
			repeat, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			// Extract the second last element from the stack, which is the outer string
			str2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Add the repeated inner string to the outer string and push the result back to the stack
			stack = append(stack, str2+strings.Repeat(str1, repeat))
			// If the character is neither a digit, '[' nor ']', add it to the current string in the stack
		} else {
			stack[len(stack)-1] = stack[len(stack)-1] + string(s[i])
		}
	}
	// Return the concatenated elements of the stack as a single string
	return strings.Join(stack, "")
}
