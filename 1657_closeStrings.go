package main

import "slices"

// Function closeStrings checks if two words are "close" under the given conditions.
// Two words are considered "close" if they are of the same length and have the same set of letters
// with the same frequency of each letter, where any letter from one word can be transformed into
// another letter in the other word by swapping some letters.
// It returns true if the words are "close", and false otherwise.
func closeStrings(word1 string, word2 string) bool {
	// Check if the lengths of the two words are different.
	if len(word1) != len(word2) {
		return false
	}

	// Initialize arrays to count the frequency of each letter in the words.
	arrLetters1 := make([]int, 26)
	arrLetters2 := make([]int, 26)

	// Count the frequency of each letter in both words.
	for i := 0; i < len(word1); i++ {
		arrLetters1[word1[i]-'a']++
		arrLetters2[word2[i]-'a']++
	}

	// Check if the set of letters is the same in both words.
	// If any letter is present in one word but not in the other, return false.
	for i := 0; i < len(arrLetters1); i++ {
		if (arrLetters1[i] != 0 && arrLetters2[i] == 0) || (arrLetters1[i] == 0 && arrLetters2[i] != 0) {
			return false
		}
	}

	// Sort the arrays of letter frequencies.
	slices.Sort(arrLetters1)
	slices.Sort(arrLetters2)

	// Check if the frequency of each letter is the same in both words.
	// If any frequency is different, return false.
	for i := 0; i < len(arrLetters1); i++ {
		if arrLetters1[i] != arrLetters2[i] {
			return false
		}
	}

	// If all conditions are met, return true, indicating that the words are "close".
	return true
}
