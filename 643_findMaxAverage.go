package main

// findMaxAverage calculates the maximum average of subarrays of length `k`
// within the given slice of integers `nums`.
func findMaxAverage(nums []int, k int) float64 {
	// Calculate the sum of the first `k` elements using the `sum` function.
	prefix := sum(nums, k)
	// Initialize `maxAvg` to the average of the first `k` elements.
	maxAvg := float64(prefix) / float64(k)
	// Iterate through the remaining elements of `nums`.
	for i := k; i < len(nums); i++ {
		// Update the `prefix` sum by adding the current element and subtracting
		// the element that falls outside the sliding window of size `k`.
		prefix = prefix + nums[i] - nums[i-k]
		// Calculate the average of the current subarray and update `maxAvg`
		// if the current average is greater than the current maximum average.
		maxAvg = max(maxAvg, float64(prefix)/float64(k))
	}
	// Return the maximum average found.
	return maxAvg
}

// sum calculates the sum of the first `index` elements in the given slice of integers `nums`.
func sum(nums []int, index int) int {
	result := 0
	// Iterate through the first `index` elements and accumulate their sum.
	for i := 0; i < index; i++ {
		result += nums[i]
	}
	// Return the accumulated sum.
	return result
}
