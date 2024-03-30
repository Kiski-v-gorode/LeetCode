package main

// MergeAlternately merges two strings alternating between the characters of the two strings.
// If one of the strings is shorter than the other, the remaining characters are appended to the result.
func mergeAlternately(word1, word2 string) string {
	smallest := min(len(word1), len(word2)) // compute smallest word
	result := ""                            // returned result
	i, j := 0, 0                            // two pointers
	for i < smallest && j < smallest {      // walk for loop until i or j not become equal smallest length word
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}

	if smallest == len(word1) { // paste remaining word
		result += word2[smallest:]
	} else {
		result += word1[smallest:]

	}

	return result
}
