package main

// Function maxVowels calculates the maximum number of vowels in any substring of length k in the given string s.
// It returns the maximum count of vowels among all possible substrings of length k.
func maxVowels(s string, k int) int {
	// Define a map to store vowels for efficient lookup.
	vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	// Initialize variables to track the maximum count of vowels and the current count of vowels in the current substring.
	result := 0
	subResult := 0
	// Initialize the left pointer for the sliding window.
	left := 0

	// Iterate over the characters of the string using the right pointer.
	for right := 0; right < len(s); right++ {
		// If the current character is a vowel, increment the current count of vowels.
		if vowels[string(s[right])] {
			subResult++
		}

		// If the size of the current substring exceeds k, adjust the left pointer and decrement the current count of vowels accordingly.
		if right-left == k {
			// If the character at the left end of the substring is a vowel, decrement the current count of vowels.
			if vowels[string(s[left])] {
				subResult--
			}
			// Move the left pointer to the right to slide the window.
			left++
		}

		// Update the maximum count of vowels encountered so far.
		result = max(result, subResult)
	}

	// Return the maximum count of vowels among all substrings of length k.
	return result
}
