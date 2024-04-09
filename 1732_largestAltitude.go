package main

// Function largestAltitude calculates the largest altitude reached during a journey.
// The function takes a slice of integers representing the gain in altitude at each step of the journey.
// It iterates over the gain values, accumulating the total altitude reached (prefix sum),
// and updates the result with the maximum altitude encountered.
// It returns the largest altitude reached during the journey.
func largestAltitude(gain []int) int {
	// Initialize variables to track the result (largest altitude reached) and the prefix sum.
	result := 0
	prefix := 0

	// Iterate over the gain values.
	for i := 0; i < len(gain); i++ {
		// Update the prefix sum by adding the gain at the current step.
		prefix += gain[i]
		// Update the result with the maximum altitude encountered so far.
		result = max(result, prefix)
	}

	// Return the largest altitude reached during the journey.
	return result
}
