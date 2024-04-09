package main

// Function uniqueOccurrences checks if the number of occurrences of each value
// in the input array arr is unique. It returns true if each value occurs a unique
// number of times, and false otherwise.
func uniqueOccurrences(arr []int) bool {
	// Initialize a map to count the occurrences of each value in arr.
	counter := map[int]int{}
	// Iterate over each element in arr.
	for _, value := range arr {
		// Check if the value already exists in the counter map.
		_, ok := counter[value]
		// If it doesn't exist, initialize its count to 1.
		if !ok {
			counter[value] = 1
		} else {
			// If it exists, increment its count.
			counter[value]++
		}
	}

	// Initialize a map to store unique occurrence counts.
	setCounter := map[int]bool{}
	// Iterate over the counts in the counter map.
	for _, value := range counter {
		// If the count is not already in setCounter, mark it as present.
		if !setCounter[value] {
			setCounter[value] = true
		}
	}

	// Check if the number of unique occurrence counts is equal to the number of distinct values.
	return len(setCounter) == len(counter)
}
